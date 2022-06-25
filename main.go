package main

import (
	"errors"
	"fmt"
	"github.com/oderwat/timg/humz"
	"github.com/oderwat/timg/must"
	"github.com/oderwat/timg/out"
	"github.com/urfave/cli/v2"
	"github.com/ysmood/gotrace"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

const (
	LogDebug  = 1
	LogInfo   = 2
	LogNormal = 3
	LogWarn   = 4
	LogErr    = 5
)

var LogLevel = LogNormal

var flagLogLevel = &cli.IntFlag{
	Name:  "log-level",
	Value: LogLevel,
	Usage: "logging level",
}

var flagPerf = &cli.BoolFlag{
	Name:  "perf",
	Usage: "Giving some performance statistics at the end",
	Value: false,
}

var flagRecursive = &cli.BoolFlag{
	Name:    "recursive",
	Aliases: []string{"r"},
	Usage:   "Giving some performance statistics at the end",
	Value:   false,
}
var flagShowUnknownTypes = &cli.BoolFlag{
	Name:    "show-unknown-extensions",
	Aliases: []string{"uext"},
	Value:   false,
	Usage:   "show files that are not of a known extension and get skipped",
}

var flagUseExtensions = &cli.BoolFlag{
	Name:    "use-extensions",
	Aliases: []string{"x"},
	Value:   true,
	Usage:   "use extensions to predict what files can be shown",
}

var flagWidth = &cli.IntFlag{
	Name:    "width",
	Aliases: []string{"w"},
	Value:   0,
	Usage:   "sets the with of the output",
}

var flagHeight = &cli.IntFlag{
	Name:    "height",
	Aliases: []string{"h"},
	Value:   0,
	Usage:   "sets the height of the output",
}

var flagVersion = &cli.BoolFlag{
	Name:    "print-version",
	Aliases: []string{"V"},
	Usage:   "print only the version",
}

var cmdPanic = &cli.Command{
	Name:  "panic",
	Usage: "just for testing a panic",
	Action: func(c *cli.Context) error {
		must.Ok(errors.New("don't panic"))
		return nil
	},
}

func main() {
	cli.VersionFlag = flagVersion
	started := time.Now()
	app := &cli.App{
		Name:  "timg",
		Usage: "shows a single image or all images of a directory inside the terminal (e.g. ITerm2)",
		UsageText: "timg [global options] [commands] [comand options]\n\n" +
			"The default command is show and can be omitted. See 'timg h show' for help.\n\n" +
			"Examples: 'timg ~/Pictures/', 'timg -r -h 80 ~/Pictures/Photos\\ Library.photoslibrary/originals/'",
		Version:  "v0.1.0",
		Compiled: time.Now(),
		Flags: []cli.Flag{
			flagLogLevel,
			flagPerf,
			flagWidth,
			flagHeight,
			flagUseExtensions,
			flagShowUnknownTypes,
			flagRecursive,
		},
		Before: func(c *cli.Context) error {
			LogLevel = c.Int(flagLogLevel.Name)
			if LogLevel <= LogDebug {
				out.Outf("Set Log Level %d\n", LogLevel)
			}
			return nil
		},
		ArgsUsage: "[imagepath, ...]",
		Action:    runTImg,

		After: func(c *cli.Context) error {
			if len(os.Args) > 1 && os.Args[len(os.Args)-1] == "--generate-bash-completion" {
				return nil
			}
			if c.Bool(flagPerf.Name) {
				var m runtime.MemStats
				// For info on each, see: https://golang.org/pkg/runtime/#MemStats
				runtime.ReadMemStats(&m)

				out.Outf("Performance:\n")
				out.Outf("  Time taken: %s\n", humz.Duration(time.Since(started)))
				out.Outf("  Alloc     : %s\n", humz.Bytes(int64(m.Alloc)))
				out.Outf("  TotalAlloc: %s\n", humz.Bytes(int64(m.TotalAlloc)))
				out.Outf("  Sys       : %s\n", humz.Bytes(int64(m.Sys)))
				out.Outf("  NumGC     : %s\n", humz.Count(int64(m.NumGC)))
			}
			return nil
		},
	}
	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"?"},
		Usage:   "show help",
	}

	app.Commands = append(app.Commands, cmdPanic)
	app.Commands = append(app.Commands, cmdTImg())

	app.EnableBashCompletion = true

	// add all our commands
	exitCode := 0
	must.Protect(func() {
		defer must.CheckLock(func(x gotrace.Traces) {
			// here we end up with CTRL+C (and other signals)
			out.OutfKeep("*** Aborted\n")
			if LogLevel <= LogDebug {
				fmt.Println(x)
			}
			exitCode = 1
		})()
		go func() {
			err := app.Run(os.Args)
			// this is the regular ending
			if err != nil {
				out.OutfKeep("%s\n", err.Error())
				if LogLevel <= LogDebug {
					debug.PrintStack()
				}
				exitCode = 2
			}
		}()
	}, func(x any) {
		// here we end up if a panic happened
		out.OutfKeep("*** Panic: %v!\n", x)
		if LogLevel <= LogDebug {
			debug.PrintStack()
		}
		exitCode = 3
	})
	os.Exit(exitCode)
}
