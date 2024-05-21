package menu

import (
	"github.com/Facalder/Planify/pkg"
	"github.com/Facalder/Planify/pkg/component"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func ManageTaskEmployee() {
	component.Header(putils.CenterText("Welcome to Manage Task (Employee) Menu!, \n" +
		"Manage Your Employees to Work Harder, and Harder!"))

	for {
		pterm.DefaultSection.Println("Available Commands")
		pterm.DefaultBasicText.Println("(1) - Add New Data Task")
		pterm.DefaultBasicText.Println("(2) - Edit Data Task")
		pterm.DefaultBasicText.Println("(3) - Show All Task")
		pterm.DefaultBasicText.Println("(4) - Remove Data Task")
		pterm.DefaultBasicText.Println("(5) - Manage Task Cost")

		pterm.DefaultSection.Println("Others Available Commands?")
		pterm.DefaultBasicText.Println("(6) - Back")
		pterm.DefaultBasicText.Println("(7) - Exit Program")

		choose, _ := pterm.DefaultInteractiveTextInput.Show("Choose Command")
		pkg.ChooseMenu = choose

		switch pkg.ChooseMenu {
		case "1":
			//controller.RegisterAdmin(&Admins, &NAdmin, func() {
			//	component.Spinner(pkg.SaveData, "Create and Save Your Data...", func() {
			//		MenuAdmin()
			//	})
			//})
		case "2":
			//controller.RegisterAdmin(&Admins, &NAdmin, func() {
			//	component.Spinner(pkg.SaveData, "Create and Save Your Data...", func() {
			//		MenuAdmin()
			//	})
			//})
		case "3":
		//controller.RegisterAdmin(&Admins, &NAdmin, func() {
		//	component.Spinner(pkg.SaveData, "Create and Save Your Data...", func() {
		//		MenuAdmin()
		//	})
		//})

		case "4":
		//controller.RegisterAdmin(&Admins, &NAdmin, func() {
		//	component.Spinner(pkg.SaveData, "Create and Save Your Data...", func() {
		//		MenuAdmin()
		//	})
		
		//})
		case "5":
			component.Spinner(pkg.Loading, "Loading...", func() {
				ManageTaskCost()
			})
		case "6":
			component.Spinner(pkg.Loading, "Leaving Manage Task Menu....", func() {
				MenuEmployee()
			})
		case "7":
			component.Spinner(pkg.ExitProgram, "Exiting Program, Please Wait a Second...", func() {
				Exit()
			})
		default:
			pterm.Error.Println("Menu Option Is Not Valid!, Please Fill Command Correctly!")
		}
	}
}
