package cards

type Card struct {
	ID          string      `json:"id"`
	User        int64       `json:"user"`
	Title       string      `json:"title"`
	Description Description `json:"description"`
	Pictures    []Picture   `json:"pictures"`
}
type Description struct {
	PlainText string `json:"plain_text"`
	Html      string `json:"html"`
}

type Picture struct {
	ID  int64  `json:"id"`
	Url string `json:"url"`
}
