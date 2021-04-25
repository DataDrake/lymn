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
	"github.com/DataDrake/ld48/script"
	"sort"
)

// Runes is used to provide a sorted list of rune
type Runes []rune

// Len provides the length of the list
func (r Runes) Len() int {
	return len(r)
}

// Less returns true if a rune has a lower codepoint
func (r Runes) Less(i, j int) bool {
	return r[i] < r[j]
}

// Swap exchanges elements when sorting
func (r Runes) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// choices contains a list of the available choices
var choices struct {
	list    Runes
	changed bool
}

// SetChoices changes the list of available choices
func SetChoices(next script.Choices) {
	choices.list = make(Runes, len(next))
	for r := range next {
		choices.list = append(choices.list, r)
	}
	sort.Sort(choices.list)
	choices.changed = true
}

// HaveChoicesChanged checks if any changes have been made to the choices
func HaveChoicesChanged() bool {
	return choices.changed
}

// GetChoices returns the list of Choices and clears the changed attribute
func GetChoices() Runes {
	choices.changed = false
	return choices.list
}

// ClearChoices resets the choices list
func ClearChoices() {
	choices.list = make(Runes, 0)
	choices.changed = true
}

// choice keeps track of the selected choice
var choice struct {
	value rune
	set   bool
}

// SetChoice changes the selected choice
func SetChoice(value rune) {
	choice.value = value
	choice.set = true
}

// GetChoice returns the choice if it is set
func GetChoice() (value rune, ok bool) {
	if !choice.set {
		return
	}
	value = choice.value
	ok = true
	choice.set = false
	return
}
