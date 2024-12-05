package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"ewallet-server-v2/internal/constant"
	"ewallet-server-v2/internal/dto/httpdto"
	"ewallet-server-v2/internal/pkg/apperror"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var codeMap = map[int]int{
	apperror.DefaultClientErrorCode: http.StatusBadRequest,
	apperror.DefaultServerErrorCode: http.StatusInternalServerError,
	apperror.NotFoundErrorCode:      http.StatusNotFound,
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		errLen := len(c.Errors)
		if errLen > 0 {
			// get only the last error
			err := c.Errors[len(c.Errors)-1]

			var appErr apperror.AppError
			var vErr validator.ValidationErrors
			var utErr *json.UnmarshalTypeError
			var sErr *json.SyntaxError
			var pErr *time.ParseError

			switch {
			case errors.As(err, &sErr):
				handleJsonSyntaxError(c, sErr)
				return
			case errors.As(err, &utErr):
				handleJsonUnmarshalTypeError(c, utErr)
				return
			case errors.As(err, &pErr):
				handleParseTimeError(c, pErr)
				return
			case errors.As(err, &vErr):
				handleValidationError(c, vErr)
			case errors.As(err, &appErr):
				c.AbortWithStatusJSON(codeMap[appErr.GetCode()], httpdto.ErrorResponse{
					Message: appErr.DisplayMessage(),
				})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, httpdto.ErrorResponse{
					Message: constant.InternalServerErrorMessage,
				})
			}
		}
	}
}

func handleJsonSyntaxError(c *gin.Context, err *json.SyntaxError) {
	c.AbortWithStatusJSON(http.StatusBadRequest, httpdto.ErrorResponse{
		Message: constant.JsonSyntaxError,
	})
}

func handleJsonUnmarshalTypeError(c *gin.Context, err *json.UnmarshalTypeError) {
	c.AbortWithStatusJSON(http.StatusBadRequest, httpdto.ErrorResponse{
		Message: fmt.Sprintf(constant.InvalidJsonValueTypeError, err.Field),
	})
}

func handleParseTimeError(c *gin.Context, err *time.ParseError) {
	c.AbortWithStatusJSON(http.StatusBadRequest, httpdto.ErrorResponse{
		Message: fmt.Sprintf("please send time in format of %s, got: %s", constant.ConvertGoTimeLayoutToReadable(err.Layout), err.Value),
	})
}

func handleValidationError(c *gin.Context, err validator.ValidationErrors) {
	ve := []httpdto.ValidationErrorResponse{}

	for _, fe := range err {
		ve = append(ve, httpdto.ValidationErrorResponse{
			Field:   fe.Field(),
			Message: tagToMsg(fe),
		})
	}

	c.AbortWithStatusJSON(http.StatusBadRequest, httpdto.ErrorResponse{
		Message: constant.ValidationError,
		Details: ve,
	})
}

// instead of cryptic error given by the validator, we overwrite and define the error message by ourself
func tagToMsg(fe validator.FieldError) string {
	switch fe.Tag() { // add more as needed
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "len":
		return fmt.Sprintf("%s must be %v characters", fe.Field(), fe.Param())
	case "max":
		return fmt.Sprintf("%s must not exceed %v characters", fe.Field(), fe.Param())
	case "dgte":
		return fmt.Sprintf("%s must be greater than or equal to %v", fe.Field(), fe.Param())
	case "dlte":
		return fmt.Sprintf("%s must be less than or equal to %v", fe.Field(), fe.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %v", fe.Field(), fe.Param())
	case "lte":
		return fmt.Sprintf("%s must be lower than or equal to %v", fe.Field(), fe.Param())
	case "email":
		return fmt.Sprintf("%s has invalid email format", fe.Field())
	case "eq":
		return fmt.Sprintf("%s must be: %v", fe.Field(), fe.Param())
	case "min":
		return fmt.Sprintf("%s must be %v characters long", fe.Field(), fe.Param())
	default:
		return "invalid input"
	}
}
