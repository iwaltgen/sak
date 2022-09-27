package rand

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"io"
)

const defaultLength = 32

var defaultRander = rand.Reader

// New creates a new random []byte or panics.
func New() []byte {
	return must(NewRandom())
}

// NewBase64 creates a new random []byte and returns it as a base64 string or panics.
func NewBase64() string {
	return base64.URLEncoding.EncodeToString(must(NewRandom()))
}

// NewBase32 creates a new random []byte and returns it as a base32 string or panics.
func NewBase32() string {
	return base32.StdEncoding.EncodeToString(must(NewRandom()))
}

// NewHex creates a new random []byte and returns it as a hex string or panics.
func NewHex() string {
	return hex.EncodeToString(must(NewRandom()))
}

// NewRandom returns a random []byte.
// The strength of the []bytes is based on the strength of the crypto/rand package.
func NewRandom() ([]byte, error) {
	return NewRandomFromReader(defaultRander, defaultLength)
}

// NewWithLength creates a new random []byte or panics.
func NewWithLength(length int) []byte {
	return must(NewRandomWithLength(length))
}

// NewBase64WithLength creates a new random []byte and returns it as a base64 string or panics.
func NewBase64WithLength(length int) string {
	return base64.URLEncoding.EncodeToString(must(NewRandomWithLength(length)))
}

// NewBase32WithLength creates a new random []byte and returns it as a base32 string or panics.
func NewBase32WithLength(length int) string {
	return base32.StdEncoding.EncodeToString(must(NewRandomWithLength(length)))
}

// NewHexWithLength creates a new random []byte and returns it as a hex string or panics.
func NewHexWithLength(length int) string {
	return hex.EncodeToString(must(NewRandomWithLength(length)))
}

// NewRandomWithLength returns a random []byte.
// The strength of the []bytes is based on the strength of the crypto/rand package.
func NewRandomWithLength(length int) ([]byte, error) {
	return NewRandomFromReader(defaultRander, length)
}

// NewRandomFromReader returns a []byte based on bytes read from a given io.Reader.
func NewRandomFromReader(reader io.Reader, length int) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := io.ReadFull(reader, bytes[:])
	return bytes, err
}

func must(bytes []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return bytes
}
