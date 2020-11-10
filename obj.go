package api

import "image"

type StreamDeckInfo struct {
	Cols int `json:"cols,omitempty"`
	Rows int `json:"rows,omitempty"`
	IconSize int `json:"icon_size,omitempty"`
	Page int `json:"page"`
}

type Page []Key

type Config struct {
	Modules []string `json:"modules,omitempty"`
	Pages []Page `json:"pages"`
}

type Key struct {
	Icon       string `json:"icon,omitempty"`
	SwitchPage int   `json:"switch_page,omitempty"`
	Text       string `json:"text,omitempty"`
	TextSize	int	`json:"text_size,omitempty"`
	TextAlignment string `json:"text_alignment,omitempty"`
	Keybind    string `json:"keybind,omitempty"`
	Command    string `json:"command,omitempty"`
	Brightness int   `json:"brightness,omitempty"`
	Url        string `json:"url,omitempty"`
	IconHandler    string `json:"icon_handler,omitempty"`
	KeyHandler string `json:"key_handler,omitempty"`
	IconHandlerFields map[string]string `json:"icon_handler_fields,omitempty"`
	KeyHandlerFields map[string]string `json:"key_handler_fields,omitempty"`
	Buff image.Image `json:"-"`
	IconHandlerStruct IconHandler `json:"-"`
	KeyHandlerStruct KeyHandler `json:"-"`
}

type Module struct {
	Name string
	IconFields []Field
	KeyFields []Field
	IsIcon bool
	IsKey bool
}


type Field struct {
	Title string
	Name string
	Type string
	FileTypes []string
}