package models

import "time"

type Workout struct {
	Title     string     `json:"title"`
	CreatedAt int64      `json:"created_at"`
	CreatedBy string     `json:"created_by"`
	Exercises []Exercise `json:"exercises"`
}

func NewWorkout(title string, createdBy string, exercises []Exercise) *Workout {
	return &Workout{
		Title:     title,
		CreatedAt: time.Now().Unix(),
		CreatedBy: createdBy,
		Exercises: exercises,
	}
}

type Exercise struct {
	Name        string        `json:"name"`
	MuscleGroup string        `json:"muscle_group"`
	Weight      float64       `json:"weight"`
	Sets        int           `json:"sets"`
	Reps        int           `json:"reps"`
	Time        time.Duration `json:"time"`
}

func NewExercise(name string, muscleGroup string, weight float64, sets int, reps int, exTime string) (*Exercise, error) {
	if exTime == "" {
		exTime = "0s"
	}
	if t, err := time.ParseDuration(exTime); err == nil {
		return &Exercise{
			Name:        name,
			MuscleGroup: muscleGroup,
			Weight:      weight,
			Sets:        sets,
			Reps:        reps,
			Time:        t,
		}, nil
	} else {
		return &Exercise{}, err
	}
}
