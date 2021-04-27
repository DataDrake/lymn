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
	_ "embed" // required for embedding
	"encoding/json"
	"fmt"
	"image"
	"log"
	"os"
	"text/tabwriter"
	"time"
)

// XMargin sets a per glyph clipping of Width/4 on the left and right of a glyph
var XMargin int

// Default is the Font to use for rendering things
var Default *Font

//go:embed font.json
var rawFont []byte

func init() {
	var err error
	Default, err = NewFont(rawFont)
	if err != nil {
		log.Fatal(err)
	}
	Default.Describe()
	println()
	LoadColors()
	Default.SetPalette(Colors)
	Default.Render()
}

// Font represents one or more glyphs belonging to a single face
type Font struct {
	Name     string    `json:"name"`
	Author   string    `json:"author"`
	Date     time.Time `json:"date"`
	Revision int       `json:"revision"`
	Glyphs   []*Glyph  `json:"sprites"`
	width    int
	height   int
}

// NewFont decodes a Font from a embedded JSON file
func NewFont(data []byte) (f *Font, err error) {
	f = &Font{}
	err = json.Unmarshal(data, f)
	f.width = f.Glyphs[0].src.Stride
	f.height = f.width
	return
}

// SetPalette sets the color Palette for the Glyphs in this Font
func (f *Font) SetPalette(p *Palette) {
	for _, g := range f.Glyphs {
		g.SetPalette(p)
	}
}

// Render generates images for each Glyph to be later used in the rendering process
func (f *Font) Render() {
	f.width = f.height
	XMargin = f.height / 8
	crop := image.Rect(XMargin, 0, f.width-XMargin, f.height)
	for _, g := range f.Glyphs {
		g.Render(crop)
	}
	f.width = f.height - 2*XMargin
}

// Describe summarizes a Font according to its metadata
func (f *Font) Describe() {
	tw := tabwriter.NewWriter(os.Stdout, 0, 4, 1, ' ', 0)
	fmt.Fprintf(tw, "Name\t: %s\n", f.Name)
	fmt.Fprintf(tw, "Author\t: %s\n", f.Author)
	fmt.Fprintf(tw, "Date\t: %s\n", f.Date)
	fmt.Fprintf(tw, "Revision\t: %d\n", f.Revision)
	fmt.Fprintf(tw, "Number of Glyphs\t: %d\n", len(f.Glyphs))
	tw.Flush()
}

// Size returns the square dimension of each glyph
func (f *Font) Size() (width, height int) {
	return f.width, f.height
}
