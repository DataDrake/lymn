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
	"log"
)

// Choices is a wheel that prints available choices
type Choices struct {
	x, y int
	cols int
	bg   color.Color
}

// NewChoices creates a new empty Choices of a specific color
func NewChoices(x, y, cols int, bg color.Color) (cs *Choices) {
	return &Choices{
		x:    x,
		y:    y,
		cols: cols,
		bg:   bg,
	}
}

// Draw renders Choices to a screen
func (cs *Choices) Draw(grid *Grid) {
	if !model.HaveChoicesChanged() {
		return
	}
	choices := model.GetChoices()
	rows := make([]string, 3)
	switch len(choices) {
	case 4:
		choice := choices[3]
		for margin := (cs.cols - len(choice)) / 2; margin > 0; margin-- {
			rows[2] += " "
		}
		rows[2] += choice
		fallthrough
	case 3:
		choice := choices[2]
		for margin := (cs.cols - len(choice)) / 2; margin > 0; margin-- {
			rows[0] += " "
		}
		rows[0] += choice
		fallthrough
	case 2:
		rows[1] = choices[0] + "  " + choices[1]
		for margin := (cs.cols/2 - 1) - len(choices[0]); margin > 0; margin-- {
			rows[1] = " " + rows[1]
		}
	case 1:
		rows[1] += choices[0]
		for margin := (cs.cols - len(choices[0])) / 2; margin > 0; margin-- {
			rows[1] += " "
		}
	case 0:
		// clear
	default:
		log.Fatal("Only 1-4 choices can be displayed")
	}
	for i, row := range rows {
		grid.Text(cs.x, cs.y+i, cs.cols, 1, []rune(row), cs.bg, false)
	}
}
