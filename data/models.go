package data

//BookData is the struct for holding the model to store
type BookData struct {
	ID          string `json:"book_id"`
	Title       string `json:"title"`
	Authors     string `json:"authors"`
	ISBN        string `json:"isbn"`
	Description string `json:"description"`
}
