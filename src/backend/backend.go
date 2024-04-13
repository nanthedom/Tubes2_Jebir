package backend

type FormData struct {
	StartArticle string `json:"startArticle"`
	StartUrl     string `json:"startUrl"`
	EndArticle   string `json:"endArticle"`
	EndUrl       string `json:"endUrl"`
}
