package servers

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-04/faisal.yudiansah/assignment-e-wallet-rest-api/apperrors"
	"github.com/gin-gonic/gin"
)

func InvalidRoute(c *gin.Context) {
	c.Error(apperrors.ErrUrlNotFound)
	c.Abort()
}
