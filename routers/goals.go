package routers

import (
	"errors"
	"kv/tross/database"
	"math"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func createGoal(c *fiber.Ctx) error {
	var goal database.CreateGoalParams
	err := c.BodyParser(&goal)
	if err != nil {
		panic(err)
	}
	db := database.DB
	_, err = db.CreateGoal(c.Context(), goal)
	if err != nil {
		panic(err)
	}
	return c.SendString("Created")
}

type GetFirstGoalResponse struct {
	Date time.Time `json:"date"`
}

func getFirstGoalDate(c *fiber.Ctx) error {
	db := database.DB
	userId := c.Query("userId")
	date, err := db.FetchFirstGoalDate(c.Context(), userId)
	if err != nil {
		return err
	}
	return c.JSON(GetFirstGoalResponse{Date: date})
}

type ProgressResult struct {
	RelapsedCount int64   `json:"relapsed"`
	TotalDays     int64   `json:"total"`
	Progress      float32 `json:"progress"`
}

func getProgress(c *fiber.Ctx) error {
	db := database.DB
	goalIdParam := c.Query("goal_id")
	if goalIdParam == "" {
		return errors.New("missing param 'goal_id'")
	}
	goalId, err := strconv.Atoi(goalIdParam)
	if err != nil {
		return err
	}
	progress, err := db.GetProgress(c.Context(), int32(goalId))
	if err != nil {
		return err
	}
	duration := time.Since(progress.CreatedAt)

	days := int64(math.Ceil(duration.Hours() / 24))
	pr := float32(days-progress.Count) / float32(days)

	return c.JSON(ProgressResult{
		RelapsedCount: progress.Count,
		TotalDays:     days,
		Progress:      pr * 100,
	})
}

func RegisterGoalsRouter(app *fiber.App) {
	groupRouter := app.Group("/goals")

	groupRouter.Post("/create", createGoal)
	groupRouter.Get("/first", getFirstGoalDate)
	groupRouter.Get("/progress", getProgress)
}
