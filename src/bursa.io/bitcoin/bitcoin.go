package backend

import (
	// "strings"

	// "github.com/conformal/btcrpcclient"
)

// Not quite sure what we want here yet
type TransferResponse struct {
	msg  string
	code int
}

type CurrencyBackend interface {
	ExecuteTransfer(transfer *Transfer) *TransferResponse
}

type BitcoinBackend struct {
}
