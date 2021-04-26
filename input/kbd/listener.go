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

package kbd

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Listener is a function to call in response to a Key being pressed
type Listener func(key ebiten.Key)

// listeners contains all of the currently Registered listeners
var listeners = make(map[ebiten.Key]Listener)

// Register associates a KeyListener with a specific Key
func Register(key ebiten.Key, listener Listener) {
	listeners[key] = listener
}

// Update scans for currently pressed keys and debounces input
func Update() {
	curr := make(state)
	// currently pressed Keys
	for key := range listeners {
		if ebiten.IsKeyPressed(key) {
			curr[key] = true
		}
	}
	// execute listeners
	for _, key := range pressed(curr) {
		listeners[key](key)
	}
	// save state
	prev = curr
}
