package loader

type BookData struct {
	ID          string `json:"book_id"`
	Title       string `json:"title"`
	Authors     string `json:"authors"`
	ISBN        string `json:"isbn"`
	Description string `json:"description"`
}
