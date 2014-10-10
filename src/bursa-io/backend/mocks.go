package backend

type MockTransfer struct {
	*Transfer
	// causes MockBackend to return a failure if false
	success bool
}

// type MockWallet struct {
//   *Wallet
// }
// The above doesn't really Mock a wallet. The Wallet is an anonymous embedded
// object inside of MockWallet but go does not treat it as an instanceof Wallet
// AFAICT.
//
// My gut says that go favours interfaces and has no notion of inheritance.
// So how do we create a mock that has the same attributes as Wallet and can be 
// substituted for a real wallet without forcing us to turn every public attribute
// of Wallet into a function (so that we can code against an interface).
//
// Very possibly - you can't? When you say that a parameter takes a Wallet
// that parameter can only ever be a Wallet. Since there may not be inheritance in
// go we have no way of extending an existing object.

type MockBackend struct {
}

// type MockBackend struct {
//   *CurrencyBackend
// }

// This (above) doesn't work. CurrencyBackend is an interface, which doesn't have
// a concrete size at runtime - different implementors of the interface have
// different sizes. If go structs are like c structs, then they have the characteristic
// of having a fairly consistent memory layout - e.g.
//
// type SomeStruct struct {
//  x int32
//  y int32
// }
// ptr := new(SomeStruct)
// *(ptr + 4) - first byte of y. go doesn' actually allow pointer arithmetic.

// Define Mock Backend
func (b *MockBackend) ExecuteTransfer(t *Transfer) *TransferResponse {

  response := new(TransferResponse)

	if t.IsSuccess() {
		response.msg = "Transfer succeeded"
		response.code = 200
	} else {
		response.msg = "Transfer failed"
		response.code = 500
	}

	return response
}

// Not defined in the interface
// func (b *MockBackend) ProvisionWallet() (*Wallet)  {
//   wallet := new(MockWallet)
//   return wallet;
// }
