package main

import (
	"neiwang/Common"
	"neiwang/Scan"
)

func main() {
	var Cmd Common.Cmd
	Common.Flag(&Cmd)
	Scan.CheckURL(Cmd)

}
