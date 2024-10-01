package proofsvc

import "github.com/bwesterb/go-ristretto"

func generateCommitments(req ProofRequestObject) (Commitments, error) {
	var commitments Commitments

	// Convert the amount to big.Int
	amount := floatToBigInt(req.Amount)

	// Generate random scalars for sender, receiver, and amount
	rSender := new(ristretto.Scalar)
	rReceiver := new(ristretto.Scalar)
	rAmount := new(ristretto.Scalar)
	rSender.Rand()
	rReceiver.Rand()
	rAmount.Rand()

	// Generate H, a random point used in the commitment
	H := generateH()

	// Set the amount as a scalar
	var senderAmountScalar, receiverAmountScalar ristretto.Scalar
	senderAmountScalar.SetBigInt(amount)
	receiverAmountScalar.SetBigInt(amount)

	// Commit to sender and receiver using the H point
	senderCommit := commitTo(&H, rSender, &senderAmountScalar)
	receiverCommit := commitTo(&H, rReceiver, &receiverAmountScalar)
	amountCommit := commitTo(&H, rAmount, &senderAmountScalar)

	// Convert commitments to strings (using proper encoding)
	commitments.SenderCommit = senderCommit.String()
	commitments.ReceiverCommit = receiverCommit.String()
	commitments.AmountCommit = amountCommit.String()

	return commitments, nil
}
