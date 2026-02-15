package handlers

import (
	"encoding/json"
	"fmt"

	"gotrack/internal/course"
	"gotrack/internal/database"
	"gotrack/internal/models"
)

// Handler provides the JS bindings for the webview.
type Handler struct {
	db *database.DB
}

// New creates a new Handler with the given database.
func New(db *database.DB) *Handler {
	return &Handler{db: db}
}

// GetModules returns all modules with current progress and unlock status.
func (h *Handler) GetModules() (string, error) {
	completions, err := h.db.GetLessonCompletions()
	if err != nil {
		return "[]", err
	}

	modules := course.AllModules()
	for mi := range modules {
		var modCompleted, modTotal int
		for ci := range modules[mi].Chapters {
			ch := &modules[mi].Chapters[ci]

			// Check unlock status
			if req, ok := course.ExerciseRequirements[ch.Number]; ok {
				unlocked, _ := h.db.IsChapterComplete(req)
				ch.Unlocked = unlocked
			} else {
				ch.Unlocked = true
			}

			// Calculate chapter progress
			var chCompleted int
			for li := range ch.Lessons {
				if completions[ch.Lessons[li].ID] {
					ch.Lessons[li].Completed = true
					chCompleted++
				}
			}
			total := len(ch.Lessons)
			if total > 0 {
				ch.Progress = float64(chCompleted) / float64(total)
			}
			modCompleted += chCompleted
			modTotal += total
		}
		if modTotal > 0 {
			modules[mi].Progress = float64(modCompleted) / float64(modTotal)
		}
	}

	data, err := json.Marshal(modules)
	return string(data), err
}

// ToggleLesson toggles a lesson's completed status and returns updated modules.
func (h *Handler) ToggleLesson(lessonID string) (string, error) {
	_, err := h.db.ToggleLesson(lessonID)
	if err != nil {
		return "", err
	}
	return h.GetModules()
}

// IsChapterUnlocked checks whether a chapter's exercises are available.
func (h *Handler) IsChapterUnlocked(chapter int) (bool, error) {
	req, ok := course.ExerciseRequirements[chapter]
	if !ok {
		return true, nil
	}
	return h.db.IsChapterComplete(req)
}

// SaveTimerSession records a completed pomodoro session.
func (h *Handler) SaveTimerSession(minutes, chapter int) error {
	return h.db.SaveTimerSession(minutes, chapter)
}

// GetStats returns overall progress statistics as JSON.
func (h *Handler) GetStats() (string, error) {
	stats, err := h.db.GetStats()
	if err != nil {
		return "{}", err
	}
	data, err := json.Marshal(stats)
	return string(data), err
}

// SeedLessons populates the database with course data.
func (h *Handler) SeedLessons() error {
	chapters := course.AllChapters()
	var all []models.Lesson
	for _, ch := range chapters {
		all = append(all, ch.Lessons...)
	}
	return h.db.SeedLessons(all)
}

// GetYouTubeURL returns the YouTube embed URL for a lesson.
func (h *Handler) GetYouTubeURL(lessonID string) (string, error) {
	ytID, err := h.db.GetYouTubeID(lessonID)
	if err != nil {
		return "", fmt.Errorf("lesson not found: %w", err)
	}
	return "https://www.youtube.com/embed/" + ytID, nil
}
