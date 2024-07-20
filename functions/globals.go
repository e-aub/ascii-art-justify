package functions

import (
	"fmt"
	"regexp"
	"strings"
)

// flags struct
type ColorFlag struct {
	Color            Color
	ToColor          string
	ToColorIndices   [][]int
	RtfHeader        string
	RtfColor         string
	RtfColorCtrlWord string
	RtfReset         string
	RtfResetCtrlWord string
	NewLineCtrlWord  string
	AnsiColor        string
	AnsiReset        string
}

type Color struct {
	R int
	G int
	B int
}

type arguments struct {
	ToDraw    string
	Banner    string
	FileName  string
	Align     string
	ColorFlag ColorFlag
}

var Args = arguments{
	ToDraw:   "",
	Banner:   "",
	FileName: "",
	Align:    "",
	ColorFlag: ColorFlag{
		Color:            Color{},
		ToColor:          "",
		ToColorIndices:   [][]int{},
		RtfHeader:        "{\\rtf1\\ansi\\deff0{\\fonttbl{\\f0\\fmodern Courier New;}}\\margl0\\margr0{\\colortbl ;\\red%d\\green%d\\blue%d;}\\f0 ",
		RtfColor:         "cf1",
		RtfColorCtrlWord: "\\cf1 ",
		RtfReset:         "cf0",
		RtfResetCtrlWord: "\\cf0 ",
		NewLineCtrlWord:  "\\line",
		AnsiColor:        "\033[38;2;%d;%d;%dm",
		AnsiReset:        "\033[0m",
	},
}

var (
	JustifyPattern = regexp.MustCompile(`^-{1,2}align`)
	JustifyCheck   = regexp.MustCompile(`^(?:--align=)(center|left|right|justify)$`)
	OutputPattern  = regexp.MustCompile(`^-{1,2}output`)
	OutputCheck    = regexp.MustCompile(`^(?:--output=)(.+.txt|.+.rtf)$`)
	ColorPattern   = regexp.MustCompile(`^-{1,2}color`)
	ColorCheck     = regexp.MustCompile(`^(?:--color=)(.+)$`)
	ValidBanner    = regexp.MustCompile(`^standard$|^shadow$|^enigma$|^nirvana$|^standard.txt$|^shadow.txt$|^enigma.txt$|^nirvana.txt$`)
	RgbPattern     = regexp.MustCompile(`^rgb\(`)
	RgbCheck       = regexp.MustCompile(`^rgb\((\d{1,3}), (\d{1,3}), (\d{1,3})\)$`)
	HexPattern     = regexp.MustCompile(`^#`)
	HexCheck       = regexp.MustCompile(`^#[0-9A-Fa-f]{6}$`)
)

var Errors map[string]string = map[string]string{
	"output":       "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . EX: go run . --output=<filename.txt> \"something\"",
	"color":        "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . EX: go run . --color=<color> <letters to be colored> \"something\"",
	"fs":           "Invalid banner\n\nAvailable banners: standard, shadow, thinkertoy, enigma, nirvana",
	"invalidColor": AvailableColors(),
	"internal":     "Internal program error",
	"rgbValue":     "Invalid RGB value\n\nThe numbers must be between 0 and 255 (ex: rgb(<0-255>, <0-255>, <0-255>))",
	"rgbFormat":    "Invalid RGB format\n\nThe format must same as: rgb(<0-255>, <0-255>, <0-255>)",
	"invalidhex":   "Invalid Hex Color\n\nThe hex code must be between #000000 -> #FFFFFF",
	"char":         "Invalid Character\n\nThe characters must be between ' ' and '~'",
}

func AvailableColors() string {
	var str strings.Builder
	str.WriteString("Available colors:\n\n")
	for name, color := range Colors {
		str.WriteString(fmt.Sprintf("\033[38;2;%d;%d;%dm%s%s"+", ", color.R, color.G, color.B, name, "\033[0m, "))
	}
	return str.String()
}

var Colors map[string]Color = map[string]Color{
	"black":                {0, 0, 0},
	"red":                  {255, 0, 0},
	"green":                {0, 255, 0},
	"blue":                 {0, 0, 255},
	"yellow":               {255, 255, 0},
	"orange":               {255, 165, 0},
	"magenta":              {255, 0, 255},
	"cyan":                 {0, 255, 255},
	"purple":               {128, 0, 128},
	"quartz":               {81, 65, 79},
	"lightblue":            {0, 255, 255},
	"lightgreen":           {0, 255, 0},
	"lightred":             {255, 0, 0},
	"lightyellow":          {255, 255, 0},
	"redpurple":            {149, 53, 83},
	"darkgray":             {50, 50, 50},
	"lightgray":            {200, 200, 200},
	"maroon":               {128, 0, 0},
	"olivegreen":           {85, 107, 47},
	"navyblue":             {0, 0, 128},
	"teal":                 {0, 128, 128},
	"limegreen":            {0, 255, 127},
	"goldenrod":            {218, 165, 32},
	"indigo":               {75, 0, 130},
	"darkorange":           {204, 85, 37},
	"forestgreen":          {34, 139, 34},
	"royalblue":            {65, 105, 225},
	"fuchsia":              {255, 0, 255},
	"turquoise":            {49, 241, 242},
	"springgreen":          {0, 250, 154},
	"darkgoldenrod":        {184, 134, 11},
	"violet":               {135, 60, 240},
	"peru":                 {205, 133, 63},
	"indianred":            {205, 92, 92},
	"darkseagreen":         {143, 188, 143},
	"steelblue":            {70, 130, 180},
	"deeppink":             {255, 0, 102},
	"mediumseagreen":       {60, 179, 113},
	"cornflowerblue":       {100, 149, 237},
	"chartreuse":           {127, 255, 0},
	"darkorchid":           {153, 50, 204},
	"palegreen":            {152, 255, 152},
	"slateblue":            {106, 90, 205},
	"mediumpurple":         {147, 112, 219},
	"coral":                {255, 128, 80},
	"saddlebrown":          {139, 69, 19},
	"seagreen":             {46, 139, 87},
	"mediumblue":           {0, 0, 205},
	"lightpink":            {255, 182, 193},
	"lightcoral":           {240, 128, 128},
	"lightsalmon":          {255, 160, 122},
	"darksalmon":           {233, 150, 122},
	"lightgoldenrodyellow": {255, 255, 141},
	"palegoldenrod":        {238, 232, 170},
	"khaki":                {240, 230, 140},
	"beige":                {245, 245, 220},
	"lavender":             {230, 230, 250},
	"mintcream":            {245, 255, 250},
	"ivory":                {255, 255, 240},
	"wheat":                {245, 222, 179},
	"moccasin":             {255, 228, 181},
	"oldlace":              {253, 245, 230},
	"papayawhip":           {255, 239, 213},
	"blanchedalmond":       {255, 235, 205},
	"peachpuff":            {255, 218, 185},
	"lemonchiffon":         {255, 250, 220},
	"pink":                 {255, 192, 203},
	"hotpink":              {255, 105, 180},
	"mediumvioletred":      {199, 21, 133},
	"orchid":               {214, 133, 181},
	"plum":                 {221, 160, 221},
	"mediumorchid":         {186, 85, 211},
	"darkviolet":           {148, 0, 211},
	"blueviolet":           {138, 43, 226},
	"darkslateblue":        {79, 75, 255},
	"mediumslateblue":      {123, 104, 238},
	"dodgerblue":           {30, 144, 255},
	"mediumturquoise":      {72, 209, 204},
	"lightseagreen":        {32, 178, 170},
	"sienna":               {160, 82, 45},
	"royalpurple":          {29, 0, 110},
}
