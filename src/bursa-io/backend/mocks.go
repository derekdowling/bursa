package backend

type MockTransfer struct {
	*Transfer
	// causes MockBackend to return a failure if false
	success bool
}

type MockWallet struct {
	*Wallet
}

type MockBackend struct {
	*CurrencyBackend
}

// Define Mock Backend
func (b *MockBackend) executeTransfer(t *Transfer) *TransferResponse {

	response = new(TransferResponse)

	if t.success {
		response.msg = "Transfer succeeded"
		response.code = 200
	} else {
		response.msg = "Transfer failed"
		response.code = 500
	}

	return response
}

func (b *MockBackend) provisionWallet() *Wallet {
	wallet := new(MockWallet)
}
