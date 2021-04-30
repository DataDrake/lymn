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

// updateStats performs stat modifications and then forces a UI refresh
func updateStats(e *Engine) bool {
	event := e.getEvent()
	model.UpdateStats(event.Stats)
	e.event.index++
	e.state = decodeEvent
	return true // Need to update any displayed stats
}
