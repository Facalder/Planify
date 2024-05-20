package menu

import (
	"github.com/Facalder/Planify/pkg"
	"github.com/Facalder/Planify/pkg/component"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func InitialMenu() {
	component.Header(putils.CenterText("The beginning of the story of this program begins here!," +
		"\nChoose Your Role Before Use Application"))

	for {
		pterm.DefaultSection.Println("Pick Your Role?")
		pterm.DefaultBasicText.Println("(1) - Admin")
		pterm.DefaultBasicText.Println("(2) - Manager")
		pterm.DefaultBasicText.Println("(3) - Employee")

		pterm.DefaultSection.Println("Others Available Commands?")
		pterm.DefaultBasicText.Println("(4) - Exit Program")

		choose, _ := pterm.DefaultInteractiveTextInput.Show("Choose Command")
		pkg.ChooseMenu = choose

		switch pkg.ChooseMenu {
		case "1":
			component.Spinner(1, "Entering Admin...", func() {
				Admin()
			})
		case "2":
			component.Spinner(1, "Entering Manager...", func() {
				Manager()
			})

		case "3":
			component.Spinner(1, "Entering Employee...", func() {
				Employee()
			})

		case "4":
			component.Spinner(5, "Exiting Program, Please Wait a Second...", func() {
				Exit()
			})

		default:
			pterm.Error.Println("Menu Option Is Not Valid!, Please Fill Command Correctly!")
		}
	}
}
