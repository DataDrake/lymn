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

package main

import (
	_ "embed" // needed for embeddable data
	"github.com/DataDrake/ld48"
	"github.com/DataDrake/ld48/font"
	"github.com/DataDrake/ld48/script"
)

//go:embed script.yml
var rawScript []byte

//go:embed palette.json
var rawPalette []byte

//go:embed font.json
var rawFont []byte

func main() {
	s := script.Decode(rawScript)
	p := font.DecodePalette(rawPalette)
	p.Describe()
	f := font.Decode(rawFont)
	f.Describe()
	f.SetPalette(p)
	f.Render()
	g := ld48.NewGame(s, f)
	g.DisplayStat("HP", 100)
	g.DisplayStat("SP", 100)
	g.DisplayStat("MP", 100)
	g.Play()
}
