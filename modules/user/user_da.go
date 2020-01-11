package user

import (
	"uangq-be/bridge/db"
)

func GetUserByUsername(username string) (User, error) {
	var u User
	err := db.MyDbService.Db.QueryRow("SELECT id, username, password_hash FROM user_tab WHERE username=?", username).Scan(&u.Id, &u.Username, &u.PasswordHash)

	if err != nil {
		return User{}, err
	}
	return u, nil
}

func CreateUser(username, hashedPassword string) (User, error) {
	var u User
	res, err := db.MyDbService.Db.Exec("INSERT INTO user_tab (username, password_hash) VALUES(?, ?)", username, hashedPassword)

	if err != nil {
		return User{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return User{}, err
	}
	u.Id = id

	return u, nil
}
