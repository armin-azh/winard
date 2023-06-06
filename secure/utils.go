package secure

import "golang.org/x/crypto/bcrypt"

// ComparePassword Compare a given password with relative hashed value
func ComparePassword(payload *string, password *string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(*payload), []byte(*password)); err != nil {
		return false, err
	}
	return true, nil
}

// CipherPassword generate hash from password
func CipherPassword(password *string, secret *[]byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
