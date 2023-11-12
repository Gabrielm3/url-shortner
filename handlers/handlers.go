package handlers

import (
	"fmt"
	"net/http"

	"github.com/Gabrielm3/url-shortner/services"
	"github.com/Gabrielm3/url-shortner/templates"
)

func HandleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/shorten", http.StatusSeeOther)
		return
	}

	templates.RenderTemplate(w, "form.html", nil)
}

func HandleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid operation", http.StatusMethodNotAllowed)
		return
	}

	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "No URL present", http.StatusBadRequest)
		return
	}

	shortKey := services.GenerateShortKey()
	services.Url[shortKey] = originalURL

	shortenedURL := fmt.Sprintf("http://localhost:8080/short/%s", shortKey)

	data := map[string]string{
		"OriginalURL": originalURL,
		"ShortenedURL": shortenedURL,
	}

	templates.RenderTemplate(w, "result.html", data)
}

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	services.RedirectShortKey(w, r)
}
