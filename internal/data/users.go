package data

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type UserModel struct {
	db *sql.DB
}

func (model *UserModel) Insert(u *User) error {
	query := `INSERT INTO msvc_users (username, password, email, role) values
	($1, $2, $3, $4) RETURNING id; 
	`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return model.db.QueryRowContext(ctx, query, u.Username, u.Password, u.Email, u.Role).Scan(&u.Id)
}

func (model *UserModel) FindById(id int) (User, error) {
	query := `Select id, username, password, email, role from msvc_users where id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	res := model.db.QueryRowContext(ctx, query, id)

	if res.Err() != nil {
		return user, res.Err()
	}

	res.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Role)

	return user, nil

}
