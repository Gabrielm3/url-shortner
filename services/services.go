package services

import (
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var Url = make(map[string]string)

func GenerateShortKey() string {
	const char = "acbaxwrfeqwoeriuhiugb234114239750932"
	const keyLength = 8

	rand.Seed(time.Now().UnixNano())
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = char[rand.Intn(len(char))]
	}
	return string(shortKey)
}

func RedirectShortKey(w http.ResponseWriter, r *http.Request) {
	shortKey := strings.TrimPrefix(r.URL.Path, "/short/")
	if shortKey == "" {
		http.Error(w, "Shortened key is missing", http.StatusBadRequest)
		return
	}

	originalURL, found := Url[shortKey]
	if !found {
		http.Error(w, "Shortened key not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, originalURL, http.StatusPermanentRedirect)
}
