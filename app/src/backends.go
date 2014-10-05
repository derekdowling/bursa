package backends

import "strings"

// Not quite sure what we want here yet
type TransferResponse struct {
	msg  string
	code int
}

type CurrencyBackend interface {
	executeTransfer(transfer *Transfer) *TransferResponse
}
