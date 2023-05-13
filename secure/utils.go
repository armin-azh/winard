package secure

import "golang.org/x/crypto/bcrypt"

// ComparePassword Compare a given password with relative hashed value
func ComparePassword(payload string, password string) (bool, error) {

	if err := bcrypt.CompareHashAndPassword([]byte(payload), []byte(password)); err != nil {
		return false, err
	}
	return true, nil
}
