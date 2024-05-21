package menu

import (
	"github.com/Facalder/Planify/internal/controller"
	"github.com/Facalder/Planify/models"
	"github.com/Facalder/Planify/pkg"
	"github.com/Facalder/Planify/pkg/component"
	"github.com/pterm/pterm"
)

func Admin() {
	component.Header("Welcome to Admin, Create account or Log in to get into Admin Dashboard")

	for {
		pterm.DefaultSection.Println("Have an account?")
		pterm.DefaultBasicText.Println("(1) - Login")

		pterm.DefaultSection.Println("Don't Have Any Account?")
		pterm.DefaultBasicText.Println("(2) - Create Account")

		pterm.DefaultSection.Println("Others Available Commands?")
		pterm.DefaultBasicText.Println("(3) - Back")
		pterm.DefaultBasicText.Println("(4) - Exit Program")

		choose, _ := pterm.DefaultInteractiveTextInput.Show("Choose Command")
		pkg.ChooseMenu = choose

		switch pkg.ChooseMenu {
		case "1":
			if controller.LoginAdmin(models.Admins, models.NAdmin) {
				component.Spinner(pkg.ValidateData, "Validating Your Data, Please Wait a Second...", func() {
					MenuAdmin()
				})
			} else {
				component.Spinner(pkg.ValidateData, "Validating Your Data, Please Wait a Second...", func() {
					pterm.Error.Println("Username and password doesn't match, please try again!")
				})
			}
		case "2":
			controller.RegisterAdmin(&models.Admins, &models.NAdmin, func() {
				component.Spinner(pkg.SaveData, "Create and Save Your Data...", func() {
					MenuAdmin()
				})
			})
		case "3":
			component.Spinner(pkg.Loading, "Leaving Admin...", func() {
				InitialMenu()
			})
		case "4":
			component.Spinner(pkg.ExitProgram, "Exiting Program, Please Wait a Second...", func() {
				Exit()
			})
		default:
			pterm.Error.Println("Menu Option Is Not Valid!, Please Fill Command Correctly!")
		}
	}
}

func MenuAdmin() {
	component.Header("Welcome to Admin Dashboard, Choose Command to use Functionality")

	for {
		pterm.DefaultSection.Println("Choose Manage Function")
		pterm.DefaultBasicText.Println("(1) - Manage Manager")
		pterm.DefaultBasicText.Println("(2) - Manage Employee")
		pterm.DefaultBasicText.Println("(3) - Manage Task")

		pterm.DefaultSection.Println("Others Available Commands?")
		pterm.DefaultBasicText.Println("(4) - Back")
		pterm.DefaultBasicText.Println("(5) - Exit Program")

		choose, _ := pterm.DefaultInteractiveTextInput.Show("Choose Command")
		pkg.ChooseMenu = choose

		switch pkg.ChooseMenu {
		case "1":
			component.Spinner(pkg.Loading, "Loading...", func() {
				ManageManagerAdmin()
			})

		case "2":
			component.Spinner(pkg.Loading, "Loading...", func() {
				ManageEmployeeAdmin()
			})

		case "3":
			component.Spinner(pkg.Loading, "Loading...", func() {

			})

		case "4":
			component.Spinner(pkg.Loading, "Leaving Admin Menu....", func() {
				Admin()
			})

		case "5":
			component.Spinner(pkg.ExitProgram, "Exiting Program, Please Wait a Second...", func() {
				Exit()
			})

		default:
			pterm.Error.Println("Menu Option Is Not Valid!, Please Fill Command Correctly!")
		}
	}
}
