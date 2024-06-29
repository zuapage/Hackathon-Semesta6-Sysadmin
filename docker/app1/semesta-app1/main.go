package main

import (
	"fmt"
	"github.com/google/uuid"
	_ "html/template"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func genUUID() string {
	id := uuid.New().String()
	return id
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/" {
		_, err := os.Stat("index.html")
		if os.IsNotExist(err) {
			http.Error(w, "File index.html tidak ditemukan", http.StatusInternalServerError)
			return
		}
		http.ServeFile(w, r, "index.html")
	} else if r.URL.Path == "/aboutus" {
		err := godotenv.Load(".env")
		if err != nil {
			http.Error(w, "Gagal memuat file .env", http.StatusInternalServerError)
			return
		}
		targetURL := os.Getenv("APP2_URL")
		if !strings.HasPrefix(targetURL, "http://") && !strings.HasPrefix(targetURL, "https://") {
			targetURL = "http://" + targetURL
		}
		if targetURL == "" {
			http.Error(w, "URL Web App ke 2 tidak ditemukan", http.StatusInternalServerError)
			return
		}
		req, err := http.NewRequest("GET", targetURL, nil)
		if err != nil {
			http.Error(w, "Gagal membuat request", http.StatusInternalServerError)
			return
		}
		req.Header.Set("X-UID", genUUID())

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Fprintf(w, `<title>Semesta App 2</title>`)
			http.Error(w, "Gagal memuat konten", http.StatusInternalServerError)
			return
		}

		body, err := io.ReadAll(resp.Body)
		err = resp.Body.Close()
		if err != nil {
			return
		}

		if err != nil {
			http.Error(w, "Gagal membaca respons", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%s", body)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1><center>Halaman yang dicari tidak ditemukan</center></h1>")
	}
}

func main() {
	server := &http.Server{
		Addr:              ":3000",
		Handler:           http.HandlerFunc(handlerFunc),
		ReadHeaderTimeout: 3 * time.Second,
	}
	fmt.Println("Server running on port 3000")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	fmt.Println(run())
}

func run() string {
	return "Setup Travis CI for Golang SEMESTA Hackathon"
}
