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
	"github.com/DataDrake/lymn/model"
	"image/color"
)

// Textbox is a bordered Box which prints text messages to the player
type Textbox struct {
	x, y   int
	cols   int
	rows   int
	bg     color.Color
	inited bool
	paused bool
}

// NewTextbox creates a new empty textbox of a specific color
func NewTextbox(x, y, cols, rows int, bg color.Color) (tb *Textbox) {
	return &Textbox{
		x:    x,
		y:    y,
		cols: cols,
		rows: rows,
		bg:   bg,
	}
}

// init populates the stat UI elements
func (tb *Textbox) init(grid *Grid) {
	border := make([]rune, tb.rows)
	// left side
	for i := 0; i < tb.rows; i++ {
		border[i] = 12
	}
	grid.Text(tb.x, tb.y, 1, tb.rows, border, tb.bg, false)
	// right side
	for i := 0; i < tb.rows; i++ {
		border[i] = 14
	}
	grid.Text(tb.x+tb.cols-1, tb.y, 1, tb.rows, border, tb.bg, false)
	// bottom
	border = make([]rune, tb.cols)
	border[0] = 8
	border[tb.cols-1] = 9
	for i := 1; i < tb.cols-1; i++ {
		border[i] = 13
	}
	grid.Text(tb.x, tb.rows, tb.cols, 1, border, tb.bg, false)
	tb.inited = true
}

// Draw renders Textbox to a screen
func (tb *Textbox) Draw(grid *Grid) {
	if !tb.inited {
		tb.init(grid)
	}
	if !model.HasTextChanged() {
		return
	}
	rows := tb.rows
	if model.TextPaused() {
		if tb.paused {
			return
		}
		grid.Text(tb.cols-4, rows-1, 3, 1, []rune("..."), tb.bg, false)
		rows -= 2
		tb.paused = true
	}
	tb.paused = false
	grid.Paint(tb.x+1, tb.y+1, tb.cols-2, rows-2, color.Black)
	character, value := model.GetText()
	if len(character) > 0 {
		character = " " + character + " "
		grid.Text(tb.x+1, tb.y+1, len(character), 1, []rune(character), tb.bg, true)
		grid.Text(tb.x+1, tb.y+3, tb.cols-2, rows-4, []rune(value), tb.bg, false)
	} else {
		grid.Text(tb.x+1, tb.y+1, tb.cols-2, rows-2, []rune(value), tb.bg, false)
	}
}
