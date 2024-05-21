package controller

import (
	"fmt"
	"github.com/Facalder/Planify/models"
	"github.com/Facalder/Planify/pkg"
	"github.com/pterm/pterm"
	"time"
)

func LoginAdmin(admins models.TabAdmin, n int) bool {
	var isValidated bool = false

	userName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your UserName")
	password, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Please input your Password")

	for i := 0; i < n; i++ {
		if admins[i].UserName == userName && admins[i].Password == password {
			isValidated = true
			return isValidated
		}
	}

	return isValidated
}

func RegisterAdmin(admins *models.TabAdmin, n *int, menu func()) {
	if *n < pkg.NMAX {
		fullName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your FullName")
		shortName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your ShortName")
		userName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your UserName")
		password, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Please input your Password")

		admin := models.AdminModel{
			Id:        *n,
			FullName:  fullName,
			ShortName: shortName,
			UserName:  userName,
			Password:  password,
			CreatedAt: time.Now(),
		}

		admins[*n] = admin
		*n += 1

		menu()
	} else {
		pterm.Error.Println("You Have Reached The Maksimum Data Allowed!")
	}
}

func GetAllAdmin(admins models.TabAdmin, n int) {
	fmt.Println("Data Admin: ")
	for i := 0; i < n; i++ {
		fmt.Println(admins[i].Id)
		fmt.Println(admins[i].UserName)
		fmt.Println(admins[i].Password)
	}
}

func UpdateAdmin() {}

func DeleteAdmin() {}
