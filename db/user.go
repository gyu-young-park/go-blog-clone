package db

import (
	"log"
)

type User struct {
	Name  string
	Email string
}

func (store *Store) GetAllUserData() ([]User, error) {
	var user User
	var users []User

	rows, err := store.db.Query("SELECT user_name, email FROM User")
	if err != nil {
		log.Printf("failed to get USER data[%v]", err.Error())
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Name, &user.Email)
		if err != nil {
			log.Printf("failed to get USER row[%v]\n", err.Error())
		}
		users = append(users, user)
	}
	return users, nil
}

func (store *Store) GetUserById(id int64) (User, error) {
	var user User
	err := store.db.QueryRow("SELECT user_name, email FROM User WHERE id = ?", id).Scan(&user.Name, &user.Email)
	if err != nil {
		log.Printf("failed to get USER[%v]\n", err.Error())
		return user, err
	}
	return user, nil
}

func (store *Store) RegisterUser(name string, email string, password string) (User, error) {
	row, err := store.db.Exec("INSERT INTO User (user_name, email, password) VALUES(? , ?, ?)", name, email, password)
	if err != nil {
		log.Printf("failed to insert user[%v]\n", err.Error())
		return User{}, err
	}
	id, _ := row.LastInsertId()
	return store.GetUserById(id)
}
