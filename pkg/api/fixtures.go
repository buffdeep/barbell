package api

import (
	"bhavdeep.me/barbell/pkg/models"
)

type newWorkoutQuery struct {
	Title     string            `json:"string"`
	Exercises []models.Exercise `json:"exercises"`
}
