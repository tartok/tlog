package tlog

import (
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var std = Con
var Config = struct {
	Debug bool
	Log   bool
}{Debug: false, Log: true}

func errorCheck(vs ...interface{}) error {
	for _, v := range vs {
		if err, ok := v.(error); ok {
			return err
		}
	}
	return nil
}
func formatf(from int, to int, f string, vs ...interface{}) (res string, err error) {
	err = errorCheck(vs...)
	d := ""
	for i := from; i <= to; i++ {
		pc, fn, li, _ := runtime.Caller(i)
		d = fmt.Sprintf("%s\n%s\n  %s:%d", d, runtime.FuncForPC(pc).Name(), fn, li)
	}
	res = fmt.Sprintf(f, vs...) + d

	//s := log.Printf("%s Err: %s\n %s:%d", prefix, err, fn, li)
	return res, err
}

func Warningf(depth int, format string, log ...interface{}) bool {
	s, err := formatf(2, depth+1, format, log...)
	if err != nil {
		Print(s)
		return true
	}
	return false
}
func Warning(depth int, log ...interface{}) bool {
	if errorCheck(log...) != nil {
		s, _ := formatf(2, depth+1, "%s", fmt.Sprintln(log...))
		Print(s)
		return true
	}
	return false
}
func Swarning(depth int, log ...interface{}) error {
	if errorCheck(log...) != nil {
		s, _ := formatf(2, depth+1, "%s", fmt.Sprintln(log...))
		return fmt.Errorf(s)
	}
	return nil
}

func WarningWin1251(depth int, log ...interface{}) bool {
	if errorCheck(log...) != nil {
		s, _ := formatf(2, depth+1, "%s", fmt.Sprintln(log...))
		dec := charmap.Windows1251.NewDecoder()
		s, err := dec.String(s)
		if err == nil {
			Print(s)
		}
		return true
	}
	return false
}
func PrintWin1251(log ...interface{}) {
	s := fmt.Sprintln(log...)
	dec := charmap.Windows1251.NewDecoder()
	s, err := dec.String(s)
	if err == nil {
		Print(s)
	}
}
func Error(depth int, log ...interface{}) bool {
	if errorCheck(log...) != nil {
		s, _ := formatf(2, depth+1, "%s", fmt.Sprintln(log...))
		Print(s)
		os.Exit(1)
	}
	return false
}
func DebugF(depth int, format string, log ...interface{}) bool {
	if !Config.Debug {
		return false
	}
	s, err := formatf(2, depth+1, format, log...)
	Print(s)
	return err != nil
}
func LogF(depth int, format string, log ...interface{}) bool {
	if !Config.Log {
		return false
	}
	s, err := formatf(2, depth+1, format, log...)
	Print(s)
	return err != nil
}
func DebugFrom(from, depth int, format string, log ...interface{}) bool {
	if !Config.Debug {
		return false
	}
	s, err := formatf(from+2, depth+1, format, log...)
	Print(s)
	return err != nil
}

// Flags returns the output flags for the standard logger.
func Flags() int {
	return std.Flags()
}

// SetFlags sets the output flags for the standard logger.
func SetFlags(flag int) {
	std.SetFlags(flag)
}

// Prefix returns the output prefix for the standard logger.
func Prefix() string {
	return std.Prefix()
}

// SetPrefix sets the output prefix for the standard logger.
func SetPrefix(prefix string) {
	std.SetPrefix(prefix)
}

// These functions write to the standard logger.

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	std.Output(2, fmt.Sprint(v...))
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf(format, v...))
}

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	std.Output(2, fmt.Sprintln(v...))
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	std.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...interface{}) {
	std.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

// Panic is equivalent to Print() followed by a call to panic().
func Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	std.Output(2, s)
	panic(s)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	std.Output(2, s)
	panic(s)
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	std.Output(2, s)
	panic(s)
}

// Output writes the output for a logging event. The string s contains
// the text to print after the prefix specified by the flags of the
// Logger. A newline is appended if the last character of s is not
// already a newline. Calldepth is the count of the number of
// frames to skip when computing the file name and line number
// if Llongfile or Lshortfile is set; a value of 1 will print the details
// for the caller of Output.
func Output(calldepth int, s string) error {
	return std.Output(calldepth+1, s) // +1 for this frame.
}

func DateTimeToStr(t time.Time) string {
	return t.Format("2006/01/02 15:04:05")
}
func CreateFile() (*os.File, error) {
	ex, err := os.Executable()
	if err != nil {
		return nil, err
	}
	patch := filepath.Dir(ex)
	patch += "/log"
	err = os.MkdirAll(patch, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return os.Create(fmt.Sprintf("%s/%s.log", patch, time.Now().Format("20060102150405")))

}

//
//func ToFile() {
//	ex, err := os.Executable()
//	if Warning(1, err) {
//		return
//	}
//	patch := filepath.Dir(ex)
//	patch += "/log"
//	err = os.MkdirAll(patch, os.ModePerm)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	w.file, err = os.Create(fmt.Sprintf("%s/%s.log", patch, time.Now().Format("20060102150405")))
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//}
//func ToChan(c chan<- []byte) {
//	w.chanW = c
//}
//func init() {
//	w.con = os.Stdout
//	std = log.New(w, "", log.LstdFlags)
//}
