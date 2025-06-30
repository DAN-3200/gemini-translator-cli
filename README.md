# Gemini Translator CLI

![image](https://github.com/user-attachments/assets/76b35e1d-3a01-466c-bf11-326f3fd9ef5c)

## Description

A command-line translation tool built in Go using the Bubble Tea TUI library and integrated with the Gemini API. This app provides accurate and context-aware sentence `translation` and a `dictionary` feature for English words, including part of speech, definition, example, synonyms, and collocations. Developed as a practical project to learn and master the Go programming language.

## Installation & Usage

```bash
git clone https://github.com/DAN-3200/gemini-translator-cli.git
cd gemini-translator-cli
```

Run directly (for development):

```bash
go run .
```

Build and execute:

```bash
go build -o app.exe
./app.exe
```

## Model

```go
type DictionaryEntry struct {
  Word         string `json:"word"`
  PartOfSpeech string `json:"partOfSpeech"`
  Definition   string `json:"definition"`
  Example      string `json:"example"`
  Synonyms     string `json:"synonyms"`
  Collocations string `json:"collocations"`
}
```

## Links 

- Link Go: [https://pkg.go.dev/golang.org](https://pkg.go.dev/golang.org)
- Link Bubble Tea: [https://github.com/charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea)


## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
