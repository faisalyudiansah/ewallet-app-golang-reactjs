package httphandler

import (
	"errors"
	"net/http"

	"ewallet-server-v2/internal/dto/httpdto"
	"ewallet-server-v2/internal/pkg/apperror"
	"ewallet-server-v2/internal/pkg/apputils"
	"ewallet-server-v2/internal/pkg/ctxutils"
	"ewallet-server-v2/internal/pkg/ginutils"
	"ewallet-server-v2/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase   usecase.UserUsecase
	walletUsecase usecase.WalletUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase, walletUsecase usecase.WalletUsecase) *UserHandler {
	return &UserHandler{
		userUsecase:   userUsecase,
		walletUsecase: walletUsecase,
	}
}

func (h *UserHandler) GetUserDetails(c *gin.Context) {
	var res httpdto.GetUserDetailsResponse

	userId := ctxutils.GetUserId(c)

	user, err := h.userUsecase.GetOneById(c, userId)
	if err != nil {
		c.Error(err)
		return
	}

	wallet, err := h.walletUsecase.GetOneByUserId(c, userId)
	if err != nil {
		c.Error(err)
		return
	}

	res = httpdto.ConvertToGetUserDetailResponse(user, wallet)

	ginutils.ResponseOKData(c, res)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	var res httpdto.UpdateProfileResponse
	var req httpdto.UpdateProfileRequest

	err := c.BindJSON(&req)
	if err != nil {
		c.Error(err)
		return
	}

	if req.FullName != "" && (len(req.FullName) < 3 || len(req.FullName) > 36) {
		c.Error(apperror.NewAppError(
			errors.New("full name must be between 3 until 36 characters long"),
			int(http.StatusBadRequest),
			"full name must be between 3 until 36 characters long",
			nil,
		))
		return
	}

	if req.Email != "" && !apputils.IsValidEmail(req.Email) {
		c.Error(apperror.NewAppError(
			errors.New("invalid email address"),
			int(http.StatusBadRequest),
			"invalid email address",
			nil,
		))
		return
	}

	if req.ProfileImage != "" && !apputils.IsValidURL(req.ProfileImage) {
		c.Error(apperror.NewAppError(
			errors.New("invalid profile image url"),
			int(http.StatusBadRequest),
			"invalid profile image url",
			nil,
		))
		return
	}

	userId := ctxutils.GetUserId(c)

	user, err := h.userUsecase.GetOneById(c, userId)
	if err != nil {
		c.Error(err)
		return
	}

	user.Email = apputils.GetStringValueOrDefault(req.Email, user.Email)
	user.FullName = apputils.GetStringValueOrDefault(req.FullName, user.FullName)
	user.ProfileImage = apputils.GetStringValueOrDefault(req.ProfileImage, user.ProfileImage)

	updatedUser, err := h.userUsecase.UpdateOne(c, userId, *user)
	if err != nil {
		c.Error(err)
		return
	}

	res = httpdto.ConvertToUpdateProfileResponse(updatedUser)

	ginutils.ResponseOKData(c, res)
}
