package controllers

import (
	"net/http"
	"strconv"
	"time"

	"ewallet-server-v1/apperrors"
	"ewallet-server-v1/constants"
	"ewallet-server-v1/helpers"
	helpercontext "ewallet-server-v1/helpers/helper_context"
	"ewallet-server-v1/services"

	"github.com/gin-gonic/gin"
)

type TransactionUserController struct {
	TransactionUserService services.TransactionServiceInterface
	GetParam               helpers.GetParamInterface
}

func NewTransactionUserController(txu services.TransactionServiceInterface, gp helpers.GetParamInterface) *TransactionUserController {
	return &TransactionUserController{
		TransactionUserService: txu,
		GetParam:               gp,
	}
}

func (uc *TransactionUserController) GetListTransactionsController(c *gin.Context) {
	query := c.DefaultQuery("s", "")
	sortBy := c.DefaultQuery("sortBy", "transaction_time")
	sort := c.DefaultQuery("sort", "desc")
	limit := c.DefaultQuery("limit", "10")
	page := c.DefaultQuery("page", "1")
	startDate := c.DefaultQuery("startDate", "")
	endDate := c.DefaultQuery("endDate", "")

	_, errStartDate := time.Parse("2006-01-02", startDate)
	_, errEndDate := time.Parse("2006-01-02", endDate)
	if (startDate != "" && errStartDate != nil) || (endDate != "" && errEndDate != nil) {
		c.Error(apperrors.ErrInvalidDateFormat)
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.Error(apperrors.ErrInvalidQueryLimit)
		return
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.Error(apperrors.ErrInvalidQueryPage)
		return
	}
	offset := (pageInt - 1) * limitInt

	data, totalCount, err := uc.TransactionUserService.GetListTransactionsUserService(c, helpercontext.GetValueUserIdFromToken(c), query, sortBy, sort, int64(limitInt), int64(offset), startDate, endDate)
	if err != nil {
		c.Error(err)
		return
	}
	pageCount := (totalCount + int64(limitInt) - 1) / int64(limitInt)
	helpers.PrintResponse(c, http.StatusOK, helpers.FormatterSuccessGetListTransaction(data, constants.Ok, int64(limitInt), int64(pageInt), pageCount, totalCount))
}
