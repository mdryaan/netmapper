package output

import (
	"fmt"
	"os"
)

func Println(s string) {
	fmt.Fprintln(os.Stdout, s)
}

func Printf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, format, args...)
}

func Errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, Red(fmt.Sprintf(format, args...))+"\n")
}

func Successf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, Success(fmt.Sprintf(format, args...))+"\n")
}

func Warnf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, Warn(fmt.Sprintf(format, args...))+"\n")
}

func Infof(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, Info(fmt.Sprintf(format, args...))+"\n")
}

func Header(title string) {
	fmt.Fprintln(os.Stdout, Bold(title))
}

func Fatal(msg string) {
	Errorf("%s", msg)
	os.Exit(1)
}
