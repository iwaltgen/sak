// Copyright (c) 2022 iwaltgen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package uuid

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/hashicorp/go-uuid"
	"github.com/jxskiss/base62"
	"github.com/oklog/ulid/v2"
	"github.com/rs/xid"
	"github.com/segmentio/ksuid"
)

const uuidLen = 16

// New creates a new random UUID or panics.
func New() string {
	return must(uuid.GenerateUUID())
}

// NewBase64 creates a new random []byte and returns it as a base64 string or panics.
func NewBase64() string {
	return base64.RawURLEncoding.EncodeToString(must(uuid.GenerateRandomBytes(uuidLen)))
}

// NewBase62 creates a new random []byte and returns it as a base62 string or panics.
func NewBase62() string {
	return base62.EncodeToString(must(uuid.GenerateRandomBytes(uuidLen)))
}

// NewHex creates a new random []byte and returns it as a hex string or panics.
func NewHex() string {
	return hex.EncodeToString(must(uuid.GenerateRandomBytes(uuidLen)))
}

// NewKSUID creates a new random []byte and returns it as a KSUID string or panics.
// https://github.com/segmentio/ksuid
func NewKSUID() string {
	return ksuid.New().String()
}

// NewULID creates a new random []byte and returns it as a ULID string or panics.
// https://github.com/oklog/ulid
func NewULID() string {
	return ulid.Make().String()
}

// NewXID creates a new random []byte and returns it as a XID string or panics.
// https://github.com/rs/xid
func NewXID() string {
	return xid.New().String()
}

func must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
