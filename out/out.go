package out

import (
	"fmt"
	"github.com/oderwat/timg/must"
	"os"
)

var clearLine bool

func Outf(format string, v ...any) {
	if clearLine {
		fmt.Print("\033[2K\r")
		clearLine = false
	}
	fmt.Printf(format, v...)
}

func OutfKeep(format string, v ...any) {
	if clearLine {
		fmt.Print("\n")
		clearLine = false
	}
	fmt.Printf(format, v...)
}

func OutfR(format string, v ...any) {
	format += "\r"
	clearLine = true
	fmt.Printf(format, v...)
}

func Outln(v ...any) {
	if clearLine {
		fmt.Print("\033[2K\r")
		clearLine = false
	}
	fmt.Println(v...)
}

func Errf(format string, v ...any) {
	// even if we output it on stderr it will mix badly into our other output
	if clearLine {
		fmt.Print("\n")
		clearLine = false
	}
	must.OkSkipOne(fmt.Fprintf(os.Stderr, fmt.Errorf(format, v...).Error()))
}
