package models

type UserAndWallet struct {
	User   User
	Wallet Wallet
}

type TransactionUserAndSourceOfFund struct {
	Transaction  Transaction
	User         User
	SourceOfFund SourceOfFund
}
