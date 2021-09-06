package elementary

import (
	"testing"

	"github.com/samiam2013/learnGo/testutils"
)

func TestHelloWorld(t *testing.T) {
	got := testutils.CaptureOutput(HelloWorld, "")
	if got != "Hello World!\n" {
		t.Errorf("HelloWorld() = %s; want \"Hello World!\\n\"", got)
	}
}
