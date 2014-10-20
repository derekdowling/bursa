package testutils

import(
	"fmt"
	"time"
	"strings"
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
	return strings.Join([]string{"test",TestId(), suffix},":")
}
