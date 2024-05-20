package component

import (
	"github.com/Facalder/Planify/pkg/theme"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func HeaderBigText(title string) {
	planifyLogo, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromStringWithRGB(title, pterm.NewRGB(117, 106,
		182))).Srender()

	pterm.Println(planifyLogo)
}

func Description(description string) {
	pterm.NewRGBStyle(theme.TextForegroundRGB, theme.BackgroundRGB).AddOptions(pterm.Bold).Println(
		description)

	pterm.Println("")
}
