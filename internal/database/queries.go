package database

import (
	"time"

	"gotrack/internal/models"
)

// SeedLessons inserts lessons that don't already exist in the database.
func (db *DB) SeedLessons(lessons []models.Lesson) error {
	stmt, err := db.conn.Prepare(`
		INSERT INTO lessons (id, chapter, lesson_number, title, youtube_id, is_exercise)
		VALUES (?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			chapter = excluded.chapter,
			lesson_number = excluded.lesson_number,
			title = excluded.title,
			youtube_id = excluded.youtube_id,
			is_exercise = excluded.is_exercise
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, l := range lessons {
		isEx := 0
		if l.IsExercise {
			isEx = 1
		}
		if _, err := stmt.Exec(l.ID, l.Chapter, l.LessonNumber, l.Title, l.YouTubeID, isEx); err != nil {
			return err
		}
	}
	return nil
}

// ToggleLesson toggles the completed status of a lesson, returning the new state.
func (db *DB) ToggleLesson(id string) (bool, error) {
	var completed int
	err := db.conn.QueryRow("SELECT completed FROM lessons WHERE id = ?", id).Scan(&completed)
	if err != nil {
		return false, err
	}

	newState := 1 - completed
	completedAt := ""
	if newState == 1 {
		completedAt = time.Now().Format(time.RFC3339)
	}

	_, err = db.conn.Exec(
		"UPDATE lessons SET completed = ?, completed_at = ? WHERE id = ?",
		newState, completedAt, id,
	)
	return newState == 1, err
}

// GetLessonCompletions returns a map of lesson ID -> completed status.
func (db *DB) GetLessonCompletions() (map[string]bool, error) {
	rows, err := db.conn.Query("SELECT id, completed FROM lessons")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]bool)
	for rows.Next() {
		var id string
		var completed int
		if err := rows.Scan(&id, &completed); err != nil {
			return nil, err
		}
		result[id] = completed == 1
	}
	return result, rows.Err()
}

// IsChapterComplete checks if all lessons in a chapter are completed.
func (db *DB) IsChapterComplete(chapter int) (bool, error) {
	var total, completed int
	err := db.conn.QueryRow(
		"SELECT COUNT(*), COALESCE(SUM(completed), 0) FROM lessons WHERE chapter = ?",
		chapter,
	).Scan(&total, &completed)
	if err != nil {
		return false, err
	}
	return total > 0 && total == completed, nil
}

// SaveTimerSession records a completed pomodoro session.
func (db *DB) SaveTimerSession(minutes, chapter int) error {
	_, err := db.conn.Exec(
		"INSERT INTO timer_sessions (started_at, duration_minutes, chapter) VALUES (?, ?, ?)",
		time.Now().Format(time.RFC3339), minutes, chapter,
	)
	return err
}

// GetYouTubeID returns the YouTube ID for a lesson.
func (db *DB) GetYouTubeID(lessonID string) (string, error) {
	var ytID string
	err := db.conn.QueryRow("SELECT youtube_id FROM lessons WHERE id = ?", lessonID).Scan(&ytID)
	return ytID, err
}

// GetStats returns overall progress statistics.
func (db *DB) GetStats() (models.Stats, error) {
	var stats models.Stats

	err := db.conn.QueryRow(
		"SELECT COUNT(*), COALESCE(SUM(completed), 0) FROM lessons",
	).Scan(&stats.TotalLessons, &stats.CompletedLessons)
	if err != nil {
		return stats, err
	}

	if stats.TotalLessons > 0 {
		stats.Progress = float64(stats.CompletedLessons) / float64(stats.TotalLessons)
	}

	err = db.conn.QueryRow(
		"SELECT COALESCE(SUM(duration_minutes), 0) FROM timer_sessions",
	).Scan(&stats.TotalMinutes)

	return stats, err
}
