package function

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"
)

func ValidateToken(token string) (bool, error) {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return false, errors.New("missing secret key")
	}

	decoded, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil {
		return false, errors.New("invalid token format")
	}

	parts := strings.Split(string(decoded), ":")
	if len(parts) != 2 {
		return false, errors.New("invalid token structure")
	}

	expirationHex, providedSignature := parts[0], parts[1]

	expirationTime, err := strconv.ParseInt(expirationHex, 16, 64)
	if err != nil || time.Now().Unix() > expirationTime {
		return false, errors.New("token expired")
	}

	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(expirationHex))
	expectedSignature := hex.EncodeToString(h.Sum(nil)[:12])

	if !hmac.Equal([]byte(providedSignature), []byte(expectedSignature)) {
		return false, errors.New("invalid token signature")
	}

	return true, nil
}
