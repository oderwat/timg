# timg
timg is a tool that shows images in the terminal (iTerm2 only)

**Repeat: iTerm2 only!**

**NOTICE:** Currently, there seems to be a [bug with iTerm2](https://gitlab.com/gnachman/iterm2/-/issues/10420) that does not free memory. This tool was basically made for making the problem reproducible by the developers of iTerm2

Install with: `go install github.com/oderwat/timg@latest`

This lists all you Photos with a small thumbnail in your terminal: `timg -r -h 80 ~/Pictures/Photos\ Library.photoslibrary/originals/` 

```
NAME:
   timg - shows a single image or all images of a directory inside the terminal (e.g. ITerm2)

USAGE:
   timg [global options] [commands] [comand options]

   The default command is show and can be omitted. See 'timg h show' for help.

   Examples: 'timg ~/Pictures/', 'timg -r -h 80 ~/Pictures/Photos\ Library.photoslibrary/originals/'

VERSION:
   v0.1.0

COMMANDS:
   panic    just for testing a panic
   show     shows the image or all images of a directory inside the terminal (e.g. ITerm2). The path can be a glob path with ?/* and ** (e.g. ~/images/*.jpg, ~/**/*.jpg)
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --height value, -h value           sets the height of the output (default: 0)
   --help, -?                         show help (default: false)
   --log-level value                  logging level (default: 3)
   --perf                             Giving some performance statistics at the end (default: false)
   --print-version, -V                print only the version (default: false)
   --recursive, -r                    Giving some performance statistics at the end (default: false)
   --show-unknown-extensions, --uext  show files that are not of a known extension and get skipped (default: false)
   --use-extensions, -x               use extensions to predict what files can be shown (default: true)
   --width value, -w value            sets the with of the output (default: 0)
```
