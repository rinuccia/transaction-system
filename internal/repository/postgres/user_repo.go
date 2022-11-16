package postgres

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/rinuccia/transaction-system/internal/models"
	"math"
)

type userRepo struct {
	*sqlx.DB
}

func newUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db}
}

func (r *userRepo) FindUser(userID string) (models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE id = $1`
	row := r.QueryRow(query, userID)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Balance)
	if err != nil {
		return user, errors.New("account not found")
	}
	user.Balance = user.Balance / 100
	return user, nil
}

func (r *userRepo) InsertUser(user models.User) (string, error) {
	query := `INSERT INTO users (first_name, last_name, email, balance) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.QueryRow(query, user.FirstName, user.LastName, user.Email, 0).Scan(&user.ID)
	if err != nil {
		return "", err
	}
	return user.ID, nil
}

func (r *userRepo) Replenish(request models.UserRequest) {
	var balance int
	query1 := `SELECT balance FROM users WHERE id = $1`
	row := r.QueryRow(query1, request.UserId)
	_ = row.Scan(&balance)
	amount := int(math.Round(request.AmountOfMoney*100)) + balance
	query2 := `UPDATE users SET balance = $1 WHERE id = $2`
	_, _ = r.Exec(query2, amount, request.UserId)
}

func (r *userRepo) Withdraw(request models.UserRequest) {
	var balance int
	query1 := `SELECT balance FROM users WHERE id = $1`
	row := r.QueryRow(query1, request.UserId)
	_ = row.Scan(&balance)
	amount := int(math.Round(request.AmountOfMoney * 100))
	diff := balance - amount
	if diff > -1 {
		query2 := `UPDATE users SET balance = $1 WHERE id = $2`
		_, _ = r.Exec(query2, diff, request.UserId)
	}
}
