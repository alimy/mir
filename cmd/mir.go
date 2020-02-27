package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/alimy/mir/v2/cmd/create"
)

const (
	appVer = "0.1.0"
)

var (
	dstPath  string
	style    string
	showHelp bool
)

func init() {
	flag.StringVar(&dstPath, "dst", ".", "genereted destination target directory, default '.'")
	flag.StringVar(&style, "type", "gin", "generated engine type style(eg: gin,chi,mux,httprout), default is 'gin'")
	flag.BoolVar(&showHelp, "help", false, "print help usage info")
}

func main() {
	flag.Parse()

	if showHelp {
		usage()
		return
	}

	cmds := flag.Args()
	if len(cmds) != 1 {
		usage()
		return
	}

	switch cmds[0] {
	case "new":
		if err := create.RunCmd(dstPath, style); err != nil {
			log.Fatal(err)
		}

	case "version":
		fmt.Println(appVer)

	default:
		usage()
	}

}

func usage() {
	fmt.Printf(`NAME:
   Mir - A usefull toolkit for help build web application

USAGE:
   mir command [options]

VERSION:
   %s

COMMANDS:
   new      generate a template project
   version  print the version

OPTIONS:
   -dst destination target directory, default '.'
   -style engine type style(eg: gin,chi,mux,httprout), default is 'gin'
   -help     show help%s`, appVer, "\n")
}
