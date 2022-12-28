package testing

import (
	"database/sql"
	"fmt"

	"api-inventory-go/src/controllers/user/controller"
	"api-inventory-go/src/modules/user/model"
	"api-inventory-go/src/modules/user/repository"
)

func CreateUserDbTest(db *sql.DB) error {
	fmt.Println("===============================================================")
	fmt.Println("Create User Saved to Database Test")

	newUserObject := model.NewUser()
	newUserObject.Fullname = "Sudirman Said"
	newUserObject.Nickname = "Sudirman"
	newUserObject.Username = "sudirmans"
	newUserObject.NoHP = "081234567890"
	newUserObject.Email = "sudirmans@gmail.com"
	newUserObject.Password = "secret"
	newUserObject.Image = "https://img.okezone.com/content/2017/08/11/33/1753946/pendam-rindu-untuk-anak-jono-amstrong-masih-takut-bertemu-keluarga-mantan-istri-JfpZnS320L.jpg"
	newUserObject.Office = "Kantor 2"
	newUserObject.Role = "Administrator"

	userRepositoryPostgres := repository.NewUserRepositoryPostgres(db)

	err := controller.CreateUser(newUserObject, userRepositoryPostgres)
	if err != nil {
		return err
	}

	fmt.Println("===============================================================")

	return nil
}

func UpdateUserDbTest(db *sql.DB) error {
	fmt.Println("===============================================================")
	fmt.Println("Update User Saved to Database Test")

	newUserObject := model.UpdatedUser()
	newUserObject.ID = "P1"
	newUserObject.Fullname = "Marka Jalan Newest Banget"
	newUserObject.Nickname = "Marka"
	newUserObject.Username = "markaj"
	newUserObject.NoHP = "081234567890"
	newUserObject.Email = "markaj@gmail.com"
	newUserObject.Password = "secret"
	newUserObject.Image = "https://img.okezone.com/content/2017/08/11/33/1753946/pendam-rindu-untuk-anak-jono-amstrong-masih-takut-bertemu-keluarga-mantan-istri-JfpZnS320L.jpg"
	newUserObject.Office = "Kantor 2"
	newUserObject.Role = "Administrator"

	userRepositoryPostgres := repository.NewUserRepositoryPostgres(db)

	err := controller.UpdateUser(newUserObject, userRepositoryPostgres)
	if err != nil {
		return err
	}

	fmt.Println("===============================================================")

	return nil
}

func DeleteUserDbTest(db *sql.DB) error {
	fmt.Println("===================================================================")
	fmt.Println("Delete User in Database Test")

	id := "U3"
	userRepositoryPostgres := repository.NewUserRepositoryPostgres(db)

	err := controller.DeleteUser(id, userRepositoryPostgres)
	if err != nil {
		return err
	}

	fmt.Println("===================================================================")
	return nil
}

func GetUserDbTest(db *sql.DB) error {
	fmt.Println("===================================================================")
	fmt.Println("Get User from Database Test")

	id := "U1"
	userRepositoryPostgres := repository.NewUserRepositoryPostgres(db)

	user, err := controller.GetUser(id, userRepositoryPostgres)
	if err != nil {
		return err
	}

	fmt.Println("Result:", user)
	fmt.Println("===================================================================")
	return nil
}

func GetUsersDbTest(db *sql.DB) error {
	fmt.Println("===================================================================")
	fmt.Println("Get Users from Database Test")

	userRepositoryPostgres := repository.NewUserRepositoryPostgres(db)

	users, err := controller.GetUsers(userRepositoryPostgres)
	if err != nil {
		return err
	}

	fmt.Println("Results:")
	for _, v := range users {
		fmt.Println(v)
	}
	fmt.Println("===================================================================")
	return nil
}
