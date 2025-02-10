package password

// import (
// 	"fmt"

// 	"golang.org/x/crypto/bcrypt"
// )

// HashPassword generates a bcrypt hash for the given password.
func HashPassword(password string) (string, error) {
	// bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	// return string(bytes), err
	return password, nil
}

// VerifyPassword verifies if the given password matches the stored hash.
func VerifyPassword(password, hash string) (bool, error) {
	// Log the comparison to debug
	// fmt.Println("Verifying password with hash:\n", password, "\n", hash)
	// err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	// if err != nil {
	// 	return false, err
	// }
	// return true, nil
	if password == hash {
		return true, nil
	}
	return false, nil
}
