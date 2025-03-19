package function

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"os"
	"time"
)

func GenerateReferenceCode() string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 4)
	for i := range b {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		b[i] = letters[randomIndex.Int64()]
	}
	return string(b)
}

func GenerateToken() (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("missing secret key")
	}

	expirationTime := time.Now().Add(24 * time.Hour).Unix()
	expirationHex := fmt.Sprintf("%x", expirationTime)

	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(expirationHex))
	signature := h.Sum(nil)

	shortSignature := hex.EncodeToString(signature[:12])

	token := fmt.Sprintf("%s:%s", expirationHex, shortSignature)

	return base64.RawURLEncoding.EncodeToString([]byte(token)), nil
}
