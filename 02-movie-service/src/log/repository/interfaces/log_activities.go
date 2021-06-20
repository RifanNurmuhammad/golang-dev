package interfaces

import (
	"context"

	"github.com/rifannurmuhammad/02-movie-service/src/log/domain"
)

// LogActivitiesRepository interface abstraction
type LogActivitiesRepository interface {
	Insert(c context.Context, data *domain.LogActivities) <-chan error
}
