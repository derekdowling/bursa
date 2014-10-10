package backend

import (
	// "strings"
	"testing"
  "bursa-io/backend"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	Convey("Mocks Testing", t, func() {
    backend.Initialize()

		// backend := new(MockBackend)
		// transfer := new(MockTransfer)
		// mockWallet := new(MockWallet)

		// Convey("MockBackend Type", func() {
		//   So(backend, ShouldHaveSameTypeAs, *CurrencyBackend)
		// })

		// Convey("MockTransfer Type", func() {
		//   So(transfer, ShouldHaveSameTypeAs, *Transfer)
		// })

		// Convey("MockWallet Type", func() {
		//   So(wallet, ShouldHaveSameTypeAs, *Wallet)
		// })

		// Convey("MockBackend Functions", func() {

		//   Convey(".executeTransfer()", func() {
		//     Convey("if successful", func() {
		//       transfer.success = true
		//       response = backend.executeTransfer(transfer)

		//       So(response, ShouldHaveSameTypeAs, *ResponseTransfer)
		//       So(response.code, ShouldEqual, 200)
		//     })
		//     Convey("if failed", func() {
		//       transfer.success = false
		//       response = backend.executeTransfer(transfer)

		//       So(response, ShouldHaveSameTypeAs, *ResponseTransfer)
		//       So(response.code, ShouldEqual, 500)
		//     })
		//   })

		//   Convey(".provisionWallet()", func() {
		//     Convey("if successful", func() {
		//       wallet = backend.provisionWallet()

		//       So(wallet, ShouldHaveSameTypeAs, *mockWallet)
		//       So(wallet.balance, ShouldEqual, 0)
		//     })
		//     Convey("if failure", func() {
		//       wallet = backend.provisionWallet()

		//       So(wallet, ShouldHaveSameTypeAs, null)
		//     })
		//   })
		// })
	})
}
