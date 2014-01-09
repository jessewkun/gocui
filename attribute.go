package gocui

import "github.com/nsf/termbox-go"

type Attribute termbox.Attribute

const (
	ColorDefault Attribute = Attribute(termbox.ColorDefault)
	ColorBlack             = Attribute(termbox.ColorBlack)
	ColorRed               = Attribute(termbox.ColorRed)
	ColorGreen             = Attribute(termbox.ColorGreen)
	ColorYellow            = Attribute(termbox.ColorYellow)
	ColorBlue              = Attribute(termbox.ColorBlue)
	ColorMagenta           = Attribute(termbox.ColorMagenta)
	ColorCyan              = Attribute(termbox.ColorCyan)
	ColorWhite             = Attribute(termbox.ColorWhite)
)

const (
	AttrBold      Attribute = Attribute(termbox.AttrBold)
	AttrUnderline           = Attribute(termbox.AttrUnderline)
	AttrReverse             = Attribute(termbox.AttrReverse)
)