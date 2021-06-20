package domain

// Parameters data structure
type Parameters struct {
	SearchWord string `json:"searchword" query:"searchword" `
	Pagination int    `json:"pagination" query:"page" validate:"omitempty,numeric"`
}

type Movies struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}

// MovieDetail data structure
type MovieDetail struct {
	Title      string         `json:"title"`
	Year       string         `json:"year"`
	Rated      string         `json:"rated"`
	Released   string         `json:"released"`
	Runtime    string         `json:"runtime"`
	Genre      string         `json:"genre"`
	Director   string         `json:"director"`
	Writer     string         `json:"writer"`
	Actors     string         `json:"actors"`
	Plot       string         `json:"plot"`
	Language   string         `json:"language"`
	Country    string         `json:"country"`
	Awards     string         `json:"awards"`
	Poster     string         `json:"poster"`
	Ratings    []RatingDetail `json:"ratings"`
	MetaScore  string         `json:"metaScore"`
	ImdbRating string         `json:"imdbRating"`
	ImdbVotes  string         `json:"imdbVotes"`
	ImdbID     string         `json:"imdbID"`
	Type       string         `json:"type"`
	DVD        string         `json:"dvd"`
	BoxOffice  string         `json:"boxOffice"`
	Production string         `json:"production"`
	Website    string         `json:"website"`
}

type RatingDetail struct {
	Source string `json:"source"`
	Value  string `json:"value"`
}
