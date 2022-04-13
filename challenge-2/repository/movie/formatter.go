package movie

type GetResponse struct {
	Movies_id string `json:"movies_id"`
	Title     string `json:"title"`
	Description string `json:"description"`
	Duration  string `json:"duration"`
	Artist    string `json:"artist"`
	Genres    string `json:"genres"`
	Image     string `json:"image"`
}

type GetResponses struct {
	Responses []GetResponse `json:"responses"`
}
