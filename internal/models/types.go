package models

type DictionaryEntry struct {
	Word         string `json:"word"`
	PartOfSpeech string `json:"partOfSpeech"`
	Definition   string `json:"definition"`
	Example      string `json:"example"`
	Synonyms     string `json:"synonyms"`
}
