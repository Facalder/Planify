package menu

import (
	"github.com/Facalder/Planify/pkg"
	"github.com/Facalder/Planify/pkg/component"
	"github.com/pterm/pterm"
)

func Exit() {
	component.HeaderBigText("Thanks!")
	component.Header("Have a nice day, Be Careful!")

	for {
		pterm.DefaultSection.Println("Others Available Commands?")
		pterm.DefaultBasicText.Println("(1) - Reload Project")

		choose, _ := pterm.DefaultInteractiveTextInput.Show("Choose Command")
		pkg.ChooseMenu = choose

		switch pkg.ChooseMenu {
		case "1":
			component.Spinner(5, "Reloading Project, Please Wait a Second...", func() {
				IntroScreen()
			})
		default:
			pterm.DefaultLogger.Error("Menu Option Is Not Valid!, Please Fill Command Correctly!")
		}
	}
}
