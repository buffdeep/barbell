package api

import (
	"bhavdeep.me/barbell/pkg/models"
)

func parseExercises(workout *newWorkoutQuery) ([]models.Exercise, error) {
	var exercises []models.Exercise
	for _, dirtyExercise := range workout.Exercises {
		if cleanExercise, err := models.NewExercise(
			dirtyExercise.Name,
			dirtyExercise.MuscleGroup,
			dirtyExercise.Weight,
			dirtyExercise.Sets,
			dirtyExercise.Reps,
			dirtyExercise.Time,
		); err != nil {
			return exercises, err
		} else {
			exercises = append(exercises, *cleanExercise)
		}
	}
	return exercises, nil
}
