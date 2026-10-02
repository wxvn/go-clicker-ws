package hasher

import "golang.org/x/crypto/bcrypt"

const bcryptCost = bcrypt.DefaultCost

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(password, hash string) error
}

type Bcrypt struct {
	cost int
}

func NewBcrypt() *Bcrypt {
	return &Bcrypt{
		cost: bcryptCost,
	}
}

func (h *Bcrypt) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (h *Bcrypt) Compare(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
