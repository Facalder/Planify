package menu

import (
	"github.com/Facalder/Planify/internal/controller"
	"github.com/Facalder/Planify/models"
	"github.com/Facalder/Planify/pkg"
	"github.com/Facalder/Planify/pkg/component"
	"github.com/pterm/pterm"
)

func Manager() {
	component.Header("Welcome to Manager, Create account or Log in to get into Manager Dashboard")

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
			if controller.LoginManager(models.Managers, models.NManager) {
				component.Spinner(pkg.ValidateData, "Validating Your Data, Please Wait a Second...", func() {
					MenuManager()
				})
			} else {
				component.Spinner(pkg.ValidateData, "Validating Your Data, Please Wait a Second...", func() {
					pterm.Error.Println("Username and password doesn't match, please try again!")
				})
			}
			MenuManager()
		case "2":
			controller.RegisterManager(&models.Managers, &models.NManager, func() {
				component.Spinner(pkg.SaveData, "Create and Save Your Data...", func() {
					MenuManager()
				})
			})
		case "3":
			component.Spinner(pkg.Loading, "Leaving Manager...", func() {
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

func MenuManager() {
	component.Header("Welcome to Manager Dashboard, Choose Command to use Functionality")

	for {
		pterm.DefaultSection.Println("Choose Manage Function")
		pterm.DefaultBasicText.Println("(1) - Manage Employee")
		pterm.DefaultBasicText.Println("(2) - Manage Task")

		pterm.DefaultSection.Println("Others Available Commands?")
		pterm.DefaultBasicText.Println("(3) - Back")
		pterm.DefaultBasicText.Println("(4) - Exit Program")

		choose, _ := pterm.DefaultInteractiveTextInput.Show("Choose Command")
		pkg.ChooseMenu = choose

		switch pkg.ChooseMenu {
		case "1":
			component.Spinner(pkg.Loading, "Loading...", func() {
				ManageEmployeeManager()
			})

		case "2":
			component.Spinner(pkg.Loading, "Loading...", func() {
				ManageTaskManager()
			})

		case "3":
			component.Spinner(pkg.Loading, "Leaving Manager Menu....", func() {
				Manager()
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
