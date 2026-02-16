package models

// Lesson representa uma aula em vídeo do curso.
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

// Chapter agrupa aulas sob um capítulo numerado.
type Chapter struct {
	Number     int       `json:"number"`
	Title      string    `json:"title"`
	Lessons    []Lesson  `json:"lessons"`
	IsExercise bool      `json:"isExercise"`
	RequiredBy int       `json:"requiredBy,omitempty"` // capítulo que deve ser concluído primeiro
	Unlocked   bool      `json:"unlocked"`
	Progress   float64   `json:"progress"` // 0.0 to 1.0
}

// Module agrupa capítulos em etapas de aprendizado.
type Module struct {
	Name     string    `json:"name"`
	Chapters []Chapter `json:"chapters"`
	Progress float64   `json:"progress"`
}

// TimerSession registra uma sessão pomodoro concluída.
type TimerSession struct {
	ID              int    `json:"id"`
	StartedAt       string `json:"startedAt"`
	DurationMinutes int    `json:"durationMinutes"`
	Chapter         int    `json:"chapter,omitempty"`
}

// Stats armazena as estatísticas gerais de progresso.
type Stats struct {
	TotalLessons     int     `json:"totalLessons"`
	CompletedLessons int     `json:"completedLessons"`
	Progress         float64 `json:"progress"`
	TotalMinutes     int     `json:"totalMinutes"`
}
