package models

// Lesson represents a single video lesson in the course.
type Lesson struct {
	ID           string `json:"id"`
	Chapter      int    `json:"chapter"`
	LessonNumber int    `json:"lessonNumber"`
	Title        string `json:"title"`
	YouTubeID    string `json:"youtubeId"`
	IsExercise   bool   `json:"isExercise"`
	Completed    bool   `json:"completed"`
	CompletedAt  string `json:"completedAt,omitempty"`
}

// Chapter groups lessons under a numbered chapter.
type Chapter struct {
	Number     int       `json:"number"`
	Title      string    `json:"title"`
	Lessons    []Lesson  `json:"lessons"`
	IsExercise bool      `json:"isExercise"`
	RequiredBy int       `json:"requiredBy,omitempty"` // chapter that must be completed first
	Unlocked   bool      `json:"unlocked"`
	Progress   float64   `json:"progress"` // 0.0 to 1.0
}

// Module groups chapters into learning stages.
type Module struct {
	Name     string    `json:"name"`
	Chapters []Chapter `json:"chapters"`
	Progress float64   `json:"progress"`
}

// TimerSession records a completed pomodoro session.
type TimerSession struct {
	ID              int    `json:"id"`
	StartedAt       string `json:"startedAt"`
	DurationMinutes int    `json:"durationMinutes"`
	Chapter         int    `json:"chapter,omitempty"`
}

// Stats holds overall progress statistics.
type Stats struct {
	TotalLessons     int     `json:"totalLessons"`
	CompletedLessons int     `json:"completedLessons"`
	Progress         float64 `json:"progress"`
	TotalMinutes     int     `json:"totalMinutes"`
}
