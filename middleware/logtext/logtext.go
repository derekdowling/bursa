package logtext

// A Logrus hook that adds filename/line number/stack trace to our log outputs

import (
	"github.com/Sirupsen/logrus"
	"runtime"
	"strings"
)

var blacklist = []string{
	"github.com/Sirupsen/logrus",
	"github.com/derekdowling/bursa/middleware/logtext",
	"github.com/smartystreets/gconvey",
	"github.com/smartystreets/goconvey",
	"github.com/jtolds/gls",
}

// Log depth is how many levels to ascend to find where the actual log call occurred
// while debugOnly sets whether or not stack traces should be printed outside of
// debug prints
type Logtext struct{
	Formatter logrus.Formatter
	LogDepth int
	DebugOnly bool
}

func NewLogtext(formatter logrus.Formatter, debugOnly bool) *Logtext {
	return &Logtext{
		LogDepth: 4,
		Formatter: formatter,
		DebugOnly: debugOnly,
	}
}

// Creates a hook to be added to an instance of logger. This is called with
func (hook *Logtext) Format(entry *logrus.Entry) ([]byte, error) {
	if !hook.DebugOnly || entry.Level == logrus.DebugLevel {
		stack := getTrace()
		entry.Data["stack"] = stack[0:4]
		if len(stack) > 0 {
			entry.Data["line"] = stack[1]
			entry.Data["file"] = stack[1]
		}
	} else if _, file, line, ok := runtime.Caller(hook.LogDepth); ok {
		entry.Data["line"] = line
		entry.Data["file"] = file
	}

	return hook.Formatter.Format(entry);
}

// Returns the trace as an array. Yes this is crazy expensive - probably.
func getTrace() []string {
	stack := make([]byte, 1024*32)
	size := runtime.Stack(stack, true)

	return filterTrace(
		strings.Split(string(stack[:size]),"\n"),
	)
}

func filterTrace(frames []string) []string {

	var result = []string{}
	for _, frame := range frames {
			if isBlackListed(frame) == false {
				result = append(result, frame)
			}
	}
	return result
}

func isBlackListed(frame string) bool {
	for _, exclude := range blacklist {
		if strings.Contains(frame, exclude) == true {
			return true
		}
	}
	return false
}
