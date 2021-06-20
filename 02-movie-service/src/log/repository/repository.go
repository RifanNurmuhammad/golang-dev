package repository

import (
	"github.com/rifannurmuhammad/02-movie-service/component"
	"github.com/rifannurmuhammad/02-movie-service/src/log/repository/interfaces"
	mysql "github.com/rifannurmuhammad/02-movie-service/src/log/repository/mysql"
)

// Repository parent
type Repository struct {
	db            *component.Mysql
	LogActivities interfaces.LogActivitiesRepository
}

// NewRepository create new repository
func NewRepository(db *component.Mysql) *Repository {
	return &Repository{
		db:            db,
		LogActivities: mysql.NewLogActivitiesRepo(db),
	}
}
