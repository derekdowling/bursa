package models

import (
  "time"
)

const (
  BitcoinCurrency = "bitcoin"
)

const (
  CompletedStatus = "completed"
  PendingStatus = "pending"
  FailedStatus = "failed"
  RolledBackStatus = "rolledback"
)

type User struct {
  Id int64
  Name string `sql:"size:255"`
  CreatedAt time.Time
  UpdatedAt time.Time

  Email `sql:"size:255"`
}

type Transfer struct {
  // TODO Refer to bitcoin protocol.
  Id string `sql:"size:255"`

  // Amount in satoshis when using Bitcoin.
  Currency string `sql:"size:255"`
  Amount int64

  SourceWallet Wallet
  DestinationWallet Wallet

  CreatedBy User
  CreatedAt time.Time
  CompletedAt time.Time
  Status string `sql:"size:32"`
}

type Wallet struct {
  Id string `sql:"size:255"`
  Name string `sql:"size:255"`
  Balance int64
  Transfers []Transfer

  CreatedAt time.Time
  CreatedBy User
  OwnedBy User
}
