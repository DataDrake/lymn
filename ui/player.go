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
	"image/color"
)

// Player presents player stats to the player
type Player struct {
	x, y   int
	label  color.Color
	stats  []*Statbar
	inited bool
}

// NewPlayer creates a new filled Player
func NewPlayer(x, y int, label color.Color) *Player {
	return &Player{
		x:     x,
		y:     y,
		label: label,
	}
}

// Draw renders a Player
func (p *Player) Draw(grid *Grid) {
	for _, stat := range p.stats {
		stat.Draw(grid)
	}
}

// AddStat registers a stat for the Player to render
func (p *Player) AddStat(name string, max int) {
	y := p.y + len(p.stats)
	sb := NewStatbar(p.x, y, name, max, p.label)
	p.stats = append(p.stats, sb)
}
