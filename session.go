package core

import (
	"fmt"
	"strings"
	"strconv"
)

func UnParseSession(s string) (Session, error) {
	fragments := strings.Split(s, ".")
	if len(fragments) != 2 {
		return Session{}, fmt.Errorf("invalid session")
	}

	id := fragments[0]
	expiration, err := strconv.Atoi(fragments[1])
	if err != nil {
		return Session{}, fmt.Errorf("invalid session")
	}

	return Session{
		ID:             id,
		Expiration: int64(expiration),
	}, nil
}

func ParseSession(s Session) string {
	return fmt.Sprintf("%s.%d", s.ID, s.Expiration)
}