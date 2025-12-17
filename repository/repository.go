package repository

import "session-18/database"

type Repository struct {
	AssignmentRepo AssignmentRepository
	SubmissionRepo SubmissionRepo
	UserRepo       UserRepository
}

func NewRepository(db database.PgxIface) Repository {
	return Repository{
		AssignmentRepo: NewAssignmentRepository(db),
		SubmissionRepo: NewSubmissionRepo(db),
		UserRepo:       NewUserRepository(db),
	}
}
