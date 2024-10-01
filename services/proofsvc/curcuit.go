package proofsvc

import (
	"errors"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

type BalanceCircuit struct {
	Balance frontend.Variable `gnark:",public"`  // Public input (sender's balance)
	Amount  frontend.Variable `gnark:",private"` // Private input (transaction amount)
}

func (circuit *BalanceCircuit) Define(api frontend.API) error {
	api.AssertIsLessOrEqual(circuit.Amount, circuit.Balance)
	return nil
}

func floatToBigInt(f float64) *big.Int {
	scaled := new(big.Float).SetFloat64(f)
	scalingFactor := new(big.Float).SetFloat64(1e6) // 6 decimal places
	scaled.Mul(scaled, scalingFactor)
	result := new(big.Int)
	scaled.Int(result) // truncate to an integer
	return result
}

func generateZKProof(req ProofRequestObject) (groth16.Proof, error) {
	var proof groth16.Proof

	// Convert float64 amounts to big.Int for precise arithmetic
	senderBalance := uint64ToBigInt(req.SendersBalance)
	amount := solToLamports(req.Amount)

	// Check if the sender has enough balance
	if senderBalance.Cmp(amount) < 0 {
		return proof, errors.New("insufficient balance")
	}

	// Define the circuit with the balance and amount as variables
	circuit := BalanceCircuit{}

	// Compile the circuit to R1CS (Rank-1 Constraint System)
	r1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	if err != nil {
		return proof, err
	}

	// Generate the proving and verifying keys
	pk, vk, err := groth16.Setup(r1cs)
	if err != nil {
		return proof, err
	}

	assignment := BalanceCircuit{
		Balance: senderBalance,
		Amount:  amount,
	}

	// Create witness (inputs for the circuit)
	witness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	if err != nil {
		return proof, err
	}

	// Generate the zkSNARK proof
	proof, err = groth16.Prove(r1cs, pk, witness)
	if err != nil {
		return proof, err
	}

	// Verify the proof
	publicWitness, err := witness.Public()
	if err != nil {
		return proof, err
	}

	err = groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		return proof, err
	}

	return proof, nil
}

func solToLamports(sol float64) *big.Int {
	// Create a big.Float to represent SOL
	solAmount := new(big.Float).SetFloat64(sol)

	// Define 1 SOL = 1,000,000,000 lamports
	lamportsPerSol := new(big.Float).SetFloat64(1e9)

	// Multiply SOL by 1e9 to get lamports
	lamports := new(big.Float).Mul(solAmount, lamportsPerSol)

	// Convert the result to big.Int for precision
	lamportsInt := new(big.Int)
	lamports.Int(lamportsInt) // Truncate to integer
	return lamportsInt
}

func uint64ToBigInt(n uint64) *big.Int {
	return new(big.Int).SetUint64(n)
}
