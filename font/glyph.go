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
	"github.com/DataDrake/ld48/font/encoding"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

// Glyph holds the data for a single glyph
type Glyph struct {
	src *image.Paletted
	img *ebiten.Image
}

// Bounds returns the bounding box for this Glyph
func (g *Glyph) Bounds() image.Rectangle {
	return g.src.Bounds()
}

// Image returns the internal image as an ebiten image
func (g *Glyph) Image() *ebiten.Image {
	return g.img
}

// SetPalette changes the color palette for this Glyph
func (g *Glyph) SetPalette(colors *Palette) {
	g.src.Palette = color.Palette(colors.Colors)
	g.img = ebiten.NewImageFromImage(g.src)
}

// UnmarshalJSON is a custom unmarshaler for the Glyph type
func (g *Glyph) UnmarshalJSON(b []byte) (err error) {
	var j encoding.Sprite
	if err = json.Unmarshal(b, &j); err != nil {
		return
	}
	(*g).src, err = j.Image()
	return
}
