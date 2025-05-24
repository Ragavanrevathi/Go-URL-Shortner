package request

type ShortURL struct {
	URL string `json:"url" binding:"required"`
}
