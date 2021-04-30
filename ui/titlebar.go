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

// Titlebar is a topbar which prints both a Title and Location
type Titlebar struct {
	x, y   int
	cols   int
	bg     color.Color
	inited bool
}

// NewTitlebar creates a new empty titlebar of a specific color
func NewTitlebar(x, y, cols int, bg color.Color) (tb *Titlebar) {
	return &Titlebar{
		x:    x,
		y:    y,
		cols: cols,
		bg:   bg,
	}
}

func (tb *Titlebar) init(grid *Grid) {
	grid.Paint(tb.x, tb.y, tb.cols, 1, tb.bg)
	tb.inited = true
}

// Draw renders Titlebar to a screen
func (tb *Titlebar) Draw(grid *Grid) {
	if !tb.inited {
		tb.init(grid)
	}
	if !model.HasSceneChanged() {
		return
	}
	title, loc := model.GetScene()
	l := len(title) + len(loc)
	max := tb.cols - 2
	if l > max {
		grid.Text(tb.x+1, tb.y, len(title), 1, []rune(title), tb.bg, true)
	} else {
		// pad spaces to align Location right
		for i := max - l; i > 0; i-- {
			title += " "
		}
		title += loc
		grid.Text(tb.x+1, tb.y, len(title), 1, []rune(title), tb.bg, true)
	}
}
