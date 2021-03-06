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

import (
	"gopkg.in/yaml.v3"
	"log"
)

// Script provides all of the information needed for the Engine to run
type Script struct {
	Title  string `yaml:"title"`
	Author string `yaml:"author"`
	Date   string `yaml:"date"`
	Start  string `yaml:"start"`
	Scenes Scenes `yaml:"scenes"`
}

// Decode parses an embedded Script from raw bytes
func Decode(raw []byte) (s Script) {
	if err := yaml.Unmarshal(raw, &s); err != nil {
		log.Fatalf("Failed to decode script: %s\n", err)
	}
	return
}
