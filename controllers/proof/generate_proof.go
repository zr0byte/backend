package proof

import (
	"0byte/zeroerrors"
	"fmt"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                               Generate Proof                               */
/* -------------------------------------------------------------------------- */
func (h *Proofhandler) GenerateProof(ctx *gin.Context) {
	req, err := validateProofRequest(ctx)
	fmt.Println("validation error : ", err)
	if err != nil {
		zeroerrors.Validation(ctx, err.Error())
		return
	}

	baseRes, res, err := h.proofSvc.GenerateProof(ctx, req)
	fmt.Println("server error : ", err)
	if err != nil {
		zeroerrors.InternalServer(ctx, err.Error())
		return
	}
	ctx.JSON(200, gin.H{
		"success":     baseRes.Success,
		"status_code": baseRes.StatusCode,
		"data":        res,
		"message":     baseRes.Message,
	})

}
