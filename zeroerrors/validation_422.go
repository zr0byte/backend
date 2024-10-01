package zeroerrors

import (
	"0byte/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

/* -------------------------------------------------------------------------- */
/*                              Validation errors                             */
/* -------------------------------------------------------------------------- */
func Validation(ctx *gin.Context, err string) {
	var res models.BaseResponse
	var zeroerror Error
	errorCode := http.StatusUnprocessableEntity

	zeroerror.Code = errorCode
	zeroerror.Type = "validation"
	zeroerror.Message = err

	res.Error = zeroerror

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
