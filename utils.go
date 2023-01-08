package core

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func RandomToken(source interface{}) (string, error) {
	randNumber := rand.Intn(9999999 - 1000000) + 1000000
	hashed, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%v", source)), 10)
	if err != nil {
		return "", err
	}

	hashed = []byte(strings.Replace(string(hashed), ".", "", -1))
	token := fmt.Sprintf("%d-%s", randNumber, hashed)
	

	return token, nil
}
func GetHashPassword(localPassword, userPassword string) (string) {
	return fmt.Sprintf("%s%s", localPassword, userPassword)
}

func ParseJSON(v interface{}) (interface{}, error) {
	i, err := json.Marshal(v)
	if err != nil {
		return i, err
	}

	return i, nil
}

func UnParseJSON(str string, v interface{}) error {
	return json.Unmarshal([]byte(str), &v)
}