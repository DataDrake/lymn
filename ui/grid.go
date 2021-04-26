//
// Copyright 2021 Bryan T. Meyers <root@datadrake.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package ui

import (
	"github.com/DataDrake/ld48/font"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

const (
	// Cols specifies the width of the display in characters
	Cols = 40
	// Rows specifies the height of the dispaly in characters
	Rows = 20
	// XMargin specifies the horizontal margin built into the Font Glyphs
	XMargin = 2
)

// Grid is a 2D matrix of text, rendered to an Image
type Grid struct {
	img *ebiten.Image
}

// NewGrid creates an empty Grid of max size
func NewGrid() (g *Grid) {
	width := 2 * Cols * (font.Default.Size() - (2 * XMargin))
	height := 2 * Rows * font.Default.Size()
	g = &Grid{
		img: ebiten.NewImage(width, height),
	}
	return
}

// Draw renders Game to a screen
func (g *Grid) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.img, nil)
}

func (g *Grid) calcXY(x, y int) (ix, iy, iw, ih int) {
	iw = 2 * (font.Default.Size() - (2 * XMargin))
	ih = 2 * font.Default.Size()
	ix = x*iw - (2 * XMargin)
	iy = y * ih
	return
}

// Paint replaces a rectangular region of the screen with the specified color
func (g *Grid) Paint(x, y, cols, rows int, bg color.Color) {
	x, y, xInc, yInc := g.calcXY(x, y)
	background := ebiten.NewImage(cols*xInc, rows*yInc)
	background.Fill(bg)
	op := &ebiten.DrawImageOptions{
		CompositeMode: ebiten.CompositeModeCopy,
	}
	op.GeoM.Translate(float64(x+2*XMargin), float64(y))
	g.img.DrawImage(background, op)
}

// Text replaces a rectangular region of the screen with the specified runes
func (g *Grid) Text(x, y, cols, rows int, data []rune, fg color.Color, inverted bool) {
	g.Paint(x, y, cols, rows, fg)
	xStart, y, xInc, yInc := g.calcXY(x, y)
	x = xStart
	f := font.Default
	mode := ebiten.CompositeModeDestinationIn
	if inverted {
		mode = ebiten.CompositeModeDestinationOut
	}
	for col, r := range data {
		if col != 0 && (col%cols) == 0 {
			x = xStart
			y += yInc
		}
		op := &ebiten.DrawImageOptions{
			CompositeMode: mode,
		}
		op.GeoM.Scale(2, 2)
		op.GeoM.Translate(float64(x), float64(y))
		g.img.DrawImage(f.Glyphs[r].Image(), op)
		x += xInc
	}
}

// Size gets the dimensions of the Grid when rendered
func (g *Grid) Size() (width, height int) {
	bounds := g.img.Bounds()
	width = bounds.Dx()
	height = bounds.Dy()
	return
}
