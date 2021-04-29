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
