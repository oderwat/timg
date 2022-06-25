package must

import (
	"fmt"
	"github.com/ysmood/gotrace"
	"log"
	"os"
	"runtime/debug"
	"syscall"
)

// BeQuiet sets stdout, stderr and the logger to null and returns a function to revert to the former values
func BeQuiet() func() {
	null, _ := os.Open(os.DevNull)
	sout := os.Stdout
	serr := os.Stderr
	os.Stdout = null
	os.Stderr = null
	log.SetOutput(null)
	return func() {
		defer func(null *os.File) {
			_ = null.Close()
		}(null)
		os.Stdout = sout
		os.Stderr = serr
		log.SetOutput(os.Stderr)
	}
}

// Protect is used to catch panics
func Protect(run func(), reportPanic func(any)) {
	defer func() {
		if x := recover(); x != nil {
			if reportPanic == nil {
				log.Printf("*** run time panic: %v", x)
				debug.PrintStack()
			} else {
				reportPanic(x)
			}
		}
	}()
	run()
}

// CheckLock let's us ctrl-c (or KILL) and we get a stacktrace of where it was at that moment
// good for checking deadlocks or running with long timeouts
func CheckLock(report func(traces gotrace.Traces)) func() {
	ctx := gotrace.Signal(syscall.SIGINT, syscall.SIGTERM)
	ignored := gotrace.IgnoreCurrent()
	return func() {
		traces := gotrace.Wait(ctx, ignored)
		if traces.Any() {
			if report != nil {
				report(traces)
			} else {
				fmt.Println(traces)
			}
		}
	}
}
