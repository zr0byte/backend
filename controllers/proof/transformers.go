package proof

import (
	"0byte/models"
	"0byte/services/proofsvc"
)

func generateProofTransformer(baseRes models.BaseResponse, res proofsvc.ZKProofResponse) models.BaseResponse {
	var finalRes models.BaseResponse

	var data proofsvc.ZKProofResponse
	data.Proof = res.Proof
	data.ReceiverCommit = res.ReceiverCommit
	data.SenderCommit = res.SenderCommit
	data.AmountCommit = res.AmountCommit

	finalRes.Data = data
	finalRes.Success = baseRes.Success
	finalRes.StatusCode = baseRes.StatusCode
	finalRes.Message = baseRes.Message

	return baseRes
}
