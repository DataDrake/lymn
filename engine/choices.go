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
	"github.com/DataDrake/lymn/script"
	"log"
	"os"
)

// getChoices gets the choices provided by the current event
func (e *Engine) getChoices() script.Choices {
	return e.getEvent().Choices
}

// printChoices sets the choices to be displayed
func printChoices(e *Engine) bool {
	model.SetChoices(e.getChoices())
	e.state = waitChoice
	return true // Need to display choices to Player
}

// waitChoice blocks until a choice has been made
func waitChoice(e *Engine) bool {
	var ok bool
	e.choice, ok = model.GetChoice()
	if !ok {
		return true // Need to give another chance for input
	}
	e.state = decodeChoice
	return false
}

// decodeChoice executes the current Choice
func decodeChoice(e *Engine) bool {
	choices := e.getChoices()
	if e.choice > len(choices) {
		log.Fatalf("missing choice: %d\n", e.choice)
	}
	choice := choices[e.choice]
	model.ClearChoices()
	e.event.index++
	switch choice.Type {
	case script.ChoiceScene:
		e.scene.next = choice.Value
		e.state = loadScene
	case script.ChoiceEvent:
		e.event.next = choice.Value
		e.state = loadEvent
	case script.ChoiceQuit:
		os.Exit(0)
	default:
		log.Fatalf("invalid choice type: %s\n", choice.Type)
	}
	return false
}
