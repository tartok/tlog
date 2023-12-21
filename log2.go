package tlog

import (
	"github.com/TwiN/go-color"
	"io"
	"log"
	"os"
)

// Loggers
type Loggers struct {
	Debug *log.Logger
	Log   *log.Logger
	Con   *log.Logger
	Err   *log.Logger
}

var Debug = log.New(io.Discard, "", log.Lmsgprefix+log.LstdFlags)
var Log = log.New(io.Discard, "", log.Lmsgprefix+log.LstdFlags)
var Con = log.New(io.Discard, "", log.Lmsgprefix+log.LstdFlags)
var Err = log.New(io.Discard, "", log.Lmsgprefix+log.LstdFlags)

var DefLoggers = Loggers{
	Debug: Debug,
	Log:   Log,
	Con:   Con,
	Err:   Err,
}

func InitDebug(prefix string, outs ...func(p []byte) (n int, err error)) {
	Debug.SetPrefix(prefix)
	Debug.SetOutput(GetOut(color.Blue, outs...))
}
func InitLog(prefix string, outs ...func(p []byte) (n int, err error)) {
	Log.SetPrefix(prefix)
	Log.SetOutput(GetOut(color.Gray, outs...))
}
func InitCon(prefix string, outs ...func(p []byte) (n int, err error)) {
	Con.SetPrefix(prefix)
	Con.SetOutput(GetOut(color.Green, outs...))
}
func InitErr(prefix string, outs ...func(p []byte) (n int, err error)) {
	Err.SetPrefix(prefix)
	Err.SetOutput(GetOut(color.Red, outs...))
}

func init() {
	InitErr("", os.Stdout.Write)
	InitCon("", os.Stdout.Write)
}
