package api

import (
	"bhavdeep.me/barbell/pkg/models"
	"github.com/gin-gonic/gin"
)

func handleCreateWorkout(c *gin.Context) {
	validator := &newWorkoutQuery{}
	if err := c.ShouldBindJSON(validator); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	if len(validator.Exercises) < 1 {
		c.JSON(400, gin.H{"msg": "No exercises provided"})
		return
	}
	exercises, err := parseExercises(validator)
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	// Logic here to get the user's ID
	userID := "SAMPLE_USER_ID"
	model := models.NewWorkout(
		validator.Title,
		userID,
		exercises,
	)
	// Persist the model in DB
	c.JSON(200, model)
}
