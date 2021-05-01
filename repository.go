package main

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type extensions []string

var imageExtensions = extensions{
	"jpeg",
	"jpg",
	"png",
	"tiff",
	"gif",
}

// Repository is a repository of images located on the local machine
type Repository struct {
	root string
	fs   fs.FS
}

// NewRepository initializes the application's state with an image repository located at the path `root`
func NewRepository(root string) *Repository {
	return &Repository{
		root: root,
		fs:   os.DirFS(root),
	}
}

// GetImages returns a slice of relative files paths for all the images in the repository
func (r *Repository) GetImages() []fs.DirEntry {
	images := make([]fs.DirEntry, 0)

	fs.WalkDir(r.fs, ".", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path)
		if isImage(path) {
			images = append(images, d)
		}
		return nil
	})

	return images
}

// isImage returns true if the provided file path is an image
func isImage(path string) bool {
	ext := strings.TrimPrefix(filepath.Ext(path), ".")
	return imageExtensions.contains(ext)
}

// contains returns true if s is an element of l
func (l extensions) contains(s string) bool {
	for _, e := range l {
		if e == s {
			return true
		}
	}
	return false
}

// AddImage downloads the image at the provided url and adds it to the repository
func (r *Repository) AddImage(url string) error {
	if !isImage(url) {
		return fmt.Errorf("not a supported image type")
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching image: %w", err)
	}
	defer resp.Body.Close()

	image, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading image: %w", err)
	}
	err = os.WriteFile("/Users/adammitha/Downloads/test.jpeg", image, 0666)
	if err != nil {
		return fmt.Errorf("error saving file: %w", err)
	}

	return nil
}
