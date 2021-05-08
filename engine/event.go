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
)

// getEvents gets the current Event group from the current Scene
func (e *Engine) getEvents() []script.Event {
	scene := e.getScene()
	events, ok := scene.Events[e.event.curr]
	if !ok {
		log.Fatalf("missing event group: %s\n", e.event.curr)
	}
	return events
}

// getEvent gets the current Event from the current Event group
func (e *Engine) getEvent() script.Event {
	events := e.getEvents()
	if e.event.index >= len(events) {
		log.Fatalf("missing event: %s[%d]\n", e.event.curr, e.event.index)
	}
	return events[e.event.index]
}

// loadEvent moves to the next Event group
func loadEvent(e *Engine) bool {
	e.event.curr = e.event.next
	e.event.index = 0
	e.state = decodeEvent
	return false
}

// decodeEvent executes the current Event
func decodeEvent(e *Engine) bool {
	event := e.getEvent()
	if !model.CheckStats(event.Condition) {
		e.event.index++
		return false
	}
	switch event.Type {
	case script.EventText:
		e.state = printText
	case script.EventScene:
		e.scene.next = event.Scene
		e.state = loadScene
	case script.EventEvent:
		e.event.next = event.Event
		e.state = loadEvent
	case script.EventChoice:
		e.state = printChoices
	case script.EventStats:
		e.state = updateStats
	default:
		log.Fatalf("invalid event type: %s[%d] %s\n", e.event.curr, e.event.index, event.Type)
	}
	return false
}
