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

// Statbar is a labelled progressbar that represents some Player stat quantity
type Statbar struct {
	x, y   int
	stat   string
	max    int
	value  int
	label  color.Color
	inited bool
}

// NewStatbar creates a new filled Statbar
func NewStatbar(x, y int, stat string, max int, label color.Color) (sb *Statbar) {
	return &Statbar{
		x:     x,
		y:     y,
		stat:  stat,
		max:   max,
		label: label,
	}
}

func (sb *Statbar) init(grid *Grid) {
	grid.Text(sb.x, sb.y, len(sb.stat), 1, []rune(sb.stat), sb.label, false)
	bar := []rune{2, 1, 1, 1, 1, 1, 1, 1, 1, 3}
	grid.Text(sb.x+len(sb.stat)+1, sb.y, len(bar), 1, bar, sb.label, false)
	sb.inited = true
}

// Draw renders Statbar to a screen
func (sb *Statbar) Draw(grid *Grid) {
	if !sb.inited {
		sb.init(grid)
	}
	stats, ok := model.GetStats("Player")
	if !ok {
		return
	}
	value := stats[sb.stat]
	if sb.value == value {
		return
	}
	sb.value = value
	percent := float64(value*100) / float64(sb.max)
	bar := make([]rune, 10)
	idx := 0
	for i := 10; i <= 100; i += 10 {
		part := float64(i) - percent
		switch {
		case part <= 0.0:
			bar[idx] = 7
		case part <= 2.5:
			bar[idx] = 6
		case part <= 5.0:
			bar[idx] = 5
		case part <= 7.5:
			bar[idx] = 4
		default:
			println(idx)
			switch idx {
			case 0:
				bar[idx] = 2
			case 9:
				bar[idx] = 3
			default:
				bar[idx] = 1
			}
		}
		idx++
	}
	var c color.Color
	switch {
	case percent <= 25:
		c = color.RGBA{0xFF, 0x00, 0x00, 0xFF}
	case percent <= 50:
		c = color.RGBA{0xFF, 0x99, 0x00, 0xFF}
	case percent <= 75:
		c = color.RGBA{0xFF, 0xFF, 0x00, 0xFF}
	default:
		c = color.RGBA{0x00, 0xFF, 0x00, 0xFF}
	}
	grid.Text(sb.x+len(sb.stat)+1, sb.y, len(bar), 1, bar, c, false)
}
