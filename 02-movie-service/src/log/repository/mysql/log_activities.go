package mysql

import (
	"context"
	"time"

	"github.com/rifannurmuhammad/02-movie-service/component"
	"github.com/rifannurmuhammad/02-movie-service/src/log/domain"
	"github.com/rifannurmuhammad/02-movie-service/src/log/repository/interfaces"
)

type logActivitiesRepoGorm struct {
	db *component.Mysql
}

// NewLogRequestRepo create new log request repository
func NewLogActivitiesRepo(db *component.Mysql) interfaces.LogActivitiesRepository {
	return &logActivitiesRepoGorm{db: db}
}

func (r *logActivitiesRepoGorm) Insert(c context.Context, data *domain.LogActivities) <-chan error {
	output := make(chan error)
	go func() {
		defer close(output)

		data.CreatedAt = time.Now()

		dbConnection := r.db.Write

		if err := dbConnection.Create(&data).Error; err != nil {
			output <- err
			return
		}

		output <- nil
	}()

	return output
}
