package users

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	model "github.com/RamadanRangkuti/multifinance-api/models"
	"github.com/google/uuid"
)

// service ini tergantung pada user repositorynya
type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *store {
	return &store{
		db: db,
	}
}

type UserRepository interface {
	UserRegister(req model.Users) (*uuid.UUID, error)
	GetUserDetail(req model.Users) (*model.Users, error)
}

func (s *store) UserRegister(req model.Users) (*uuid.UUID, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Parsing tanggal lahir dari string
	parsedDate, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		return nil, fmt.Errorf("invalid date format, expected YYYY-MM-DD")
	}

	var userID uuid.UUID
	// Insert ke tabel users
	queryUser := `
		INSERT INTO users(
			nik,
		    fullname,
			legal_name,
 			place_of_birth,
 			date_of_birth,
			salary,
			password,
			role
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
		) RETURNING id
	`
	if err := tx.QueryRow(
		queryUser,
		req.NIK,
		req.Fullname,
		req.LegalName,
		req.PlaceOfBirth,
		parsedDate,
		req.Salary,
		req.Password,
		req.Role,
	).Scan(&userID); err != nil {
		return nil, err
	}

	// Definisi tenor dan multiplier
	tenors := []struct {
		Tenor      int
		Multiplier float64
	}{
		{1, 2.0},
		{2, 2.5},
		{3, 3.0},
		{4, 3.5},
	}

	// Insert ke tabel user_limits dalam transaksi yang sama
	queryLimit := `
		INSERT INTO user_limits (user_id, tenor, limit_amount, used_amount)
		VALUES ($1, $2, $3, $4)
	`
	for _, t := range tenors {
		_, err := tx.Exec(queryLimit, userID, t.Tenor, req.Salary*t.Multiplier, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to insert user limits: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &userID, nil
}

func (s *store) GetUserDetail(req model.Users) (*model.Users, error) {
	queryArgs := `
		SELECT
			*
		FROM
		    users
	`

	var queryConditions []string
	if req.NIK != "" {
		queryConditions = append(queryConditions, fmt.Sprintf("nik = '%s'", req.NIK))
	}

	if req.ID != uuid.Nil {
		queryConditions = append(queryConditions, fmt.Sprintf("id = '%v'", req.ID))
	}

	if req.Fullname != "" {
		queryConditions = append(queryConditions, fmt.Sprintf("fullname = '%s'", req.Fullname))
	}

	if len(queryConditions) > 0 {
		queryArgs += " WHERE " + strings.Join(queryConditions, " AND ")
	}

	queryArgs += `
		ORDER BY created_at DESC limit 1
	`

	fmt.Println("Executing query:", queryArgs)

	var response model.Users
	rows, err := s.db.Query(queryArgs)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&response.ID,
			&response.NIK,
			&response.Fullname,
			&response.LegalName,
			&response.PlaceOfBirth,
			&response.DateOfBirth,
			&response.Salary,
			&response.Password,
			&response.Role,
			&response.CreatedAt,
			&response.UpdatedAt,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("no partner found")
			}
			return nil, fmt.Errorf("failed to fetch user data")
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed iterate over user: %v", err)
	}

	return &response, nil
}
