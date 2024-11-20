package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func setupTestDirectory() (string, func()) {
	dir, err := os.MkdirTemp("", "fileserver_test")
	if err != nil {
		panic(err)
	}

	os.WriteFile(filepath.Join(dir, "test.txt"), []byte("Group D Test Module"), 0644)
	os.WriteFile(filepath.Join(dir, "image.png"), []byte("\x89PNG\r\n\x1a\n"), 0644)

	return dir, func() {
		os.RemoveAll(dir)
	}
}

func TestFileServer(t *testing.T) {
	dir, cleanup := setupTestDirectory()
	defer cleanup()

	fs := http.FileServer(http.Dir(dir))
	server := httptest.NewServer(fs)
	defer server.Close()

	t.Run("Serve existing file", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/test.txt")
		if err != nil {
			t.Fatalf("Failed to send GET request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
		}

		body, _ := io.ReadAll(resp.Body)
		if string(body) != "Hello, World!" {
			t.Errorf("Unexpected body content: %s", body)
		}
	})

	t.Run("Non-existent file", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/nonexistent.txt")
		if err != nil {
			t.Fatalf("Failed to send GET request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Expected status 404 Not Found, got %d", resp.StatusCode)
		}
	})

	t.Run("MIME type for text file", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/test.txt")
		if err != nil {
			t.Fatalf("Failed to send GET request: %v", err)
		}
		defer resp.Body.Close()

		if contentType := resp.Header.Get("Content-Type"); !strings.HasPrefix(contentType, "text/plain") {
			t.Errorf("Expected Content-Type 'text/plain', got %s", contentType)
		}
	})

	t.Run("MIME type for image file", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/image.png")
		if err != nil {
			t.Fatalf("Failed to send GET request: %v", err)
		}
		defer resp.Body.Close()

		if contentType := resp.Header.Get("Content-Type"); !strings.HasPrefix(contentType, "image/png") {
			t.Errorf("Expected Content-Type 'image/png', got %s", contentType)
		}
	})
}
