package wallet_interface

type BlockSubscription struct {
	Out   chan BlockInfo
	Close func()
}

type TransactionSubscription struct {
	Out         chan Transaction
	Subscribe   chan []Address
	Unsubscribe chan []Address
	Close       func()
}

type ChainClient interface {
	GetBlockchainInfo() (BlockInfo, error)

	GetAddressTransactions(addr Address, fromHeight uint64) ([]Transaction, error)

	GetTransaction(id TransactionID) (Transaction, error)

	IsBlockInMainChain(block BlockInfo) (bool, error)

	SubscribeTransactions(addrs []Address) (*TransactionSubscription, error)

	SubscribeBlocks() (*BlockSubscription, error)

	Broadcast(serializedTx []byte) error

	Open() error

	Close() error
}
