package helpers

import (
	"fmt"

	"ewallet-server-v1/dtos"

	"github.com/go-playground/validator/v10"
)

func FormatterErrorInput(ve validator.ValidationErrors) []dtos.ResponseApiError {
	result := make([]dtos.ResponseApiError, len(ve))
	for i, fe := range ve {
		result[i] = dtos.ResponseApiError{
			Field: jsonFieldName(fe.Field()),
			Msg:   msgForTag(fe.Tag(), fe.Namespace()),
		}
	}
	return result
}

func FormatterMessageWithOneUser(data *dtos.ResponseUser, msg string) dtos.ResponseShowOneUser {
	res := dtos.ResponseShowOneUser{
		Message: msg,
		Result:  *data,
	}
	return res
}

func FormatterSuccessCreateGachaBox(data []int, msg string) dtos.ResponseShowListGachaBox {
	res := dtos.ResponseShowListGachaBox{
		Message: msg,
		Result:  data,
	}
	return res
}

func FormatterSuccessChooseGame(data float64, msg string) dtos.ResponseShowReward {
	total := fmt.Sprintf("Rp. %v", data)
	res := dtos.ResponseShowReward{
		Message: msg,
		Result:  total,
	}
	return res
}

func FormatterMessageWithOneUserAndWallet(data *dtos.ResponseUserAndWallet, msg string) dtos.ResponseShowOneUserWithWallet {
	res := dtos.ResponseShowOneUserWithWallet{
		Message: msg,
		Result:  *data,
	}
	return res
}

func FormatterSuccessLogin(data *dtos.ResponseAccessToken, msg string) dtos.ResponseLoginUser {
	res := dtos.ResponseLoginUser{
		Message: msg,
		Result:  *data,
	}
	return res
}

func FormatterSuccessForgetPassword(data *dtos.ResponseTokenResetPassword, msg string) dtos.ResponseForgetPassword {
	res := dtos.ResponseForgetPassword{
		Message: msg,
		Result:  *data,
	}
	return res
}

func FormatterSuccessTransfer(data *dtos.ResponseSuccessTransfer, msg string) dtos.ResponseSuccessTransferWithMessage {
	res := dtos.ResponseSuccessTransferWithMessage{
		Message: msg,
		Result:  *data,
	}
	return res
}

func FormatterSuccessGetListTransaction(data []dtos.ResponseSingleDataTransactionUser, msg string, limit, page, pageCount, totalCount int64) dtos.ResponseListTransactions {
	res := dtos.ResponseListTransactions{
		Message:        msg,
		Limit:          limit,
		Page:           page,
		PageCount:      pageCount,
		TotalCountData: totalCount,
		Result:         data,
	}
	return res
}
