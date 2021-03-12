package model

type Channel struct {
	Channel string `json:"channel"`
}

type Message struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}
