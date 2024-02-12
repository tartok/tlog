package tlog_test

import (
	"github.com/tartok/tlog"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	tlog.Log.SetOutput(tlog.GetOut("", os.Stdout))
	tlog.Log.Println("1")
}
