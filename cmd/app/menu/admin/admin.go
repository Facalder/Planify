package admin

import (
	"fmt"
	"github.com/Facalder/Planify/cmd/app/menu"
	"github.com/Facalder/Planify/internal/controller"
	"github.com/Facalder/Planify/pkg"
	"github.com/pterm/pterm"
)

func Admin() {
	fmt.Println("----Welcome in Admin Page, Create account or Log in to get into Admin Dashboard----")

	for {
		fmt.Println("--------------------------------------------")
		fmt.Println("Have an account?")
		fmt.Println("1. Log in")

		fmt.Println("")

		fmt.Println("Do you have an account?")
		fmt.Println("2. Create Account")

		fmt.Println("")

		fmt.Println("3. Back")
		fmt.Println("4. Exit")
		fmt.Println("----------------- END ----------------------")
		fmt.Println("--------------------------------------------")
		fmt.Scan(&pkg.ChooseMenu)

		switch pkg.ChooseMenu {
		case 1:
			if controller.LoginAdmin() {
				MenuAdmin()
			} else {
				pterm.DefaultLogger.Error("Username and password doesn't match, please try again!")
			}
		case 2:
			controller.RegisterAdmin(func() {
				MenuAdmin()
			})
		case 3:
			menu.Exit()
		default:
			fmt.Println("Menu option is not valid. Please try again")
		}
	}
}

func MenuAdmin() {
	fmt.Println("Hello Worlds")
}
