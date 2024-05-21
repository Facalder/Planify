package controller

import (
	"github.com/Facalder/Planify/models"
	"github.com/Facalder/Planify/pkg"
	"github.com/pterm/pterm"
	"time"
)

func LoginManager(managers models.TabManager, n int) bool {
	var isValidated bool = false

	userName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your UserName")
	password, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Password")

	for i := 0; i < n; i++ {
		if managers[i].UserName == userName && managers[i].Password == password {
			isValidated = true
			return isValidated
		}
	}

	return isValidated
}

func RegisterManager(managers *models.TabManager, n *int, menu func()) {
	if *n < pkg.NMAX {
		fullName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your FullName")
		shortName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your ShortName")
		userName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your UserName")
		password, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Please input your Password")
		bio, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Bio")
		department, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Department")

		manager := models.ManagerModel{
			Id:         *n,
			FullName:   fullName,
			ShortName:  shortName,
			UserName:   userName,
			Password:   password,
			Bio:        bio,
			Department: department,
			CreatedAt:  time.Now(),
		}

		managers[*n] = manager
		*n += 1

		menu()
	} else {
		pterm.Error.Println("You Have Reached The Maksimum Data Allowed!")
	}
}

func CreateManager(managers *models.TabManager, n *int) {
	if *n < pkg.NMAX {
		fullName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your FullName")
		shortName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your ShortName")
		userName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your UserName")
		password, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Please input your Password")
		bio, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Bio")
		department, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Department")

		manager := models.ManagerModel{
			Id:         *n,
			FullName:   fullName,
			ShortName:  shortName,
			UserName:   userName,
			Password:   password,
			Bio:        bio,
			Department: department,
			CreatedAt:  time.Now(),
		}

		managers[*n] = manager
		*n += 1
	} else {
		pterm.Error.Println("You Have Reached The Maksimum Data Allowed!")
	}
}

func UpdateManager() {}

func ShowAllManager() {}

func DeleteManager() {}
