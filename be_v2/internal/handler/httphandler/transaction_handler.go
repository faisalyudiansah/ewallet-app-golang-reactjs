package httphandler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"ewallet-server-v2/internal/constant"
	"ewallet-server-v2/internal/dto/appdto"
	"ewallet-server-v2/internal/dto/httpdto"
	"ewallet-server-v2/internal/dto/pagedto"
	"ewallet-server-v2/internal/pkg/apperror"
	"ewallet-server-v2/internal/pkg/ctxutils"
	"ewallet-server-v2/internal/pkg/database"
	"ewallet-server-v2/internal/pkg/ginutils"
	"ewallet-server-v2/internal/usecase"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionUsecase usecase.TransactionUsecase
	userUsecase        usecase.UserUsecase
	walletUsecase      usecase.WalletUsecase
	transactor         database.Transactor
}

func NewTransactionHandler(
	transactionUsecase usecase.TransactionUsecase,
	userUsecase usecase.UserUsecase,
	walletUsecase usecase.WalletUsecase,
	transactor database.Transactor,
) *TransactionHandler {
	return &TransactionHandler{
		transactionUsecase: transactionUsecase,
		userUsecase:        userUsecase,
		walletUsecase:      walletUsecase,
		transactor:         transactor,
	}
}

func (h *TransactionHandler) TopUp(c *gin.Context) {
	var req httpdto.TopUpRequest
	var res httpdto.TopUpResponse

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(err)
		return
	}

	userId := ctxutils.GetUserId(c)

	err = h.transactor.Transaction(c, func(txCtx context.Context) error {
		wallet, err := h.walletUsecase.GetOneByUserId(txCtx, userId)
		if err != nil {
			return err
		}

		transaction, err := h.transactionUsecase.Topup(txCtx, wallet.WalletId, req.Amount, req.SourceOfFundsId)
		if err != nil {
			return err
		}

		res = httpdto.ConvertToTopUpResponse(transaction, wallet)

		return nil
	})
	if err != nil {
		c.Error(err)
		return
	}

	ginutils.ResponseSuccessJSON(c, http.StatusCreated, constant.MessageResponseSuccess, res)
}

func (h *TransactionHandler) Transfer(c *gin.Context) {
	var req httpdto.TransferRequest
	var res httpdto.TransferResponse

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(err)
		return
	}

	userId := ctxutils.GetUserId(c)

	err = h.transactor.Transaction(c, func(txCtx context.Context) error {
		walletFrom, err := h.walletUsecase.GetOneByUserId(txCtx, userId)
		if err != nil {
			return err
		}

		walletTo, err := h.walletUsecase.GetOneByNumber(txCtx, req.WalletTo)
		if err != nil {
			return err
		}

		transaction, err := h.transactionUsecase.Transfer(txCtx, walletFrom.WalletId, req.WalletTo, req.Amount, req.Description)
		if err != nil {
			return err
		}

		res = httpdto.ConvertToTransferResponse(transaction, walletFrom, walletTo)

		return nil
	})
	if err != nil {
		c.Error(err)
		return
	}

	ginutils.ResponseSuccessJSON(c, http.StatusCreated, constant.MessageResponseSuccess, res)
}

func (h *TransactionHandler) GetTransactionList(c *gin.Context) {
	var req httpdto.GetTransactionListRequest
	var res httpdto.GetTransactionListResponse

	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.Error(err)
		return
	}

	sorts := []pagedto.SortDto{}
	for i, s := range req.SortBy {
		sortDir := "asc"
		if len(req.SortDir) > i {
			sortDir = req.SortDir[i]
		}

		sorts = append(sorts, pagedto.SortDto{
			Column:   s,
			OrderDir: sortDir,
		})
	}

	pageDto := pagedto.PageSortDto{
		Page:            req.Page,
		Limit:           req.Limit,
		Sorts:           sorts,
		Search:          req.Search,
		FilterStartDate: req.FilterStartDate,
		FilterEndDate:   req.FilterEndDate,
		FilterType:      req.FilterType,
	}

	userId := ctxutils.GetUserId(c)

	wallet, err := h.walletUsecase.GetOneByUserId(c, userId)
	if err != nil {
		c.Error(err)
		return
	}

	transactionList, err := h.transactionUsecase.GetListByWalletId(c, wallet.WalletId, pageDto)
	if err != nil {
		c.Error(err)
		return
	}

	res = httpdto.ConvertToGetTransactionListResponse(transactionList)

	ginutils.ResponseOKData(c, res)
}

func (h *TransactionHandler) GetTransactiontype(c *gin.Context) {
	var res []appdto.TransactionTypeResponse

	transactionTypes, err := h.transactionUsecase.GetTransactionType(c)
	if err != nil {
		c.Error(err)
	}

	res = appdto.ConvertToTransactionTypeDto(transactionTypes)

	ginutils.ResponseOKData(c, res)
}

func (h *TransactionHandler) GetThisMonthExpenseSum(c *gin.Context) {
	userId := ctxutils.GetUserId(c)

	res, err := h.transactionUsecase.GetThisMonthExpenseSum(c, userId, time.Now())
	if err != nil {
		c.Error(err)
		return
	}

	ginutils.ResponseOKData(c, gin.H{
		"sum":   res.Sum,
		"month": res.Month.Month(),
	})
}

func ConvertMonthNameToNumber(monthName string) (int, error) {
	t, err := time.Parse("January", monthName)
	if err != nil {
		return 0, err
	}
	monthNumber := int(t.Month())
	return monthNumber, nil
}

func (h *TransactionHandler) GetExpenseSumByMonth(c *gin.Context) {
	req := appdto.TransactionByMonthParams{}

	userId := ctxutils.GetUserId(c)
	err := c.ShouldBindUri(&req)
	if err != nil {
		c.Error(err)
		return
	}

	monthNum, err := ConvertMonthNameToNumber(req.Month)
	if err != nil {
		c.Error(apperror.NewAppError(
			errors.New("please input a valid month name ('January', 'February', etc...)"),
			http.StatusBadRequest,
			"please input a valid month name ('January', 'February', etc...)",
			nil,
		))
		return
	}

	res, err := h.transactionUsecase.GetExpenseSumByMonth(c, userId, time.Month(monthNum))
	if err != nil {
		c.Error(err)
		return
	}

	ginutils.ResponseOKData(c, gin.H{
		"sum":   res.Sum,
		"month": res.Month.Month(),
	})
}

// TODO: Refactor get expense code
// TODO: Get income by month
