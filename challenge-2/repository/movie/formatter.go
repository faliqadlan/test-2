package movie

type GetResponse struct {
	Movie_uid   string `json:"movie_uid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	Artist      string `json:"artist"`
	Genres      string `json:"genres"`
	Image       string `json:"image"`
}

type GetResponses struct {
	Responses []GetResponse `json:"responses"`
}
