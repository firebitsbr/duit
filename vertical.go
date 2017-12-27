package duit

import (
	"image"

	"9fans.net/go/draw"
)

type Vertical struct {
	Kids  []*Kid
	Split func(height int) (heights []int)

	size    image.Point
	heights []int
}

var _ UI = &Vertical{}

func (ui *Vertical) Layout(env *Env, size image.Point) image.Point {
	heights := ui.Split(size.Y)
	if len(heights) != len(ui.Kids) {
		panic("bad number of heights from split")
	}
	ui.heights = heights
	cur := image.ZP
	for i, k := range ui.Kids {
		childSize := k.UI.Layout(env, image.Pt(size.X, heights[i]))
		k.r = rect(childSize).Add(cur)
		cur.Y += heights[i]
	}
	ui.size = image.Pt(size.X, cur.Y)
	return ui.size
}

func (ui *Vertical) Draw(env *Env, img *draw.Image, orig image.Point, m draw.Mouse) {
	kidsDraw(env, ui.Kids, ui.size, img, orig, m)
}

func (ui *Vertical) Mouse(env *Env, m draw.Mouse) (result Result) {
	return kidsMouse(env, ui.Kids, m)
}

func (ui *Vertical) Key(env *Env, orig image.Point, m draw.Mouse, k rune) (result Result) {
	return kidsKey(env, ui, ui.Kids, orig, m, k)
}

func (ui *Vertical) FirstFocus(env *Env) *image.Point {
	return kidsFirstFocus(env, ui.Kids)
}

func (ui *Vertical) Focus(env *Env, o UI) *image.Point {
	return kidsFocus(env, ui.Kids, o)
}

func (ui *Vertical) Print(indent int, r image.Rectangle) {
	uiPrint("Vertical", indent, r)
	kidsPrint(ui.Kids, indent+1)
}
