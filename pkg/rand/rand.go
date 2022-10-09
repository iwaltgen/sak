// Copyright (c) 2022 iwaltgen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package rand

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"

	"github.com/hashicorp/go-uuid"
	"github.com/jxskiss/base62"
)

const defaultLength = 16

var (
	base64Encoder = base64.RawURLEncoding
	base32Encoder = base32.StdEncoding.WithPadding(base32.NoPadding)
)

// New creates a new random []byte or panics.
func New() []byte {
	return must(GenerateRandomBytes())
}

// NewBase64 creates a new random []byte and returns it as a base64 string or panics.
func NewBase64() string {
	return base64Encoder.EncodeToString(must(GenerateRandomBytes()))
}

// NewBase62 creates a new random []byte and returns it as a base62 string or panics.
func NewBase62() string {
	return base62.EncodeToString(must(GenerateRandomBytes()))
}

// NewBase32 creates a new random []byte and returns it as a base32 string or panics.
func NewBase32() string {
	return base32Encoder.EncodeToString(must(GenerateRandomBytes()))
}

// NewHex creates a new random []byte and returns it as a hex string or panics.
func NewHex() string {
	return hex.EncodeToString(must(GenerateRandomBytes()))
}

// NewWithLength creates a new random []byte or panics.
func NewWithLength(length int) []byte {
	return must(uuid.GenerateRandomBytes(length))
}

// NewBase64WithLength creates a new random []byte and returns it as a base64 string or panics.
func NewBase64WithLength(length int) string {
	return base64Encoder.EncodeToString(must(uuid.GenerateRandomBytes(length)))
}

// NewBase62WithLength creates a new random []byte and returns it as a base62 string or panics.
func NewBase62WithLength(length int) string {
	return base62.EncodeToString(must(uuid.GenerateRandomBytes(length)))
}

// NewBase32WithLength creates a new random []byte and returns it as a base32 string or panics.
func NewBase32WithLength(length int) string {
	return base32Encoder.EncodeToString(must(uuid.GenerateRandomBytes(length)))
}

// NewHexWithLength creates a new random []byte and returns it as a hex string or panics.
func NewHexWithLength(length int) string {
	return hex.EncodeToString(must(uuid.GenerateRandomBytes(length)))
}

// GenerateRandomBytes returns a random []byte.
func GenerateRandomBytes() ([]byte, error) {
	return uuid.GenerateRandomBytes(defaultLength)
}

func must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
