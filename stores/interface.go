package stores

import "go-url-shortner/models"

type URL interface {
	GetShortURL(models.MyURL) (string, error)
}
