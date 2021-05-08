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

// Condition represents a scripted examination of a character's stats
type Condition struct {
	Character string   `yaml:"character"`
	Stat      string   `yaml:"stat"`
	Type      CondKind `yaml:"type"`
	Value     int      `yaml:"value"`
}

// CondKind enumerates all of the available types of Condition checks
type CondKind string

const (
	// CondEqual checks for an exact match
	CondEqual CondKind = "equal"
	// CondLess checks for less than value
	CondLess CondKind = "less"
	// CondGreater checks for greater than value
	CondGreater CondKind = "greater"
)

// Conditional contains a list of Condition checks to be carried out
type Conditional []Condition
