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

import (
	"github.com/DataDrake/lymn/script"
)

// choices contains a list of the available choices
var choices struct {
	list    []string
	changed bool
}

// SetChoices changes the list of available choices
func SetChoices(next script.Choices) {
	choices.list = make([]string, 0)
	for _, choice := range next {
		choices.list = append(choices.list, choice.Name)
	}
	choices.changed = true
}

// HaveChoicesChanged checks if any changes have been made to the choices
func HaveChoicesChanged() bool {
	return choices.changed
}

// GetChoices returns the list of Choices and clears the changed attribute
func GetChoices() []string {
	choices.changed = false
	return choices.list
}

// ClearChoices resets the choices list
func ClearChoices() {
	choices.list = make([]string, 0)
	choices.changed = true
}

// choice keeps track of the selected choice
var choice struct {
	index int
	set   bool
}

// SetChoice changes the selected choice
func SetChoice(index int) {
	choice.index = index
	choice.set = true
}

// GetChoice returns the choice if it is set
func GetChoice() (index int, ok bool) {
	if !choice.set {
		return
	}
	index = choice.index
	ok = true
	choice.set = false
	return
}
