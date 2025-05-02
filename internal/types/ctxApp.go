package types

import "github.com/charmbracelet/bubbles/textinput"

type CtxMain struct {
	Err        error
	SwitchMode bool
	TextInput  textinput.Model
	CtxDict
	CtxTranslate
}

type CtxTranslate struct {
	Text       string
	SwitchLang bool
}

type CtxDict struct {
	Response string
}