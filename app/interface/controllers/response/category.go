package response

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Categories = []Category