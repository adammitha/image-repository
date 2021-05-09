**This is my submission for the Shopify Backend Developer Challenge**

# Image Repository

Welcome to the Image Repository command line application. It is a simple tool for downloading an arbitrary number of images from the internet and storing them in a folder on your local machine. It uses goroutines to concurrently fetch images rather than fetching them sequentially, which speeds up the retrieval of large images.

## Installation

If you have a go toolchain installed on your machine, you can install from source with the following command:

```
go install github.com/adammitha/image-repository@latest
```

Alternatively, you can download one of the prebuild binaries from the releases page.

## Usage

The tool is quite simple. It has a single flag which specifies the directory you want to save the images to. You can pass in as many urls as you want.

```
image-repository -dir /Users/me/Downloads url1 url2 url3 ... urln
```

The program checks the url to make sure that it points to a valid image. The image formats it recognizes are:

- jpeg/jpg
- png
- tiff
- gif
- heic
- svg
