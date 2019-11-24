package main

import (
	"flag"
	"os"

	"github.com/Uclok31v/image_conversion/convert"
)

var (
	before = flag.String("b", "jpg", "変換前の画像形式")
	after  = flag.String("a", "png", "返還後の画像形式")
)

func main() {
	flag.Parse()
	cli := convert.Cli{flag.Arg(0), before, after}
	os.Exit(cli.Execute())
}
