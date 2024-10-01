package proof

import "0byte/services/proofsvc"

type Proofhandler struct {
	proofSvc proofsvc.Interface
}

func Handler(proofSvc proofsvc.Interface) *Proofhandler {
	return &Proofhandler{
		proofSvc: proofSvc,
	}
}
