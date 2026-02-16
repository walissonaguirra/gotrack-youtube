package handlers

import (
	"encoding/json"
	"fmt"

	"gotrack/internal/course"
	"gotrack/internal/database"
	"gotrack/internal/models"
)

// Handler fornece as vinculações JS para o webview.
type Handler struct {
	db *database.DB
}

// New cria um novo Handler com o banco de dados fornecido.
func New(db *database.DB) *Handler {
	return &Handler{db: db}
}

// GetModules retorna todos os módulos com progresso atual e status de desbloqueio.
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

			// Verifica status de desbloqueio
			if req, ok := course.ExerciseRequirements[ch.Number]; ok {
				unlocked, _ := h.db.IsChapterComplete(req)
				ch.Unlocked = unlocked
			} else {
				ch.Unlocked = true
			}

			// Calcula o progresso do capítulo
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

// ToggleLesson alterna o status de conclusão de uma aula e retorna os módulos atualizados.
func (h *Handler) ToggleLesson(lessonID string) (string, error) {
	_, err := h.db.ToggleLesson(lessonID)
	if err != nil {
		return "", err
	}
	return h.GetModules()
}

// IsChapterUnlocked verifica se os exercícios de um capítulo estão disponíveis.
func (h *Handler) IsChapterUnlocked(chapter int) (bool, error) {
	req, ok := course.ExerciseRequirements[chapter]
	if !ok {
		return true, nil
	}
	return h.db.IsChapterComplete(req)
}

// SaveTimerSession registra uma sessão pomodoro concluída.
func (h *Handler) SaveTimerSession(minutes, chapter int) error {
	return h.db.SaveTimerSession(minutes, chapter)
}

// GetStats retorna as estatísticas gerais de progresso em JSON.
func (h *Handler) GetStats() (string, error) {
	stats, err := h.db.GetStats()
	if err != nil {
		return "{}", err
	}
	data, err := json.Marshal(stats)
	return string(data), err
}

// SeedLessons popula o banco de dados com os dados do curso.
func (h *Handler) SeedLessons() error {
	chapters := course.AllChapters()
	var all []models.Lesson
	for _, ch := range chapters {
		all = append(all, ch.Lessons...)
	}
	return h.db.SeedLessons(all)
}

// GetYouTubeURL retorna a URL de embed do YouTube para uma aula.
func (h *Handler) GetYouTubeURL(lessonID string) (string, error) {
	ytID, err := h.db.GetYouTubeID(lessonID)
	if err != nil {
		return "", fmt.Errorf("aula não encontrada: %w", err)
	}
	return "https://www.youtube.com/embed/" + ytID, nil
}
