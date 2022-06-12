package main

import "github.com/fatih/color"

type Alert struct {
	LabelText  string
	LabelColor *color.Color
	Text       string
	TextColor  *color.Color
}

func NewDoneAlert(text string) Alert {
	return Alert{
		LabelText:  "DONE",
		LabelColor: color.New(color.BgGreen, color.FgHiWhite),
		Text:       text,
		TextColor:  color.New(color.FgGreen),
	}
}

func NewErrAlert(text string) Alert {
	return Alert{
		LabelText:  "ERRO",
		LabelColor: color.New(color.BgHiRed, color.FgHiWhite),
		Text:       text,
		TextColor:  color.New(color.FgRed),
	}
}

func NewWarnAlert(text string) Alert {
	return Alert{
		LabelText:  "WARN",
		LabelColor: color.New(color.BgHiYellow, color.FgBlack, color.Bold),
		Text:       text,
		TextColor:  color.New(color.FgYellow),
	}
}

func NewInfoAlert(text string) Alert {
	return Alert{
		LabelText:  "INFO",
		LabelColor: color.New(color.BgCyan, color.FgHiWhite),
		Text:       text,
		TextColor:  color.New(color.FgCyan),
	}
}

func (a Alert) Render() {
	resetColor := color.New(color.Reset)

	resetColor.Println()
	a.LabelColor.Print(" " + a.LabelText + " ")
	a.TextColor.Print("  " + a.Text + "  ")
	resetColor.Println()
}
