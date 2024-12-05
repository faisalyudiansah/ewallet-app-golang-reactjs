package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"ewallet-server-v2/internal/pkg/apperror"
	"ewallet-server-v2/internal/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// logging the request information such as:
// - method
// - status code
// - request processing latency
// - path
// some code taken from the default logger gin code
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		param := map[string]interface{}{
			"status_code": c.Writer.Status(),
			"method":      c.Request.Method,
			"latency":     time.Since(start),
			"path":        path,
		}

		if len(c.Errors) == 0 {
			logger.Log.WithFields(param).Info("incoming request")
		} else {
			errList := []error{}
			for _, err := range c.Errors {
				var appErr apperror.AppError
				var vErr validator.ValidationErrors
				var utErr *json.UnmarshalTypeError
				var sErr *json.SyntaxError

				switch {
				case errors.As(err, &sErr):
					fallthrough
				case errors.As(err, &utErr):
					fallthrough
				case errors.As(err, &vErr):
					param["status_code"] = http.StatusBadRequest
					errList = append(errList, err)
				case errors.As(err, &appErr):
					param["status_code"] = codeMap[appErr.GetCode()]
					errList = append(errList, appErr.OriginalError())

					if datamap := appErr.AdditionalData(); datamap != nil {
						for k, v := range datamap {
							param[k] = v
						}
					}
				default:
					param["status_code"] = http.StatusInternalServerError
					errList = append(errList, err)
				}
			}

			if len(errList) > 0 {
				param["errors"] = errList
				logger.Log.WithFields(param).Error("got error")
			}
		}
	}
}
