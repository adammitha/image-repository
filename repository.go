package main

import (
	"io/fs"
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

// Repository stores the application's state
type Repository struct {
	fs fs.FS
}

// NewRepository initializes the application's state with an image repository located at the path `root`
func NewRepository(root string) *Repository {
	return &Repository{
		fs: os.DirFS(root),
	}
}

// GetImages returns a slice of relative files paths for all the images in the repository
func (r *Repository) GetImages() []string {
	images := make([]string, 0)

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
