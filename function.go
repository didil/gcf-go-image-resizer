package gcf_go_image_resizer

import (
	"bytes"
	"errors"
	"github.com/disintegration/imaging"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"strconv"
)

// Cloud Function entry point
func ResizeImage(w http.ResponseWriter, r *http.Request) {
	// parse the url query sting into ResizerParams
	p, err := ParseQuery(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// fetch input image and resize
	img, err := FetchAndResizeImage(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// encode output image to jpeg buffer
	encoded, err := EncodeImageToJpg(img)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set Content-Type and Content-Length headers
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(encoded.Len()))

	// write the output image to http response body
	_, err = io.Copy(w, encoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// struct containing the initial query params
type ResizerParams struct {
	url    string
	height int
	width  int
}

// parse/validate the url params
func ParseQuery(r *http.Request) (*ResizerParams, error) {
	var p ResizerParams
	query := r.URL.Query()
	url := query.Get("url")
	if url == "" {
		return &p, errors.New("Url Param 'url' is missing")
	}

	width, _ := strconv.Atoi(query.Get("width"))
	height, _ := strconv.Atoi(query.Get("height"))

	if width == 0 && height == 0 {
		return &p, errors.New("Url Param 'height' or 'width' must be set")
	}

	p = NewResizerParams(url, height, width)

	return &p, nil
}

// ResizerParams factory 
func NewResizerParams(url string, height int, width int) ResizerParams {
	return ResizerParams{url, height, width}
}

// fetch the image from provided url and resize it
func FetchAndResizeImage(p *ResizerParams) (*image.Image, error) {
	var dst image.Image

	// fetch input data
	response, err := http.Get(p.url)
	if err != nil {
		return &dst, err
	}
	// don't forget to close the response
	defer response.Body.Close()

	// decode input data to image
	src, _, err := image.Decode(response.Body)
	if err != nil {
		return &dst, err
	}

	// resize input image
	dst = imaging.Resize(src, p.width, p.height, imaging.Lanczos)

	return &dst, nil
}

// encode image to jpeg
func EncodeImageToJpg(img *image.Image) (*bytes.Buffer, error) {
	encoded := &bytes.Buffer{}
	err := jpeg.Encode(encoded, *img, nil)
	return encoded, err
}
