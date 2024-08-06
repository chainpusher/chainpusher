package account

type Repository interface {
	FindBySecretKey(secretKey string) (*Account, error)
}
