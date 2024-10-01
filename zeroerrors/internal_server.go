package zeroerrors

import (
	"0byte/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                            Internal server error                           */
/* -------------------------------------------------------------------------- */
func InternalServer(ctx *gin.Context, err string) {
	var res models.BaseResponse
	var zeroerror Error
	errorCode := http.StatusInternalServerError

	zeroerror.Code = errorCode
	zeroerror.Type = "server"
	zeroerror.Message = err

	res.Error = zeroerror

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
