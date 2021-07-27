package models

type PushTelegramModel struct {
	Text      string `json:"text"`
	Photo     string `json:"photo"`
	Video     string `json:"video"`
	File      string `json:"file"`
	Animation string `json:"animation"`
}
