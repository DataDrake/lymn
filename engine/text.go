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
	"github.com/DataDrake/lymn/model"
)

// printText sets the text to be displayed, pausinging if the next even involves text
func printText(e *Engine) bool {
	event := e.getEvent()
	model.SetText(event.Character, event.Text)
	e.event.index++
	event = e.getEvent()
	model.PauseText()
	e.state = waitText
	return true // Need to update the displayed text
}

// waitText holds back execution until the text is unpaused
func waitText(e *Engine) bool {
	if model.TextPaused() {
		return true // Need to wait before advancing to the next Text
	}
	e.state = decodeEvent
	return false
}
