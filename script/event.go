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

// EventKind helps enumerate all of the possible types of Event
type EventKind string

const (
	// EventText is an event that presents text to the Player
	EventText EventKind = "text"
	// EventScene is an event that changes the current scene
	EventScene EventKind = "scene"
	// EventEvent is an event that branches to a specific group of events within the current scene
	EventEvent EventKind = "event"
	// EventChoice is an event that presents a choice to the Player
	EventChoice EventKind = "choice"
	// EventStats is an event which alters the Stats of a Player or NPC
	EventStats EventKind = "stats"
)

// Event is one of a sequence of possible actions which occur during a Scene
type Event struct {
	Type      EventKind    `yaml:"type"`
	Character string       `yaml:"character"`
	Text      string       `yaml:"text"`
	Scene     string       `yaml:"scene"`
	Event     string       `yaml:"event"`
	Choices   Choices      `yaml:"choices"`
	Stats     Stats        `yaml:"stats"`
	Condition *Conditional `yaml:"condition"`
}

// Events maps a string ID to a sequential group of events
type Events map[string][]Event
