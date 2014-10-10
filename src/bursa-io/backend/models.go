package backend

import (
	"github.com/jinzhu/gorm"
	"time"
)

const (
	BitcoinCurrency = "bitcoin"
)

const (
	CompletedStatus  = "completed"
	PendingStatus    = "pending"
	FailedStatus     = "failed"
	RolledBackStatus = "rolledback"
)

const (
	CodeSuccess = 200
)

type User struct {
	Id        int64
	Name      string `sql:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Email string `sql:"size:255"`
}

type Transfer struct {
	// TODO Refer to bitcoin protocol.
	Id string `sql:"size:255"`

	// Amount in satoshis when using Bitcoin.
	Currency string `sql:"size:255"`
	Amount   int64

	SourceWallet      Wallet
	DestinationWallet Wallet

	CreatedBy   User
	CreatedAt   time.Time
	CompletedAt time.Time
	Status      string `sql:"size:32"`
	Code        int16 `sql:"size:8"`
}

func (self *Transfer) IsSuccess() bool {
  return self.Code == CodeSuccess
}

type Wallet struct {
	Id        string `sql:"size:255"`
	Name      string `sql:"size:255"`
	Pin       string `sql:"size:255"`
	Balance   int64
	Transfers []Transfer

	CreatedAt time.Time
	CreatedBy User
	OwnedBy   User
}

// Well suited to some kind of management cli.
func Initialize() {
	db, _ := gorm.Open("postgres", "user=bursa dbname=bursa sslmode=disable")
	db.CreateTable(User{})
	db.CreateTable(Transfer{})
	db.CreateTable(Wallet{})
}
