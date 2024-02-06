package bot

// func (h *Handler) getSumTransactionDebtorUser(b *gotgbot.Bot, ctx *ext.Context) error {
// 	id := c.Param("uuid")
// 	if id == "" {
// 		newErrorResponse(c, http.StatusBadRequest, "uuid is null")
// 		return nil
// 	}

// 	userUUID, err := getUserUUID(c)
// 	if err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, err.Error())
// 		return nil
// 	}
// 	value, err := h.service.User.GetSumTransactionDebtor(c, userUUID, id)
// 	if err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, err.Error())
// 		return nil
// 	}

// 	c.JSON(http.StatusOK, value)
// 	return nil
// }

// func (h *Handler) getSumTransactionDebtorsUser(b *gotgbot.Bot, ctx *ext.Context) error {

// 	userUUID, err := getUserUUID(c)
// 	if err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, err.Error())
// 		return nil
// 	}
// 	value, err := h.service.User.GetSumTransactionDebtors(c, userUUID)
// 	if err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, err.Error())
// 		return nil
// 	}

// 	c.JSON(http.StatusOK, value)
// 	return nil
// }

// func (h *Handler) getSumTransactionMy(b *gotgbot.Bot, ctx *ext.Context) error {

// 	userUUID, err := getUserUUID(c)
// 	if err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, err.Error())
// 		return nil
// 	}
// 	value, err := h.service.User.GetSumMy(c, userUUID)
// 	if err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, err.Error())
// 		return nil
// 	}

// 	c.JSON(http.StatusOK, value)
// 	return nil
// }
