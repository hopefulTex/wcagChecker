package color

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/ansi"
)

type Score int

var (
	tagWidth   int            = 8
	blockStyle lipgloss.Style = lipgloss.NewStyle().
			Width(tagWidth).Height(tagWidth / 2).
		// Border(lipgloss.HiddenBorder()).
		SetString(" ")
	borderStyle lipgloss.Style = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder())
	centerText lipgloss.Style = lipgloss.NewStyle().
			Width(tagWidth).Height(1).
		// Border(lipgloss.HiddenBorder()).
		Align(lipgloss.Center)
)

func (s Score) ToString() string {
	if s == AAA {
		return "AAA"
	}
	if s == AA {
		return "AA"
	}
	if s == FAILED {
		return "FAILED"
	}
	return "error converting Score to string"
}

const (
	FAILED Score = iota
	AA
	AAA
)

type Color struct {
	red   int
	green int
	blue  int
	hex   string
}

func FromHex(hex string) (Color, error) {
	var err error
	var tmp int64
	color := Color{}
	if strings.HasPrefix(hex, "#") {
		color.hex = hex
		hex = hex[1:]
	} else {
		color.hex = "#" + hex
	}
	if len(hex) > 6 {
		return color, fmt.Errorf("hexadecimal colors are limited to 6 places")
	}

	for len(hex) < 6 {
		hex += "0"
	}

	r := hex[:2]
	g := hex[2:4]
	b := hex[4:]

	tmp, err = strconv.ParseInt(r, 16, 32)
	if err != nil {
		return color, nil
	}
	color.red = int(tmp)

	tmp, err = strconv.ParseInt(g, 16, 32)
	if err != nil {
		return color, nil
	}
	color.green = int(tmp)

	tmp, err = strconv.ParseInt(b, 16, 32)
	if err != nil {
		return color, nil
	}
	color.blue = int(tmp)

	return color, nil
}

func Contrast(first, last Color) float64 {
	var contrast float64
	var tmp float64
	var L1 float64 = first.Luminance()
	var L2 float64 = last.Luminance()

	if L1 < L2 {
		tmp = L1
		L1 = L2
		L2 = tmp
	}

	contrast = (L1 + 0.05) / (L2 + 0.05)
	//contrast = roundToPlace(contrast, 2)

	return contrast
}

func (c Color) TagView() string {
	var view strings.Builder
	blk := blockStyle.Background(lipgloss.Color(c.hex))
	view.WriteString(blk.String())
	view.WriteRune('\n')
	view.WriteString(centerText.Render(c.hex))

	return borderStyle.Render(view.String())
}

func ComplianceView(first, last Color) string {
	var view string

	bStyle := borderStyle.Height(2).AlignHorizontal(lipgloss.Center)

	c1 := first.TagView()
	c2 := last.TagView()

	view = lipgloss.JoinHorizontal(lipgloss.Center, c1, c2)

	score, contrast := Compliance(first, last)
	contrastStr := fmt.Sprintf("%f", contrast)
	index := strings.IndexRune(contrastStr, '.')
	contrastStr = contrastStr[0 : index+2]

	newLine := strings.Index(view, "\n")
	if newLine == -1 {
		newLine = len(view)
	}

	text := fmt.Sprintf("%s\n%s", score.ToString(), contrastStr)
	bStyle = bStyle.Width(ansi.PrintableRuneWidth(view[:newLine-1]) - 3).Align(lipgloss.Center)

	return bStyle.Render(text) + "\n" + view
}

func Compliance(first, last Color) (Score, float64) {
	contrast := Contrast(first, last)

	if contrast > 7 {
		return AAA, contrast
	} else if contrast >= 4.5 {
		return AA, contrast
	} else {
		return FAILED, contrast
	}

}

func Normalize(color float64) float64 {
	if color <= 0.03928 {
		color /= 12.92
	} else {
		color += 0.055
		color /= 1.055
		color = math.Pow(color, 2.4)
	}
	return color
}

func (c Color) Luminance() float64 {
	var lum float64 = 0.0
	var r float64
	var g float64
	var b float64

	r = float64(c.red) / 255.0
	g = float64(c.green) / 255.0
	b = float64(c.blue) / 255.0

	lum += 0.2126 * r
	lum += 0.7152 * g
	lum += 0.0722 * b

	return lum
}
