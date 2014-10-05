package tests

import (
	"strings"
	"testing"

	. "bursa/backends"
)

type MockBackend struct{}

type MockTransfer struct {
	success bool
}

// Define Mock Backend
func (b *CurrencyBackend) executeTransfer(t *Transfer) *TransferResponse {

	*response = new(TransferResponse)

	if t.success {
		response.msg = "Transfer succeeded"
		response.code = 200
	} else {
		response.msg = "Transfer failed"
		response.code = 500
	}

	return response
}

func TestSpec(t *testing.T) {

	Convey("MockBackend", t, func() {
		backend := new(CurrencyBackend)

		Convey("executeTransfer", func() {

			Convey("onSuccess", func() {
				transfer = new(MockTransfer)
				transfer.success = true

				response = backend.executeTransfer(transfer)

				So(response, ShouldHaveSameTypeAs, *ReponseTransfer)
				So(response.code, ShouldEqual, 200)
			})
		})
	})
}
