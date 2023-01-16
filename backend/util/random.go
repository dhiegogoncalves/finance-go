package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const lettersAndNumbers = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(number int) string {
	var sb strings.Builder
	k := len(lettersAndNumbers)

	for i := 0; i < number; i++ {
		c := lettersAndNumbers[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomEmail(number int) string {
	return fmt.Sprintf("%s@email.com", RandomString(number))
}
