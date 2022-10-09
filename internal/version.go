// Copyright (c) 2022 iwaltgen
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package internal

import (
	"strconv"
	"time"
)

var (
	version    = "dev"
	commitHash = "dev"
	buildDate  = "1640995200" // Sat Jan 01 2022 00:00:00 GMT+0000
)

// Version applied app version
func Version() string {
	return version
}

// CommitHash applied git hash
func CommitHash() string {
	return commitHash
}

// BuildDate applied build time
func BuildDate() time.Time {
	ts, _ := strconv.ParseInt(buildDate, 10, 64)
	return time.Unix(ts, 0).UTC()
}
