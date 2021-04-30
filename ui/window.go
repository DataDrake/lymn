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

// Window presents the entire game area to the player
type Window struct {
	title   *Titlebar
	text    *Textbox
	player  *Player
	choices *Choices

	grid *Grid
}

// NewWindow creates a fully populated game screen
func NewWindow(cols, rows int, glyphs *font.Font) *Window {
	top := rows - 7
	chX := cols/2 - 18/2
	return &Window{
		title:   NewTitlebar(0, 0, cols, cyan),
		text:    NewTextbox(0, 1, cols, top, cyan),
		player:  NewPlayer(1, top+2, color.White),
		choices: NewChoices(chX, top+2, 18, color.White),

		grid: NewGrid(cols, rows, glyphs),
	}
}

// DisplayStat adds a Player Stat to the Window
func (w *Window) DisplayStat(name string, max int) {
	w.player.AddStat(name, max)
}

// Draw renders the Window to the screen
func (w *Window) Draw(screen *ebiten.Image) {
	w.title.Draw(w.grid)
	w.text.Draw(w.grid)
	w.player.Draw(w.grid)
	w.choices.Draw(w.grid)

	w.grid.Draw(screen)
}

// Size gets the dimensions of the Window when rendered
func (w *Window) Size() (width, height int) {
	return w.grid.Size()
}
