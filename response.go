package parser

type Response struct {
	Kind    string `json:"kind"`
	ID      string `json:"id"`
	LongURL string `json:"longUrl"`
	Status  string `json:"status"`
}
