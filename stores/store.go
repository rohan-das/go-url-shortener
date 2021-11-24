package stores

import (
	"encoding/csv"
	"go-url-shortner/models"
	"io"
	"os"
	"time"

	"github.com/speps/go-hashids"
)

type urlStore struct{}

func New() URL {
	return &urlStore{}
}

func (u *urlStore) GetShortURL(url models.MyURL) (string, error) {
	csvFile, err := os.OpenFile("./data.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return "", err
	}

	reader := csv.NewReader(csvFile)
	for {
		// line represents each row present in file
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}

		// line[0] is the long input URL and line[1] is the short URL present
		// return short URL if already present
		if url.LongURL == line[0] {
			return line[1], nil
		}
	}

	// generate short URL
	hd := hashids.NewData()
	h, _ := hashids.NewWithData(hd)
	id, _ := h.Encode([]int{int(time.Now().Unix())})
	url.ShortURL = "http://localhost:8080/" + id

	// write the short URL in file for future reference
	w := csv.NewWriter(csvFile)
	if err := w.Write([]string{url.LongURL, url.ShortURL}); err != nil {
		return "", err
	}

	w.Flush()

	return url.ShortURL, nil
}
