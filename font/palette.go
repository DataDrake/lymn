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

package font

import (
	"encoding/json"
	"fmt"
	"github.com/DataDrake/lymn/font/encoding"
	"log"
	"os"
	"text/tabwriter"
	"time"
)

// Palette is a color picker for a Palette of colors
type Palette struct {
	Name     string          `json:"name"`
	Author   string          `json:"author"`
	Date     time.Time       `json:"date"`
	Revision int             `json:"revision"`
	Colors   encoding.Colors `json:"colors"`
}

// DecodePalette decodes a Palette from an embedded JSON file and unmarshals it
func DecodePalette(raw []byte) (p *Palette) {
	p = &Palette{}
	if err := json.Unmarshal(raw, p); err != nil {
		log.Fatal(err.Error())
	}
	return
}

// Describe privides a simple description of the Palette with its Metadata
func (p *Palette) Describe() {
	tw := tabwriter.NewWriter(os.Stdout, 0, 4, 1, ' ', 0)
	fmt.Fprintf(tw, "Name\t: %s\n", p.Name)
	fmt.Fprintf(tw, "Author\t: %s\n", p.Author)
	fmt.Fprintf(tw, "Date\t: %s\n", p.Date)
	fmt.Fprintf(tw, "Revision\t: %d\n", p.Revision)
	fmt.Fprintf(tw, "Number of Colors\t: %d\n", len(p.Colors))
	tw.Flush()
}
