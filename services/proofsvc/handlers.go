package proofsvc

import (
	"0byte/models"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                              Service Interface                             */
/* -------------------------------------------------------------------------- */
type Interface interface {
	GenerateProof(ctx *gin.Context, req ProofRequestObject) (models.BaseResponse, ZKProofResponse, error)
}

/* -------------------------------------------------------------------------- */
/*                                  Reciever                                  */
/* -------------------------------------------------------------------------- */
type proofSvcImpl struct{}

/* -------------------------------------------------------------------------- */
/*                                   Handler                                  */
/* -------------------------------------------------------------------------- */
func Handler() *proofSvcImpl {
	return &proofSvcImpl{}
}
