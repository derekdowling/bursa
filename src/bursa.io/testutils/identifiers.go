package testutils

import (
	"fmt"
	"strings"
	"time"
)

var test_id string

func init() {
	TestId()
}

func TestId() string {
	if test_id == "" {
		test_id = fmt.Sprintf("%d", time.Now().UnixNano())
	}
	return test_id
}

func SuffixedId(suffix string) string {
	return strings.Join([]string{"test", TestId(), suffix}, ":")
}

// Returns an email safe suffixed test id
func EmailSuffixedId(suffix string) string {
	return strings.Join([]string{"test", TestId(), suffix}, "_")
}
