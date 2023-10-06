package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/tiozafrem/debtors/models"
	"golang.org/x/exp/slog"
)

type ServiceAuthorization struct {
	client *auth.Client
}

// Struct for send refresh token
type userRefresh struct {
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

// Struct for send auth date
type signInBody struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

// Struct for recieve auth date
type userResponse struct {
	IdToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
}

// Struct for recieve auth date
type refreshResponse struct {
	RefreshToken string `json:"refresh_token"`
	IdToken      string `json:"id_token"`
}

func NewAuthorizationService(client *auth.Client) *ServiceAuthorization {
	return &ServiceAuthorization{client: client}
}

func (s *ServiceAuthorization) ParseTokenToUserUUID(ctx context.Context, token string) (string, error) {
	user, err := s.client.VerifyIDTokenAndCheckRevoked(ctx, token)
	return user.UID, err
}

func (s *ServiceAuthorization) SignIn(email, password string) (*models.Tokens, error) {
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s",
		os.Getenv("FIREBASE_API_KEY"))
	reqBody := signInBody{
		Email:             email,
		Password:          password,
		ReturnSecureToken: true,
	}
	buffer := new(bytes.Buffer)
	if err := json.NewEncoder(buffer).Encode(reqBody); err != nil {
		slog.Error("error encode sign in: %v", err)
		return nil, err
	}

	resp, err := http.Post(url, "application/json", buffer)
	if err != nil {
		slog.Error("error send sign in: %v", err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)

		return nil, fmt.Errorf(string(body))
	}

	decoder := json.NewDecoder(resp.Body)

	firebaseAuthResponse := userResponse{}

	if err := decoder.Decode(&firebaseAuthResponse); err != nil ||
		strings.Trim(firebaseAuthResponse.IdToken, " ") == "" {
		slog.Error("error decode sign in: %v", err)
		return nil, err
	}
	return &models.Tokens{AccessToken: firebaseAuthResponse.IdToken,
		RefreshToken: firebaseAuthResponse.RefreshToken}, nil
}

func (s *ServiceAuthorization) SignUp(ctx context.Context, email, password string) (*models.Tokens, error) {
	createUser := auth.UserToCreate{}
	createUser.Password(password)
	createUser.Email(email)
	_, err := s.client.CreateUser(ctx, &createUser)
	if err != nil {
		slog.Error("error create user: %v", err)
		return nil, err
	}
	return s.SignIn(email, password)
}

func (s *ServiceAuthorization) RefreshToken(refreshToken string) (*models.Tokens, error) {
	url := fmt.Sprintf("https://securetoken.googleapis.com/v1/token?key=%s",
		os.Getenv("FIREBASE_API_KEY"))
	reqBody := newUserRefresh(refreshToken)
	buffer := new(bytes.Buffer)
	if err := json.NewEncoder(buffer).Encode(reqBody); err != nil {
		slog.Error("error encode sign in: %v", err)
		return nil, err
	}

	resp, err := http.Post(url, "application/json", buffer)
	if err != nil {
		slog.Error("error send refresh: %v", err)
		return nil, err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	refreshResponse := refreshResponse{}
	if err := decoder.Decode(&refreshResponse); err != nil ||
		strings.Trim(refreshResponse.IdToken, " ") == "" {
		slog.Error("error decode sign in: %v", err)
		return nil, err
	}
	return &models.Tokens{AccessToken: refreshResponse.IdToken,
		RefreshToken: refreshResponse.RefreshToken}, nil
}

func newUserRefresh(refreshToken string) *userRefresh {
	return &userRefresh{GrantType: "refresh_token", RefreshToken: refreshToken}
}
