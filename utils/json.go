package utils

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func ReturnJsonStruct(ctx *gin.Context, genericStruct interface{}) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(ctx.Writer).Encode(genericStruct)
	if err != nil {
		return
	}
}
