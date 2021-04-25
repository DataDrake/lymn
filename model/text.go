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

// text contains text to be shared between UI and Engine
var text struct {
	character string
	value     string
	changed   bool
	paused    bool
}

// SetText changes the current text
func SetText(character, value string) {
	text.character = character
	text.value = value
	text.changed = true
}

// HasTextChanged checks if any changes have been made to the text
func HasTextChanged() bool {
	return text.changed
}

// GetText returns the current text information
func GetText() (character, value string) {
	text.changed = false
	return text.character, text.value
}

// PauseText prevents the state machine from advancing while more text is available
func PauseText() {
	text.paused = true
}

// TextPaused checks if the text is paused
func TextPaused() bool {
	return text.paused
}

// UnpauseText clear a pause on text rendering
func UnpauseText() {
	text.paused = true
}
