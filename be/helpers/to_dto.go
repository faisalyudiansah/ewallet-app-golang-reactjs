package helpers

import (
	"fmt"

	"ewallet-server-v1/dtos"
	"ewallet-server-v1/models"
)

func ToResponseUser(dataUser *models.User) *dtos.ResponseUser {
	return &dtos.ResponseUser{
		ID:         dataUser.ID,
		Email:      dataUser.Email,
		ChanceGame: dataUser.ChanceGame,
		FullName:   dataUser.FullName,
		BirthDate:  dataUser.BirthDate,
		CreatedAt:  dataUser.CreatedAt,
		UpdatedAt:  dataUser.UpdatedAt,
		DeleteAt:   dataUser.DeleteAt,
	}
}

func ToResponseUserButNotPointer(dataUser models.User) dtos.ResponseUser {
	return dtos.ResponseUser{
		ID:         dataUser.ID,
		Email:      dataUser.Email,
		ChanceGame: dataUser.ChanceGame,
		FullName:   dataUser.FullName,
		BirthDate:  dataUser.BirthDate,
		CreatedAt:  dataUser.CreatedAt,
		UpdatedAt:  dataUser.UpdatedAt,
		DeleteAt:   dataUser.DeleteAt,
	}
}

func ToResponseUserAndWallet(dataUser *models.UserAndWallet) *dtos.ResponseUserAndWallet {
	return &dtos.ResponseUserAndWallet{
		ID:         dataUser.User.ID,
		Email:      dataUser.User.Email,
		ChanceGame: dataUser.User.ChanceGame,
		FullName:   dataUser.User.FullName,
		BirthDate:  dataUser.User.BirthDate,
		CreatedAt:  dataUser.User.CreatedAt,
		UpdatedAt:  dataUser.User.UpdatedAt,
		DeleteAt:   dataUser.User.DeleteAt,
		Wallet:     dtos.ResponseWallet(dataUser.Wallet),
	}
}

func ToResponseAccessToken(ac string) *dtos.ResponseAccessToken {
	return &dtos.ResponseAccessToken{
		AccessToken: ac,
	}
}

func ToResponseTokenResetPassword(ac string) *dtos.ResponseTokenResetPassword {
	URL := fmt.Sprintf("/reset-password/%s", ac)
	return &dtos.ResponseTokenResetPassword{
		LinkResetPassword:  URL,
		TokenResetPassword: ac,
	}
}

func FormatterTransactionsList(data []models.TransactionUserAndSourceOfFund) []dtos.ResponseSingleDataTransactionUser {
	var res []dtos.ResponseSingleDataTransactionUser
	for _, item := range data {
		res = append(res, dtos.ResponseSingleDataTransactionUser{
			ID:              item.Transaction.ID,
			SourceId:        item.Transaction.SourceId,
			SourceOfFund:    toSourceOfFundDTO(item.SourceOfFund),
			RecipientId:     item.Transaction.RecipientId,
			Recipient:       ToResponseUserButNotPointer(item.User),
			TransactionTime: item.Transaction.TransactionTime,
			Amount:          item.Transaction.Amount,
			Description:     item.Transaction.Description,
			CreatedAt:       item.Transaction.CreatedAt,
			UpdatedAt:       item.Transaction.UpdatedAt,
			DeleteAt:        item.Transaction.DeleteAt,
		})
	}
	return res
}

func toSourceOfFundDTO(source models.SourceOfFund) dtos.SourceOfFund {
	return dtos.SourceOfFund{
		ID:        source.ID,
		Name:      source.Name,
		CreatedAt: source.CreatedAt,
		UpdatedAt: source.UpdatedAt,
		DeleteAt:  source.DeleteAt,
	}
}
