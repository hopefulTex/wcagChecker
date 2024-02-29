package color

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/ansi"
)

var (
	tagWidth   int            = 8
	blockStyle lipgloss.Style = lipgloss.NewStyle().
			Width(tagWidth).Height(tagWidth / 2).
			Align(lipgloss.Center).
			PaddingTop(1).PaddingBottom(2).
			PaddingRight(2).
			SetString(" ")
	borderStyle lipgloss.Style = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder())
	centerText lipgloss.Style = lipgloss.NewStyle().
			Width(tagWidth).Height(1).
			Align(lipgloss.Center)
)

func (s Score) String() string {
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

func (c Color) String() string {
	c.hex = fmt.Sprintf("#%02X%02X%02X", c.red, c.green, c.blue)
	return c.hex
}

func (c Color) Lipgloss() lipgloss.Color {
	return lipgloss.Color(c.String())
}

func (c Color) TagView() string {
	var view strings.Builder
	blk := blockStyle.Background(c.Lipgloss())
	view.WriteString(blk.String())
	view.WriteRune('\n')
	view.WriteString(centerText.Render(c.String()))

	return borderStyle.Render(view.String())
}

func (c Color) TextTagView(color Color) string {
	var view strings.Builder

	blk := blockStyle.
		Background(c.Lipgloss()).
		Foreground(color.Lipgloss())

	view.WriteString(blk.Render("Text"))
	view.WriteRune('\n')
	view.WriteString(centerText.Render(c.String()))

	return borderStyle.Render(view.String())
}

func ComplianceView(first, last Color) string {
	var view string

	bStyle := borderStyle.Height(2).AlignHorizontal(lipgloss.Center)

	c1 := first.TextTagView(last)
	c2 := last.TextTagView(first)

	view = lipgloss.JoinHorizontal(lipgloss.Center, c1, c2)

	score, contrast := Compliance(first, last)
	contrastStr := ContrastString(contrast)

	newLine := strings.Index(view, "\n")
	if newLine == -1 {
		newLine = len(view)
	}

	text := fmt.Sprintf("%s\n%s", score.String(), contrastStr)
	bStyle = bStyle.Width(ansi.PrintableRuneWidth(view[:newLine-1]) - 3).Align(lipgloss.Center)

	return bStyle.Render(text) + "\n" + view
}

func ComplianceString(first, last Color) string {
	score, contrast := Compliance(first, last)
	contrastStr := ContrastString(contrast)
	return fmt.Sprintf("%s: %s", score.String(), contrastStr)
}

func ContrastString(contrast float64) string {
	contrastStr := fmt.Sprintf("%f", contrast)
	index := strings.IndexRune(contrastStr, '.')
	contrastStr = contrastStr[0 : index+2]
	return contrastStr
}
