package main

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/charmbracelet/bubbles/textinput"
)

type extensions []string

var imageExtensions = extensions{
	"jpeg",
	"jpg",
	"png",
	"tiff",
	"gif",
	"heic",
	"svg",
}

// Repository is a repository of images located at root
type Repository struct {
	root      string
	fs        fs.FS
	textInput textinput.Model
}

// NewRepository initializes the application's state with an image repository located at the path `root`
func NewRepository(root string) *Repository {
	ti := textinput.NewModel()
	ti.Placeholder = ". (Current directory)"
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 30
	return &Repository{
		root:      root,
		fs:        os.DirFS(root),
		textInput: ti,
	}
}

// GetImages returns a slice of relative files paths for all the images in the repository
func (r *Repository) GetImages() []fs.DirEntry {
	images := make([]fs.DirEntry, 0)

	fs.WalkDir(r.fs, ".", func(path string, d fs.DirEntry, err error) error {
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
func (ext extensions) contains(s string) bool {
	for _, e := range ext {
		if e == s {
			return true
		}
	}
	return false
}

// AddImages consumes a slice of urls to images, retrieves each image and adds it to the repository
func (r *Repository) AddImages(urls []string) []error {
	var errors []error
	errChan := make(chan error, len(urls))
	var wg sync.WaitGroup
	for _, url := range urls {
		url := url
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := r.addImage(url)
			errChan <- err
		}()
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}

// addImage downloads the image at the provided url and adds it to the repository
func (r *Repository) addImage(url string) error {
	if !isImage(url) {
		return fmt.Errorf("not a supported image type")
	}

	filename, err := getFilename(url)
	if err != nil {
		return fmt.Errorf("error parsing url: %w", err)
	}
	filename = filepath.Join(r.root, filename)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching image: %w", err)
	}
	defer resp.Body.Close()

	image, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading image: %w", err)
	}
	err = os.WriteFile(filename, image, 0666)
	if err != nil {
		return fmt.Errorf("error saving image: %w", err)
	}

	return nil
}

// getFilename returns the name of the file at url
func getFilename(path string) (string, error) {
	parsedURL, err := url.Parse(path)
	if err != nil {
		return "", fmt.Errorf("error parsing url: %w", err)
	}
	_, filename := filepath.Split(parsedURL.Path)
	return filename, nil
}
