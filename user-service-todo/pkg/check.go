package pkg


import "golang.org/x/crypto/bcrypt"

func CheckPassword(password, hash string) bool {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}