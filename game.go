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

package lymn

import (
	"fmt"
	"github.com/DataDrake/lymn/engine"
	"github.com/DataDrake/lymn/font"
	"github.com/DataDrake/lymn/input"
	"github.com/DataDrake/lymn/script"
	"github.com/DataDrake/lymn/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"time"
)

// Game is the global Game object
type Game struct {
	script script.Script
	window *ui.Window
	engine *engine.Engine
	last   time.Time
}

// NewGame creates a Game and populates the UI
func NewGame(s script.Script, f *font.Font) (g *Game) {
	g = &Game{
		script: s,
		window: ui.NewWindow(40, 20, f),
		engine: engine.NewEngine(s),
		last:   time.Now(),
	}
	return
}

// Update checks for all updates in the input and the internal state of the Game
func (g *Game) Update() error {
	g.last = time.Now()
	input.Update()
	g.engine.Update()
	return nil
}

// DisplayStat adds a Player stat tp the Window
func (g *Game) DisplayStat(name string, max int) {
	g.window.DisplayStat(name, max)
}

// Draw renders Game to a screen
func (g *Game) Draw(screen *ebiten.Image) {
	g.window.Draw(screen)
	elapsed := time.Now().Sub(g.last)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Time: %0.3fms", elapsed.Seconds()*1000), 8, 620)
}

// Layout specifies the layout dimensions of Game
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.window.Size()
}

// Play starts the game
func (g *Game) Play() {
	ebiten.SetWindowSize(g.window.Size())
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle(g.script.Title)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
