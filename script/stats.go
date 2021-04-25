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

// Stat represents a scripted modification to a character's stats
type Stat struct {
	Name      string   `yaml:"name"`
	Character string   `yaml:"character"`
	Type      StatKind `yaml:"type"`
	Value     int      `yaml:"value"`
}

// StatKind enumerates all of the available types of Stat modification
type StatKind string

const (
	// StatClear removes a Stat and can be used to clear status effects
	StatClear StatKind = "clear"
	// StatSet overwrites the value of a Stat
	StatSet StatKind = "set"
	// StatAdd increments the value of a Stat by the specified number
	StatAdd StatKind = "add"
	// StatSub decrements the value of a Stat be the specified number
	StatSub StatKind = "sub"
)

// Stats contains a list of Stat operations to be carried out
type Stats []Stat
