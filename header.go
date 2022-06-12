package main

import (
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
)

var HeaderHeavyBorders = HeaderBorders{
	H:  "━",
	V:  "┃",
	TL: "┏",
	TR: "┓",
	BL: "┗",
	BR: "┛",
}
var HeaderLightBorders = HeaderBorders{
	H:  "─",
	V:  "│",
	TL: "┌",
	TR: "┐",
	BL: "└",
	BR: "┘",
}
var HeaderDoubleBorders = HeaderBorders{
	H:  "═",
	V:  "║",
	TL: "╔",
	TR: "╗",
	BL: "╚",
	BR: "╝",
}

type Header struct {
	Text        string
	TextLen     int
	PaddingH    int
	PaddingV    int
	TextColor   color.Color
	BorderColor color.Color
	ResetColor  color.Color
	Borders     HeaderBorders
}

func NewHeader(text string) Header {
	return Header{
		Text:     text,
		Borders:  HeaderHeavyBorders,
		PaddingH: 4,
		PaddingV: 0,
	}
}

func HeaderTitle(header Header) Header {
	header.PaddingV = 1
	return header
}

func HeaderStyleHeavy(header Header) Header {
	header.Borders = HeaderHeavyBorders
	return header
}

func HeaderStyleLight(header Header) Header {
	header.Borders = HeaderLightBorders
	return header
}

func HeaderStyleDouble(header Header) Header {
	header.Borders = HeaderDoubleBorders
	return header
}

func HeaderStyleFlat(header Header) Header {
	header.Borders = HeaderLightBorders
	header.BorderColor.Add(color.Concealed, color.ReverseVideo)
	header.TextColor.Add(color.ReverseVideo)
	return header
}

type HeaderBorders struct {
	H  string
	V  string
	TL string
	TR string
	BL string
	BR string
}

func (h Header) Render() {
	h.TextLen = utf8.RuneCountInString(h.Text)

	h.ResetColor.Println()
	if h.PaddingV >= 0 {
		h.renderTop()
		h.renderVPadding()
	}
	h.renderText()
	if h.PaddingV >= 0 {
		h.renderVPadding()
		h.renderBottom()
	}
}

func (h Header) renderTop() {
	h.BorderColor.Print(h.Borders.TL)
	h.BorderColor.Print(strings.Repeat(h.Borders.H, h.PaddingH*2+h.TextLen-2))
	h.BorderColor.Print(h.Borders.TR)
	h.ResetColor.Println()
}

func (h Header) renderBottom() {
	h.BorderColor.Print(h.Borders.BL)
	h.BorderColor.Print(strings.Repeat(h.Borders.H, h.PaddingH*2+h.TextLen-2))
	h.BorderColor.Print(h.Borders.BR)
	h.ResetColor.Println()
}

func (h Header) renderVPadding() {
	for i := 0; i < h.PaddingV; i++ {
		h.BorderColor.Print(h.Borders.V)
		h.TextColor.Print(strings.Repeat(" ", 2*h.PaddingH+h.TextLen-2))
		h.BorderColor.Print(h.Borders.V)
		h.ResetColor.Println()
	}
}

func (h Header) renderText() {
	h.BorderColor.Print(h.Borders.V)
	h.TextColor.Print(strings.Repeat(" ", h.PaddingH-1))
	h.TextColor.Print(h.Text)
	h.TextColor.Print(strings.Repeat(" ", h.PaddingH-1))
	h.BorderColor.Print(h.Borders.V)
	h.ResetColor.Println()
}
