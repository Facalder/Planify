package repositories

import (
	"github.com/Facalder/Planify/models"
	"github.com/Facalder/Planify/pkg"
)

var admins [pkg.NMAX]models.AdminModel

func GetAllAdmin() [pkg.NMAX]models.AdminModel {
	return admins
}

func CreateAdmin(admin models.AdminModel) models.AdminModel {
	var i = 0

	if i < pkg.NMAX {
		admins[i] = admin
		i++
	}

	return admin
}

func UpdateByIdAdmin() {}

func DeleteByIdAdmin() {}
