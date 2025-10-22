package services

type HashService interface {
	Hash(password string) (string, error)
	Compare(hashed string, plain string) error
}
