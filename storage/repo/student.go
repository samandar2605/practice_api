package repo

import "time"

type Student struct {
	Id          int       `db:"id"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	UserName    string    `db:"user_name"`
	Email       string    `db:"email"`
	PhoneNumber string    `db:"phone_number"`
	CreatedAt   time.Time `db:"created_at"`
}

type StudentStorageI interface {
	Create(u []*Student) (error)
	GetAll(param GetStudentsQuery) (*GetAllStudentsResult, error)
}

type GetStudentsQuery struct {
	Page       int
	Limit      int
	Search     string
	SortByDate string
	SortByName string
}

type GetAllStudentsResult struct {
	Students []*Student
	Count    int
}
