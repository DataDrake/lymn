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
	"github.com/DataDrake/ld48/model"
	"github.com/DataDrake/ld48/script"
	"log"
)

// getScene gets the current Scene from the script
func (e *Engine) getScene() script.Scene {
	scene, ok := e.script.Scenes[e.scene.curr]
	if !ok {
		log.Fatalf("missing scene: %s\n", e.scene.curr)
	}
	return scene
}

// loadScene changes the current Scene and notifies the UI
func loadScene(e *Engine) bool {
	e.scene.curr = e.scene.next
	scene := e.getScene()
	e.event.next = scene.Start
	model.SetScene(scene.Title, scene.Location)
	e.state = loadEvent
	return true // blocked because the UI needs to update
}
