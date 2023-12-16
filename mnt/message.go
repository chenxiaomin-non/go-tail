package mnt

import "encoding/json"

type Message struct {
	FilePath string `json:"file_path"`
	Content  string `json:"content"`
	Error    string `json:"error"`
}

// new message from observer without error
func NewMessage(filePath, content string) *Message {
	return &Message{
		FilePath: filePath,
		Content:  content,
		Error:    "nil",
	}
}

// new message from observer with error
func NewErrorMessage(filePath string, err error) *Message {
	return &Message{
		FilePath: filePath,
		Content:  "",
		Error:    err.Error(),
	}
}

func (m Message) ToJson() (string, error) {
	js, err := json.Marshal(m)
	return string(js), err
}
