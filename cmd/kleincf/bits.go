package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path"
)

type Bits struct {
	TempDir string
}

func (b *Bits) DownloadBits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename='application.zip'")

	http.ServeFile(w, r, path.Join(b.TempDir, "application.zip"))
}

func (b *Bits) UploadBits(w http.ResponseWriter, r *http.Request) {
	uploadedFile, _, err := r.FormFile("application")
	if err != nil {
		panic(err)
	}
	defer uploadedFile.Close()

	filename := path.Join(b.TempDir, "application.zip")
	out, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	_, err = io.Copy(out, uploadedFile)
	if err != nil {
		panic(err)
	}

	response := map[string]interface{}{
		"metadata": map[string]interface{}{
			"guid": "job-id",
		},
		"entity": map[string]interface{}{
			"status": "finished",
			"guid":   "job-id",
		},
	}

	bytes, _ := json.Marshal(response)
	w.Write(bytes)
	w.WriteHeader(http.StatusCreated)
}
