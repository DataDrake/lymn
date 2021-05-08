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
	"github.com/DataDrake/lymn/font"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

// Grid is a 2D matrix of text, rendered to an Image
type Grid struct {
	font *font.Font
	img  *ebiten.Image
}

// NewGrid creates an empty Grid of max size
func NewGrid(cols, rows int, glyphs *font.Font) (g *Grid) {
	width, height := glyphs.Size()
	width *= 2 * cols
	height *= 2 * rows
	g = &Grid{
		font: glyphs,
		img:  ebiten.NewImage(width, height),
	}
	return
}

// Draw renders Game to a screen
func (g *Grid) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.img, nil)
}

func (g *Grid) calcXY(x, y int) (ix, iy, iw, ih int) {
	w, h := g.font.Size()
	iw = 2 * w
	ih = 2 * h
	ix = x * iw
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
	op.GeoM.Translate(float64(x), float64(y))
	g.img.DrawImage(background, op)
}

// Text replaces a rectangular region of the screen with the specified runes
func (g *Grid) Text(x, y, cols, rows int, data []rune, fg color.Color, inverted bool) {
	g.Paint(x, y, cols, rows, fg)
	xStart, y, xInc, yInc := g.calcXY(x, y)
	x = xStart
	mode := ebiten.CompositeModeDestinationIn
	if inverted {
		mode = ebiten.CompositeModeDestinationOut
	}
	var r rune
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			op := &ebiten.DrawImageOptions{
				CompositeMode: mode,
			}
			op.GeoM.Scale(2, 2)
			op.GeoM.Translate(float64(x), float64(y))
			glyph := g.font.Glyphs[0]
			if len(data) > 0 && r != '\n' {
				r, data = data[0], data[1:]
			} else {
				r = '\n'
			}
			if r != '\n' {
				glyph = g.font.Glyphs[r]
			}
			g.img.DrawImage(glyph.Image(), op)
			x += xInc
		}
		r = 0
		x = xStart
		y += yInc
	}
}

// Size gets the dimensions of the Grid when rendered
func (g *Grid) Size() (width, height int) {
	bounds := g.img.Bounds()
	width = bounds.Dx()
	height = bounds.Dy()
	return
}
