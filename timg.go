package main

import (
	"fmt"
	"github.com/oderwat/timg/must"
	"github.com/urfave/cli/v2"
	"io"
	"modernc.org/strutil"
	"os"
	"path/filepath"
)

func cmdTImg() *cli.Command {
	return &cli.Command{
		Name:      "show",
		Usage:     "shows the image or all images of a directory inside the terminal (e.g. ITerm2). The path can be a glob path with ?/* and ** (e.g. ~/images/*.jpg, ~/**/*.jpg)",
		ArgsUsage: "[imagepath, ...]",
		Action:    runTImg,
		Flags: []cli.Flag{
			flagWidth,
			flagHeight,
			flagUseExtensions,
			flagShowUnknownTypes,
		},
	}

}

func runTImg(c *cli.Context) error {
	skipDirs := map[string]struct{}{".Spotlight-V100": {}, ".Trashes": {}, ".fseventsd": {},
		".TemporaryItems": {}, ".DocumentRevisions-V100": {}}

	if c.NArg() < 1 {
		return cli.Exit("need a image file or path", 1)
	}
	iw := c.Int(flagWidth.Name) / 8
	ih := c.Int(flagHeight.Name) / 8

	recursive := c.Bool(flagRecursive.Name)

	scanPaths := c.Args().Slice()
	for _, scanPath := range scanPaths {
		topLevel := true
		e := filepath.WalkDir(scanPath, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				//fmt.Printf("can't access path %q: %v\n", path, err)
				return nil
			}
			filename := d.Name()

			// Check against out hard skip dirs
			if d.IsDir() {
				if !recursive && !topLevel {
					return filepath.SkipDir
				}
				topLevel = false
				_, ok := skipDirs[filename]
				if ok {
					return filepath.SkipDir
				}
				return nil
			}

			if c.Bool(flagUseExtensions.Name) {
				switch extensionType(filename) {
				case extUnknown:
					return nil
				case extSkipped:
					if c.Bool(flagShowUnknownTypes.Name) {
						fmt.Printf("Skipping: %q\n", path[len(scanPath):])
					}
					return nil
				}
			}

			if filename == "Thumbs.db" || filename == ".picasa.ini" || filename == ".DS_Store" {
				return nil
			}

			if LogLevel <= LogNormal {
				fmt.Println(path)
			}
			Show(path, iw, ih)
			return nil
		})
		if e != nil {
			panic(e)
		}
	}
	return nil
}

func FShow(w io.Writer, path string, width int, height int) {
	term := os.Getenv("TERM")
	if term == "xscreen" {
		must.OkSkipOne(fmt.Fprint(w, "\033Ptmux;\033\033]"))
	} else {
		must.OkSkipOne(fmt.Fprint(w, "\033]"))
	}
	fi := must.OkOne(os.Stat(path))
	data := must.OkOne(os.ReadFile(path))

	sizeFmt := ""
	if width != 0 && height == 0 {
		sizeFmt = fmt.Sprintf(";width=%d;preserveAspectRatio=true", width)
	} else if width == 0 && height != 0 {
		sizeFmt = fmt.Sprintf(";height=%d;preserveAspectRatio=true", height)
	} else if width != 0 && height != 0 {
		sizeFmt = fmt.Sprintf(";width=%d;height=%d;preserveAspectRatio=true", width, height)
	}

	must.OkSkipOne(fmt.Fprintf(w, `1337;File=name=%s;size=%d;inline=1%s:%s`,
		strutil.Base64Encode([]byte(filepath.Base(path))), fi.Size(), sizeFmt, strutil.Base64Encode(data)))
	if term == "xscreen" {
		must.OkSkipOne(fmt.Fprint(w, "\a\033\\"))
	} else {
		must.OkSkipOne(fmt.Fprint(w, "\a"))
	}
	if term == "xscreen" {
		must.OkSkipOne(fmt.Fprint(w, "\u001B[4C\u001B[Bx\n\n"))
	} else {
		must.OkSkipOne(fmt.Fprint(w, "\033[A\n\n"))
	}
}

func Show(path string, width int, height int) {
	FShow(os.Stdout, path, width, height)
}
