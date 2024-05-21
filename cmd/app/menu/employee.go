package menu

import (
	"github.com/Facalder/Planify/internal/controller"
	"github.com/Facalder/Planify/models"
	"github.com/Facalder/Planify/pkg"
	"github.com/Facalder/Planify/pkg/component"
	"github.com/pterm/pterm"
)

func Employee() {
	component.Header("Welcome to Employee, Create Account or Log in to Get Into Employee Dashboard")

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
			if controller.LoginEmployee(models.Employees, models.NEmployee) {
				component.Spinner(5, "Validating Your Data, Please Wait a Second...", func() {
					MenuEmployee()
				})
			} else {
				component.Spinner(5, "Validating Your Data, Please Wait a Second...", func() {
					pterm.Error.Println("Username and password doesn't match, please try again!")
				})
			}
		case "2":
			controller.RegisterEmployee(&models.Employees, &models.NEmployee, func() {
				component.Spinner(5, "Create and Save Your Data...", func() {
					MenuEmployee()
				})
			})
		case "3":
			component.Spinner(pkg.Loading, "Leaving Employee...", func() {
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

func MenuEmployee() {
	component.Header("Welcome to Employee Dashboard, Choose Command to use Functionality")

	for {
		pterm.DefaultSection.Println("Choose Manage Function")
		pterm.DefaultBasicText.Println("(1) - Manage Task")

		pterm.DefaultSection.Println("Others Available Commands?")
		pterm.DefaultBasicText.Println("(2) - Back")
		pterm.DefaultBasicText.Println("(3) - Exit Program")

		choose, _ := pterm.DefaultInteractiveTextInput.Show("Choose Command")
		pkg.ChooseMenu = choose

		switch pkg.ChooseMenu {
		case "1":
			component.Spinner(pkg.Loading, "Loading...", func() {
				ManageTaskEmployee()
			})

		case "2":
			component.Spinner(pkg.Loading, "Leaving Employee Menu....", func() {
				Employee()
			})

		case "3":
			component.Spinner(pkg.ExitProgram, "Exiting Program, Please Wait a Second...", func() {
				Exit()
			})

		default:
			pterm.Error.Println("Menu Option Is Not Valid!, Please Fill Command Correctly!")
		}
	}
}
