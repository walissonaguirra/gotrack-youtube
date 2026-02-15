package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"

	webview "github.com/webview/webview_go"

	"gotrack/internal/database"
	"gotrack/internal/handlers"
)

//go:embed frontend/dist
var frontendFS embed.FS

func main() {
	// Initialize database
	db, err := database.New()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize handlers and seed course data
	h := handlers.New(db)
	if err := h.SeedLessons(); err != nil {
		log.Fatalf("Failed to seed lessons: %v", err)
	}

	// Serve embedded frontend files via local HTTP server
	frontendContent, err := fs.Sub(frontendFS, "frontend/dist")
	if err != nil {
		log.Fatalf("Failed to access frontend files: %v", err)
	}

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatalf("Failed to start HTTP listener: %v", err)
	}
	defer listener.Close()

	go http.Serve(listener, http.FileServer(http.FS(frontendContent)))
	addr := fmt.Sprintf("http://%s", listener.Addr().String())

	// Create webview window
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("GoTrack - Aprenda Go")
	w.SetSize(1280, 800, webview.HintNone)

	// Bind Go functions to JavaScript
	w.Bind("goGetModules", h.GetModules)
	w.Bind("goToggleLesson", h.ToggleLesson)
	w.Bind("goIsChapterUnlocked", h.IsChapterUnlocked)
	w.Bind("goSaveTimerSession", h.SaveTimerSession)
	w.Bind("goGetStats", h.GetStats)
	w.Bind("goGetYouTubeURL", h.GetYouTubeURL)

	w.Navigate(addr)
	w.Run()
}
