package proofsvc

import (
	"0byte/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

/* -------------------------------------------------------------------------- */
/*                               generate Proof                               */
/* -------------------------------------------------------------------------- */
func (h *proofSvcImpl) GenerateProof(ctx *gin.Context, req ProofRequestObject) (models.BaseResponse, ZKProofResponse, error) {
	var baseRes models.BaseResponse
	var res ZKProofResponse
	var err error

	// default response
	baseRes.StatusCode = http.StatusInternalServerError
	baseRes.Message = "something went wrong"

	// get wallet balance
	walletResponse, err := getWalletBalance(req.SendersAddress)
	if err != nil {
		return baseRes, res, errors.Wrap(err, "[GenerateProof][GetWalletBalance]")
	}

	req.SendersBalance = walletResponse.Lamports

	// generate Proof
	proof, err := generateZKProof(req)
	if err != nil {
		return baseRes, res, errors.Wrap(err, "[generateProof][generateProof]")
	}

	// generate Commitments
	commiments, err := generateCommitments(req)
	if err != nil {
		return baseRes, res, errors.Wrap(err, "[generateProof][generateCommitments]")
	}

	// map the response
	res.Proof = proof
	res.AmountCommit = commiments.AmountCommit
	res.ReceiverCommit = commiments.ReceiverCommit
	res.SenderCommit = commiments.SenderCommit

	// success response
	baseRes.Success = true
	baseRes.StatusCode = http.StatusOK
	baseRes.Message = "proof generated succesfully"

	return baseRes, res, err
}
