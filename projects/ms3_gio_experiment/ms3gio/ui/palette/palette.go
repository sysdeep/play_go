package palette

import (
	"image/color"
	"strconv"
	"strings"
)

// BlueGray = PaletteModel("#F8FAFC", "#F1F5F9", "#E2E8F0", "#CBD5E1", "#94A3B8", "#64748B", "#475569", "#334155", "#1E293B", "#0F172A")

// var (
// 	BlueGray []string = []string{"#F8FAFC", "#F1F5F9", "#E2E8F0", "#CBD5E1", "#94A3B8", "#64748B", "#475569", "#334155", "#1E293B", "#0F172A"}
// )

var palette [][]string = [][]string{
	// BlueGray
	{"#F8FAFC", "#F1F5F9", "#E2E8F0", "#CBD5E1", "#94A3B8", "#64748B", "#475569", "#334155", "#1E293B", "#0F172A", "#020617"},
	// CoolGray
	{"#F9FAFB", "#F3F4F6", "#E5E7EB", "#D1D5DB", "#9CA3AF", "#6B7280", "#4B5563", "#374151", "#1F2937", "#111827", "#030712"},
	// Gray
	{"#FAFAFA", "#F4F4F5", "#E4E4E7", "#D4D4D8", "#A1A1AA", "#71717A", "#52525B", "#3F3F46", "#27272A", "#18181B", "#09090b"},
	// TrueGray
	{"#FAFAFA", "#F5F5F5", "#E5E5E5", "#D4D4D4", "#A3A3A3", "#737373", "#525252", "#404040", "#262626", "#171717", "#0a0a0a"},
	// WarmGray
	{"#FAFAF9", "#F5F5F4", "#E7E5E4", "#D6D3D1", "#A8A29E", "#78716C", "#57534E", "#44403C", "#292524", "#1C1917", "#0c0a09"},
	//
	// Red
	{"#FEF2F2", "#FEE2E2", "#FECACA", "#FCA5A5", "#F87171", "#EF4444", "#DC2626", "#B91C1C", "#991B1B", "#7F1D1D", "#450a0a"},
	// Orange
	{"#FFF7ED", "#FFEDD5", "#FED7AA", "#FDBA74", "#FB923C", "#F97316", "#EA580C", "#C2410C", "#9A3412", "#7C2D12", "#431407"},
	// Amber
	{"#FFFBEB", "#FEF3C7", "#FDE68A", "#FCD34D", "#FBBF24", "#F59E0B", "#D97706", "#B45309", "#92400E", "#78350F", "#451a03"},
	// Yellow
	{"#FEFCE8", "#FEF9C3", "#FEF08A", "#FDE047", "#FACC15", "#EAB308", "#CA8A04", "#A16207", "#854D0E", "#713F12", "#422006"},
	// Lime
	{"#F7FEE7", "#ECFCCB", "#D9F99D", "#BEF264", "#A3E635", "#84CC16", "#65A30D", "#4D7C0F", "#3F6212", "#365314", "#1a2e05"},
	// Green
	{"#F0FDF4", "#DCFCE7", "#BBF7D0", "#86EFAC", "#4ADE80", "#22C55E", "#16A34A", "#15803D", "#166534", "#14532D", "#052e16"},
	// Emerald
	{"#ECFDF5", "#D1FAE5", "#A7F3D0", "#6EE7B7", "#34D399", "#10B981", "#059669", "#047857", "#065F46", "#064E3B", "#022c22"},
	// Teal
	{"#F0FDFA", "#CCFBF1", "#99F6E4", "#5EEAD4", "#2DD4BF", "#14B8A6", "#0D9488", "#0F766E", "#115E59", "#134E4A", "#042f2e"},
	// Cyan
	{"#ECFEFF", "#CFFAFE", "#A5F3FC", "#67E8F9", "#22D3EE", "#06B6D4", "#0891B2", "#0E7490", "#155E75", "#164E63", "#083344"},
	// LightBlue
	{"#F0F9FF", "#E0F2FE", "#BAE6FD", "#7DD3FC", "#38BDF8", "#0EA5E9", "#0284C7", "#0369A1", "#075985", "#0C4A6E", "#082f49"},
	// Blue
	{"#EFF6FF", "#DBEAFE", "#BFDBFE", "#93C5FD", "#60A5FA", "#3B82F6", "#2563EB", "#1D4ED8", "#1E40AF", "#1E3A8A", "#172554"},
	// Indigo
	{"#EEF2FF", "#E0E7FF", "#C7D2FE", "#A5B4FC", "#818CF8", "#6366F1", "#4F46E5", "#4338CA", "#3730A3", "#312E81", "#1e1b4b"},
	// Violet
	{"#F5F3FF", "#EDE9FE", "#DDD6FE", "#C4B5FD", "#A78BFA", "#8B5CF6", "#7C3AED", "#6D28D9", "#5B21B6", "#4C1D95", "#2e1065"},
	// Purple
	{"#FAF5FF", "#F3E8FF", "#E9D5FF", "#D8B4FE", "#C084FC", "#A855F7", "#9333EA", "#7E22CE", "#6B21A8", "#581C87", "#3b0764"},
	// Fuchsia
	{"#FDF4FF", "#FAE8FF", "#F5D0FE", "#F0ABFC", "#E879F9", "#D946EF", "#C026D3", "#A21CAF", "#86198F", "#701A75", "#4a044e"},
	// Pink
	{"#FDF2F8", "#FCE7F3", "#FBCFE8", "#F9A8D4", "#F472B6", "#EC4899", "#DB2777", "#BE185D", "#9D174D", "#831843", "#500724"},
	// Rose
	{"#FFF1F2", "#FFE4E6", "#FECDD3", "#FDA4AF", "#FB7185", "#F43F5E", "#E11D48", "#BE123C", "#9F1239", "#881337", "#4c0519"},
}

// type Palette struct {
// 	BlueGray
// }

const (
	BlueGray  = 0
	CoolGray  = 1
	Gray      = 2
	TrueGray  = 3
	WarmGray  = 4
	Red       = 5
	Orange    = 6
	Amber     = 7
	Yellow    = 8
	Lime      = 9
	Green     = 10
	Emerald   = 11
	Teal      = 12
	Cyan      = 13
	LightBlue = 14
	Blue      = 15
	Indigo    = 16
	Violet    = 17
	Purple    = 18
	Fuchsia   = 19
	Pink      = 20
	Rose      = 21
)

const (
	P50  = 0
	P100 = 1
	P200 = 2
	P300 = 3
	P400 = 4
	P500 = 5
	P600 = 6
	P700 = 7
	P800 = 8
	P900 = 9
	P950 = 10
)

var Colors = []int{
	BlueGray,
	CoolGray,
	Gray,
	TrueGray,
	WarmGray,
	Red,
	Orange,
	Amber,
	Yellow,
	Lime,
	Green,
	Emerald,
	Teal,
	Cyan,
	LightBlue,
	Blue,
	Indigo,
	Violet,
	Purple,
	Fuchsia,
	Pink,
	Rose,
}
var Scales = []int{P50, P100, P200, P300, P400, P500, P600, P700, P800, P900, P950}

var cache = map[int]map[int]color.NRGBA{}

func Color(name, scale int) color.NRGBA {

	// TODO
	if _, ok := cache[name]; ok {
		if _, oks := cache[name][scale]; !oks {
			cache[name][scale] = makeColor(name, scale)
		}
	} else {
		cache[name] = map[int]color.NRGBA{
			scale: makeColor(name, scale),
		}
	}

	return cache[name][scale]
}

func makeColor(name, scale int) color.NRGBA {

	colorStr := palette[name][scale]
	colorStr = strings.TrimLeft(colorStr, "#")
	v, _ := strconv.ParseUint(colorStr, 16, 64)
	return Torgb(v)
}
