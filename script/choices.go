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

// ChoiceKind helps enumerate the supported types of Choice
type ChoiceKind string

const (
	// ChoiceScene is a Choice which transitions to a different scene
	ChoiceScene ChoiceKind = "scene"
	// ChoiceEvent is a Choice which transitions to a different group of events
	ChoiceEvent ChoiceKind = "event"
	// ChoiceQuit is a Choice which exits the game
	ChoiceQuit ChoiceKind = "quit"
)

// Choice represents a scripted Choice theat a Player made
type Choice struct {
	Type  ChoiceKind `yaml:"type"`
	Value string     `yaml:"value"`
}

// Choices relates the rune used to display the Choice to the Choice it selects
type Choices map[rune]Choice
