package proof

import (
	"0byte/services/proofsvc"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

/* -------------------------------------------------------------------------- */
/*                           Validate Proof Request                           */
/* -------------------------------------------------------------------------- */
func validateProofRequest(ctx *gin.Context) (proofsvc.ProofRequestObject, error) {
	var req proofsvc.ProofRequestObject
	var err error

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		return req, err
	}
	validate := validator.New()

	err = validate.Struct(req)
	if err != nil {
		return req, err
	}

	// validate request body
	err = validateReqBody(req)
	if err != nil {
		return req, err
	}

	// convert sol to lamports
	// req.Amount = solToLamports(req.Amount)

	return req, err
}

func validateReqBody(req proofsvc.ProofRequestObject) error {
	if req.Amount <= 0 {
		return errors.New("amount should be greater than 0")
	}

	return nil
}
