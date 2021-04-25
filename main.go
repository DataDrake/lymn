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

package main

import (
	"fmt"
	"github.com/DataDrake/ld48/engine"
	"github.com/DataDrake/ld48/script"
	"github.com/DataDrake/ld48/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"time"
)

// Game is the global Game object
type Game struct {
	grid   *ui.Grid
	engine *engine.Engine
	last   time.Time
}

// NewGame creates a Game and populates the UI
func NewGame() (g *Game) {
	g = &Game{
		grid:   ui.NewGrid(),
		engine: engine.NewEngine(script.Load()),
		last:   time.Now(),
	}
	g.grid.Set(1, 1, 2, 2, []rune("AbcD"), false)
	g.grid.Set(0, 0, 1, ui.Cols, []rune(" AbcDEFGHIJKLMNOPQRSTUVWXYZ"), true)
	g.grid.Set(10, 18, 2, 2, []rune("AbcD"), true)
	g.grid.Set(30, 2, 2, 2, []rune("AbcD"), false)
	return
}

// Update checks for all updates in the input and the internal state of the Game
func (g *Game) Update() error {
	g.last = time.Now()
	// input.Update()
	return nil
}

// Draw renders Game to a screen
func (g *Game) Draw(screen *ebiten.Image) {
	g.grid.Draw(screen)
	elapsed := time.Now().Sub(g.last)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Time: %0.3fms", elapsed.Seconds()*1000), 8, 256)
}

// Layout specifies the layout dimensions of Game
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.grid.Size()
}

func main() {
	g := NewGame()
	ebiten.SetWindowSize(g.Layout(0, 0))
	ebiten.SetWindowResizable(false)
	ebiten.SetWindowTitle("LD48")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
