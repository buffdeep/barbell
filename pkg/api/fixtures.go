package api

type ExecrciseQuery struct {
	Name        string  `json:"name"`
	MuscleGroup string  `json:"muscle_group"`
	Weight      float64 `json:"weight"`
	Sets        int     `json:"sets"`
	Reps        int     `json:"reps"`
	Time        string  `json:"time"`
}

type newWorkoutQuery struct {
	Title     string           `json:"string"`
	Exercises []ExecrciseQuery `json:"exercises"`
}
