package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestContains(t *testing.T) {
	is := is.New(t)

	is.True(imageExtensions.contains("gif"))
	is.True(!imageExtensions.contains("pdf"))
	is.True(!imageExtensions.contains(""))
}

func TestIsImage(t *testing.T) {
	is := is.New(t)

	is.True(isImage("image.jpeg"))
	is.True(isImage("/dir1/dir2/image.png"))
	is.True(!isImage("/dir1/file.pdf"))
}

func TestGetFilename(t *testing.T) {
	is := is.New(t)

	filename, err := getFilename("http://test.com/test.jpg")
	is.Equal(err, nil)
	is.Equal(filename, "test.jpg")
}
