package proofsvc

import (
	"github.com/consensys/gnark/backend/groth16"
)

type ProofRequestObject struct {
	SendersAddress   string  `json:"senders_address"`
	ReceiversAddress string  `json:"recievers_address"`
	Amount           float64 `json:"amount"`
	SendersBalance   uint64  `json:"senders_balance"`
	Message          string  `json:"message"`
}

type ProofResponse struct {
	Proof string `json:"proof"`
}

// ZKProofResponse represents the proof and commitments returned in response
type ZKProofResponse struct {
	Proof          groth16.Proof `json:"proof"`
	SenderCommit   string        `json:"sender_commitment"`
	ReceiverCommit string        `json:"receiver_commitment"`
	AmountCommit   string        `json:"amount_commitment"`
}

type Commitments struct {
	SenderCommit   string `json:"sender_commitment"`
	ReceiverCommit string `json:"receiver_commitment"`
	AmountCommit   string `json:"amount_commitment"`
}
