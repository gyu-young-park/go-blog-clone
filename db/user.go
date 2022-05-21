package db

import (
	"log"
)

type User struct {
	id       int64
	Name     string
	Email    string
	Password string
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
	err := store.db.QueryRow("SELECT user_name, email, password FROM User WHERE id = ?", id).Scan(&user.Name, &user.Email, &user.Password)
	if err != nil {
		log.Printf("failed to get USER[%v]\n", err.Error())
		return user, err
	}
	return user, nil
}

func (store *Store) GetUserIdByEmail(email string) (int64, error) {
	var id int64
	err := store.db.QueryRow("SELECT id FROM User WHERE email = ?", email).Scan(&id)
	if err != nil {
		log.Printf("failed to get USER id[%v]\n", err.Error())
		return -1, nil
	}
	return id, nil
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

func (store *Store) UpdateUserInfo(preEmail string, email string, name string) (User, error) {
	row, err := store.db.Exec("UPDATE User SET email = ?, user_name = ? WHERE email = ?", email, name, preEmail)
	if err != nil {
		log.Printf("failed to update user[%v]\n", err.Error())
		return User{}, nil
	}
	if rowCount, err := row.RowsAffected(); rowCount == 0 || err != nil {
		log.Printf("failed to update user[%v]\n", err.Error())
		return User{}, nil
	}
	id, err := store.GetUserIdByEmail(email)
	if err != nil {
		log.Printf("failed to get new user[%v]\n", err.Error())
		return User{}, nil
	}
	return store.GetUserById(id)
}

func (store *Store) DeleteUser(id int64) error {
	_, err := store.db.Exec("DELETE FROM User WHERE id = ?", id)
	if err != nil {
		log.Printf("failed to delete user[%v]\n", err.Error())
		return err
	}
	return nil
}
