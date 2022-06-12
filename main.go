package main

import (
	"flag"
	"os"
)

func main() {
	flagsHeader := flag.NewFlagSet("header", flag.ExitOnError)
	flagsHeaderTitle := flagsHeader.Bool("title", false, "A larger title header")
	flagsHeaderStyle := flagsHeader.String("style", "", "The box style (heavy, light, double, flat)")

	var element string
	if len(os.Args) > 1 {
		element = os.Args[1]
	}

	switch element {
	case "header":
		err := flagsHeader.Parse(os.Args[2:])
		if err != nil {
			return
		}
		h := NewHeader(flagsHeader.Arg(0))
		if *flagsHeaderTitle {
			h = HeaderTitle(h)
		}
		switch *flagsHeaderStyle {
		case "", "heavy":
			h = HeaderStyleHeavy(h)
		case "light":
			h = HeaderStyleLight(h)
		case "double":
			h = HeaderStyleDouble(h)
		case "flat":
			h = HeaderStyleFlat(h)
		}
		h.Render()

	case "alert-done":
		NewDoneAlert(os.Args[2]).Render()

	case "alert-info":
		NewInfoAlert(os.Args[2]).Render()

	case "alert-warn":
		NewWarnAlert(os.Args[2]).Render()

	case "alert-error":
		NewErrAlert(os.Args[2]).Render()

	default:
		println("usage: terminal-ui [header|alert-done|alert-info|alert-warn|alert-error] [flags] text")
		os.Exit(1)
	}
}
