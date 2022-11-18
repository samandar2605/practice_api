package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/samandar2605/practice_api/storage/repo"
)

type studentRepo struct {
	db *sqlx.DB
}

func NewStudent(db *sqlx.DB) repo.StudentStorageI {
	return &studentRepo{db: db}
}

func (sr *studentRepo) Create(students []*repo.Student) error {
	query := `
		INSERT INTO students(
			first_name,
			last_name,
			username,
			email,
			phone_number
		)values($1,$2,$3,$4,$5)
	`
	tx, err := sr.db.Begin()
	if err != nil {
		return err
	}

	for _, s := range students {
		_, err = tx.Exec(
			query,
			s.FirstName,
			s.LastName,
			s.UserName,
			s.Email,
			s.PhoneNumber,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func (sr *studentRepo) GetAll(param repo.GetStudentsQuery) (*repo.GetAllStudentsResult, error) {
	result := repo.GetAllStudentsResult{
		Students: make([]*repo.Student, 0),
	}

	offset := (param.Page - 1) * param.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", param.Limit, offset)
	filter := ""
	if param.Search != "" {
		str := "%" + param.Search + "%"
		filter += fmt.Sprintf(` 
			where first_name ILIKE '%s' OR last_name ILIKE '%s' OR email ILIKE '%s' 
		OR username ILIKE '%s' OR phone_number ILIKE '%s'`, str, str, str, str, str)
	}

	query := `
		SELECT 
			id,
			first_name,
			last_name,
			username,
			email,
			phone_number,
			created_at
		FROM students
		` + filter

	if param.SortByDate != "none" || param.SortByName != "none" {
		queryDate := " "
		queryName := " "
		if param.SortByDate != "none" {
			queryDate = " created_at " + param.SortByDate
		}
		if param.SortByName != "none" {
			queryName = " first_name " + param.SortByName
		}
		if param.SortByDate != "none" && param.SortByName != "none" {
			query += " order by " + queryDate + " , " + queryName
		} else {
			if param.SortByDate != "none" {
				query += " order by " + queryDate
			} else {
				query += " order by " + queryName
			}
		}
	}

	query += limit
	rows, err := sr.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var usr repo.Student
		if err := rows.Scan(
			&usr.Id,
			&usr.FirstName,
			&usr.LastName,
			&usr.PhoneNumber,
			&usr.Email,
			&usr.UserName,
			&usr.CreatedAt,
		); err != nil {
			return nil, err
		}
		result.Students = append(result.Students, &usr)
	}
	queryCount := `SELECT count(1) FROM students ` + filter
	err = sr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
