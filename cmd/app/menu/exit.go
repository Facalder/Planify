package menu

import (
	"github.com/Facalder/Planify/pkg"
	"github.com/Facalder/Planify/pkg/component"
	"github.com/pterm/pterm"
	"github.com/savioxavier/termlink"
)

func Exit() {
	component.HeaderBigText("Thanks!")
	component.Header("Have a nice day, Be Careful!")

	for {
		pterm.DefaultSection.Println("Others Available Commands?")
		pterm.DefaultBasicText.Println("(1) - Reload Project")

		pterm.DefaultSection.Println("Find and Follow Us!")
		pterm.DefaultBasicText.Println(termlink.Link("Caldera's Github", "https://github.com/Facalder"))
		pterm.DefaultBasicText.Println(termlink.Link("Caldera's Linkedin", "https://www.linkedin.com/in/facalder/"))
		pterm.DefaultBasicText.Println(termlink.Link("Caldera's Instagram", "https://www.instagram.com/facalder/"))

		pterm.Println(" ")
		pterm.DefaultBasicText.Println(termlink.Link("Avriela's Instagram", " "))

		choose, _ := pterm.DefaultInteractiveTextInput.Show("Choose Command")
		pkg.ChooseMenu = choose

		switch pkg.ChooseMenu {
		case "1":
			component.Spinner(pkg.ReloadProject, "Reloading Project, Please Wait a Second...", func() {
				IntroScreen()
			})
		default:
			pterm.DefaultLogger.Error("Menu Option Is Not Valid!, Please Fill Command Correctly!")
		}
	}
}
