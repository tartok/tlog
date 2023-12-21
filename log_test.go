package tlog

import (
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	InitLog("test ", os.Stdout.Write)
	Log.Println("1")
}
