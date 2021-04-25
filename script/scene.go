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

package script

// Scene represents a narrative location where a sequents of events is carries out
type Scene struct {
	Location string `yaml:"location"`
	Start    string `yaml:"start"`
	Events   Events `yaml:"events"`
}

// Scenes maps a string ID to a specific scene
type Scenes map[string]Scene
