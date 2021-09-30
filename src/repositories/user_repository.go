package repositories

import (
	"database/sql"
	"fmt"
	"pg-conn/src/helpers"
	"pg-conn/src/models"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetByEmail(email string) models.User
	GetById(id string) models.User
	All() ([]models.User, error)
	Create(user models.User) (models.User, error)
	Update(user models.User, id string) (models.User, error)
	UpdatePassword(pass string, id string) error
	Delete(id string) error
}

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return UserRepo{
		db: db,
	}
}

func (repo *UserRepo) GetById(id string) models.User {
	var user models.User
	stmt, err := repo.db.Prepare(`SELECT * FROM users WHERE id = $1`)
	if err != nil {
		return models.User{}
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&user.ID, &user.Name, &user.Lastname, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return models.User{}
	}

	return user
}

func (repo *UserRepo) GetByEmail(email string) models.User {
	var user models.User
	stmt, err := repo.db.Prepare(`SELECT * FROM "user" WHERE email = $1`)
	if err != nil {
		return models.User{}
	}
	defer stmt.Close()

	row := stmt.QueryRow(email)
	err = row.Scan(&user.ID, &user.Name, &user.Lastname, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return models.User{}
	}

	return user
}

func (repo *UserRepo) All() ([]models.User, error) {
	var users []models.User
	stmt, err := repo.db.Prepare(`SELECT * FROM "user"`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for result.Next() {
		var user models.User
		if err := result.Scan(&user.ID, &user.Name, &user.Lastname, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := result.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *UserRepo) Create(user models.User) (models.User, error) {
	stmt, err := repo.db.Prepare(`INSERT INTO "user" (name, lastname, password, email) VALUES ($1, $2, $3, $4)`)
	if err != nil {
		return models.User{}, err
	}

	defer stmt.Close()

	user.Password, err = helpers.GenerateHash(user.Password)

	if err != nil {
		return models.User{}, err

	}

	_, err = stmt.Exec(user.Name, user.Lastname, user.Password, user.Email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (repo *UserRepo) Update(user models.User, id string) (models.User, error) {
	if isValid := helpers.ValidateMail(user.Email); !isValid {
		return models.User{}, fmt.Errorf("invalid email")
	}

	userInDb := repo.GetById(id)
	if userInDb.ID <= 0 {
		return models.User{}, fmt.Errorf("does not exists user with this id")
	}

	e := repo.GetByEmail(user.Email)
	if len(e.Email) > 0 && userInDb.ID != e.ID {
		return models.User{}, fmt.Errorf("email already exists")
	}

	userInDb.Email = user.Email
	userInDb.Name = user.Name
	userInDb.Lastname = user.Lastname
	stmt, err := repo.db.Prepare(`UPDATE "user" SET name = $1, lastname = $2, email = $3 WHERE id = $4`)
	if err != nil {
		return models.User{}, err
	}

	_, err = stmt.Exec(userInDb.Name, userInDb.Lastname, userInDb.Email, id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (repo *UserRepo) UpdatePassword(password string, id string) error {
	user := repo.GetById(id)
	stmt, err := repo.db.Prepare(`UPDATE "user" SET password = $1 WHERE id = $2`)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == nil {
		return fmt.Errorf("new password cant be same of old password")
	}

	hashedPass, err := helpers.GenerateHash(password)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(hashedPass, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepo) DeleteUser(id string) error {
	stmt, err := repo.db.Prepare(`DELETE FROM "user" WHERE user_id = $1`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
