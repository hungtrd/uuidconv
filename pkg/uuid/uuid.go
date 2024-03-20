package uuid

import (
	"encoding/base64"
	"errors"
	"fmt"

	guuid "github.com/google/uuid"
	"github.com/hungtrd/uuidconv/pkg/base62"
)

type Format string

var (
	FormatString Format = "string"
	FormatBase64 Format = "base64"
	FormatBase62 Format = "base62"
)

type UUID struct {
	NormalString  string
	Base64Encoded string
	Base62Encoded string
}

func New() UUID {
	u := guuid.New()
	return NewWithUUID(u)
}

func NewWithUUID(u guuid.UUID) UUID {
	b64 := base64.StdEncoding.EncodeToString(u[:])
	b62 := base62.EncodeBytes(u[:])

	return UUID{
		NormalString:  u.String(),
		Base64Encoded: b64,
		Base62Encoded: b62,
	}
}

func NewFromUnknow(s string) (UUID, error) {
	if u, err := NewFromString(s); err == nil {
		fmt.Println("String format was detected")
		return u, nil
	}

	if u, err := NewFromBase64(s); err == nil {
		fmt.Println("Base64 format was detected")
		return u, nil
	}

	if u, err := NewFromBase62(s); err == nil {
		fmt.Println("Base62 format was detected")
		return u, nil
	}

	return UUID{}, errors.New("unknown input format")
}

func NewFromString(s string) (UUID, error) {
	u, err := guuid.Parse(s)
	if err != nil {
		return UUID{}, err
	}

	return NewWithUUID(u), nil
}

func NewFromBase64(s string) (UUID, error) {
	b64, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return UUID{}, err
	}
	u, err := guuid.FromBytes(b64)
	if err != nil {
		return UUID{}, err
	}

	return NewWithUUID(u), nil
}

func NewFromBase62(s string) (UUID, error) {
	b := base62.DecodeStr(s)
	u, err := guuid.FromBytes(b)
	if err != nil {
		return UUID{}, err
	}

	return NewWithUUID(u), nil
}
