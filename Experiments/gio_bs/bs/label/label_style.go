package label

import (
	"image/color"

	"gioui.org/font"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
)

// LabelStyle configures the presentation of text. If the State field is set, the
// label will be laid out as interactive (able to be selected and copied). Otherwise,
// the label will be non-interactive.
type LabelStyle struct {
	// Face defines the text style.
	Font font.Font
	// Color is the text color.
	Color color.NRGBA
	// SelectionColor is the color of the background for selected text.
	SelectionColor color.NRGBA
	// Alignment specify the text alignment.
	Alignment text.Alignment
	// MaxLines limits the number of lines. Zero means no limit.
	MaxLines int
	// WrapPolicy configures how displayed text will be broken into lines.
	WrapPolicy text.WrapPolicy
	// Truncator is the text that will be shown at the end of the final
	// line if MaxLines is exceeded. Defaults to "…" if empty.
	Truncator string
	// Text is the content displayed by the label.
	Text string
	// TextSize determines the size of the text glyphs.
	TextSize unit.Sp
	// LineHeight controls the distance between the baselines of lines of text.
	// If zero, a sensible default will be used.
	LineHeight unit.Sp
	// LineHeightScale applies a scaling factor to the LineHeight. If zero, a
	// sensible default will be used.
	LineHeightScale float32

	// Shaper is the text shaper used to display this labe. This field is automatically
	// set using by all constructor functions. If constructing a LabelStyle literal, you
	// must provide a Shaper or displaying text will panic.
	Shaper *text.Shaper
	// State provides text selection state for the label. If not set, the label cannot
	// be selected or copied interactively.
	State *widget.Selectable
}
