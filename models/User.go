package models

import "example.com/mygolangproj/db"

type User struct {
	ID       int
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users(email, password) VALUES(?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}
	return nil
}

func GetAllUsers() ([]User, error) {
	query := `SELECT id, email, password FROM users`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserByID(id int) (User, error) {
	query := `SELECT id, email, password FROM users WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func UpdateUser(user User) error {
	query := `UPDATE users SET email = ?, password = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Email, user.Password, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
