package model

import "errors"

func Login(email string, password string) (map[string]string, error) {
	// db := framework.Database{}
	// defer db.Close()
	// db.Select("id, nama, email, password, role")
	// db.From("vendor")
	// db.Where("email", email)
	// db.Where("role", 7)
	// user, err := db.Row()
	// if err != nil {
	// 	return nil, err
	// }
	// if framework.ValidPassword(password, user["password"].(string)) {
	// 	return user, nil
	// }
	return nil, errors.New("Invalid password")
}
