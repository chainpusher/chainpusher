package wallet

type PoolRepository interface {
	FindByAccountId(accountId int64) (*Pool, error)
}
