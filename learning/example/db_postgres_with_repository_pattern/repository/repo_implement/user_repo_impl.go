package repo_implement

import (
	"database/sql"
	"fmt"
	models "testProject/learning/example/db_postgres_with_repository_pattern/model"
	repo "testProject/learning/example/db_postgres_with_repository_pattern/repository"
)

type UserRepoImpl struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) repo.UserRepo {
	return &UserRepoImpl{Db: db}
}

func (u *UserRepoImpl) Select() ([]models.User, error) {
	users := make([]models.User, 0)

	rows, err := u.Db.Query("SELECT * FROM users")
	if err != nil {
		return users, err
	}

	// mapping data from DB to struct
	for rows.Next() {
		user := models.User{}

		// select id, name, gender, email from users;
		err := rows.Scan(&user.ID, &user.Name, &user.Gender, &user.Email)
		if err != nil {
			break
		}

		// append row to users
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return users, nil
	}

	return users, nil
}

func (u *UserRepoImpl) Insert(user models.User) error {
	sqlInsertStatement := `
        INSERT INTO users (name, gender, email)
        VALUES ($1, $2, $3);
    `

	_, err := u.Db.Exec(sqlInsertStatement, user.Name, user.Gender, user.Email)
	if err != nil {
		return err
	}

	fmt.Println("Inserted user")
	return nil
}
