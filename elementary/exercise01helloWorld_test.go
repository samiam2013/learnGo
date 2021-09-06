package elementary

import (
	"testing"

	"github.com/samiam2013/learnGo/testUtils"
)

func TestHelloWorld(t *testing.T) {
	got := testUtils.CaptureOutput(HelloWorld, "")
	if got != "Hello World!\n" {
		t.Errorf("HelloWorld() = %s; want \"Hello World!\\n\"", got)
	}
}
