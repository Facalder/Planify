package menu

import (
	"fmt"
	"github.com/Facalder/Planify/pkg"
	"github.com/Facalder/Planify/pkg/component"
	"github.com/Facalder/Planify/pkg/theme"
	"github.com/pterm/pterm"
)

func Exit() {
	for {
		component.HeaderBigText("thank you so much")

		component.Description("Have a nice day, and be careful!")

		pterm.NewRGBStyle(theme.TextForegroundRGB).AddOptions(pterm.Bold).Println(
			"Available Command: ")
		pterm.Println(pterm.Red("(1) - Reload Project"))

		fmt.Println("")
		component.Description("----------------- END ----------------------")

		fmt.Scan(&pkg.ChooseMenu)

		switch pkg.ChooseMenu {
		case 1:
			IntroScreen()
			pkg.Clear()
		}
	}
}
