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
	"log"
)

// characters contains all of the stats for characters
var characters = make(map[string]Stats)

// GetStats retrieves the stats of a character by name
func GetStats(character string) (m Stats, ok bool) {
	m, ok = characters[character]
	return
}

// UpdateStats applies a list of stat changes
func UpdateStats(stats script.Stats) {
	for _, stat := range stats {
		m, ok := characters[stat.Character]
		if !ok {
			m = make(Stats)
			characters[stat.Character] = m
		}
		m.Apply(stat)
	}
}

// Stats stores all of the stats related to a character
type Stats map[string]int

// Apply modifies the stats of a character with a scripted action
func (s Stats) Apply(stat script.Stat) {
	switch stat.Type {
	case script.StatClear:
		delete(s, stat.Name)
	case script.StatSet:
		s[stat.Name] = stat.Value
	case script.StatAdd:
		s[stat.Name] += stat.Value
	case script.StatSub:
		s[stat.Name] -= stat.Value
	default:
		log.Fatalf("Invalid stat type: %s\n", stat.Type)
	}
}
