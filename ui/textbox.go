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
	"github.com/DataDrake/ld48/model"
	"image/color"
)

var (
	// TextRows determines the amount of screen space allowed for text
	TextRows = Rows - 6
)

// Textbox is a bordered Box which prints text messages to the player
type Textbox struct {
	bg     color.Color
	inited bool
}

// NewTextbox creates a new empty textbox of a specific color
func NewTextbox(bg color.Color) (tb *Textbox) {
	return &Textbox{
		bg: bg,
	}
}

// init populates the stat UI elements
func (tb *Textbox) init(grid *Grid) {
	border := make([]rune, Rows-6)
	// left side
	for i := 0; i < TextRows; i++ {
		border[i] = '\u0001'
	}
	grid.Text(0, 1, 1, TextRows+1, border, tb.bg, false)
	// right side
	for i := 0; i < TextRows; i++ {
		border[i] = '\u0007'
	}
	grid.Text(Cols-1, 1, 1, TextRows+1, border, tb.bg, false)
	// bottom
	border = make([]rune, Cols)
	border[0] = '\u000B'
	border[Cols-1] = '\u000A'
	for i := 1; i < Cols-1; i++ {
		border[i] = '\u000D'
	}
	grid.Text(0, TextRows+1, Cols, 1, border, tb.bg, false)
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
	if model.TextPaused() {
		return
	}
	character, value := model.GetText()
	if len(character) > 0 {
		grid.Text(1, 2, len(character), 1, []rune(character), tb.bg, true)
		grid.Text(1, 4, Cols-2, TextRows-1, []rune(value), tb.bg, false)
	} else {
		grid.Text(1, 2, Cols-2, TextRows-1, []rune(value), tb.bg, false)
	}
}
