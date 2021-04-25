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

package model

// scene contains the scene data to be shared between UI and Engine
var scene struct {
	title    string
	location string
	changed  bool
}

// SetScene changes the scene
func SetScene(title, location string) {
	scene.title = title
	scene.location = location
	scene.changed = true
}

// HasSceneChanged checks if any changes have been made to the scene
func HasSceneChanged() bool {
	return scene.changed
}

// GetScene returns the current Scene and clears the changed status
func GetScene() (title, location string) {
	scene.changed = false
	return scene.title, scene.location
}
