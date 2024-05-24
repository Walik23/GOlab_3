package painter

import (
	"image"
	"image/color"

	"github.com/roman-mazur/architecture-lab-3/ui"

	"golang.org/x/exp/shiny/screen"
)

type Operation interface {
	Update(state *TextureState)
}

type TextureOperation interface {
	Do(t screen.Texture)
	Update(state *TextureState)
}

type OperationList []Operation

var UpdateOp = Update{}

type Update struct{}

func (op Update) Update(_ *TextureState) {}

type Fill struct {
	Color color.Color
}

func (op Fill) Do(t screen.Texture) {
	t.Fill(t.Bounds(), op.Color, screen.Src)
}

func (op Fill) Update(state *TextureState) {
	state.backgroundColor = &op
}

type Reset struct{}

var ResetOp = Reset{}

func (op Reset) Update(state *TextureState) {
	state.backgroundColor = &Fill{Color: color.Black}
	state.backgroundRect = nil
	state.figureCenters = nil
}

type BgRect struct {
	X1 float32
	Y1 float32
	X2 float32
	Y2 float32
}

func (op BgRect) Do(t screen.Texture) {
	t.Fill(
		image.Rect(
			int(op.X1*float32(t.Size().X)),
			int(op.Y1*float32(t.Size().Y)),
			int(op.X2*float32(t.Size().X)),
			int(op.Y2*float32(t.Size().Y)),
		),
		color.Black,
		screen.Src,
	)
}

func (op BgRect) Update(state *TextureState) {
	state.backgroundRect = &op
}

type Figure struct {
	X float32
	Y float32
}

func (op Figure) Do(t screen.Texture) {
	ui.DrawCross(
		t,
		image.Pt(
			int(op.X*float32(t.Size().X)),
			int(op.Y*float32(t.Size().Y)),
		),
	)
}

func (op Figure) Update(state *TextureState) {
	state.figureCenters = append(state.figureCenters, &op)
}

type Move struct {
	X float32
	Y float32
}

func (op Move) Update(state *TextureState) {
	for _, fig := range state.figureCenters {
		fig.X = op.X
		fig.Y = op.Y
	}
}
