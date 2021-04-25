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

package engine

import (
	"github.com/DataDrake/ld48/script"
)

// Engine carries out all of the game logic according to a Script
type Engine struct {
	state  stateFn
	script script.Script
	scene  struct {
		curr string
		next string
	}
	event struct {
		curr  string
		next  string
		index int
	}
	choice rune
}

// NewEngine creates an Engine that is ready to go
func NewEngine(s script.Script) *Engine {
	return &Engine{
		script: s,
		state:  start,
	}
}

// Update carries out the internal state machine until it blocks on model changes or user input
func (e *Engine) Update() {
	for {
		if e.state(e) {
			// blocking on UI/model update
			return
		}
	}
}

// start sets the scene to the beginning of the Script and starts the Loading process
func start(e *Engine) bool {
	e.scene.next = e.script.Start
	e.state = loadScene
	return false
}
