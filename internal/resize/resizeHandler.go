package resize

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/disintegration/imaging"
	"github.com/h2non/bimg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type ReSizeImage interface {
	fetch() (*CurrentImage, bool, error)
	ReSize([]byte) (string, error)
	IsReSizeAgain() (bool, []byte, error)
}
type CurrentImage struct {
	Width     int
	Height    int
	Name      string
	Size      int
	ImageByte []byte
}

type fileResize struct {
	url string
}

func NewResize(url string) ReSizeImage {
	return &fileResize{
		url: url,
	}
}
func (f *fileResize) fetch() (*CurrentImage, bool, error) {
	response, err := http.Get(f.url)
	defer response.Body.Close()

	if err != nil {
		return nil, false, err
	}

	if response.StatusCode != consts.StatusOK {
		return nil, false, errors.New("received non 200 response code")
	}
	body, err := ioutil.ReadAll(response.Body)

	img, err := imaging.Decode(bytes.NewReader(body))

	if err != nil {
		return nil, false, err
	}
	size, _ := strconv.Atoi(response.Header.Get("Content-Length"))

	if err != nil {
		return nil, false, err
	}

	return &CurrentImage{
		Width:     img.Bounds().Dx(),
		Height:    img.Bounds().Dy(),
		Name:      strings.Split(filepath.Base(f.url), "?")[0],
		Size:      size,
		ImageByte: body,
	}, true, nil
}

func (f *fileResize) ReSize(data []byte) (string, error) {
	newImage, err := bimg.NewImage(data).Resize(getConfigResize())

	if err != nil {
		return "", err
	}

	fileName := strings.Split(filepath.Base(f.url), "?")[0]
	p, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	path := fmt.Sprintf("%s/assets/%s", p, fileName)
	if err = createFile(path); err != nil {
		return "", err
	}
	err = bimg.Write(path, newImage)

	if err != nil {
		return "", err
	}

	return fileName, nil
}

// this func create in cdn or do or s3
func createFile(file string) error {
	f, err := os.Create(file)

	defer f.Close()

	if err != nil {
		return err
	}

	return nil
}

func (f *fileResize) IsReSizeAgain() (bool, []byte, error) {

	result, ok, err := f.fetch()
	if !ok {
		return true, nil, err
	}

	width, height := getConfigResize()

	if result.Height == height && result.Width == width {
		return true, nil, errors.New("the image is resize")
	}

	return false, result.ImageByte, nil
}
