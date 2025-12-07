package modals

import (
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func NewInfo(title, body string, th *material.Theme) *Message {
	return &Message{
		Title: title,
		Body:  body,
		Type:  MessageTypeInfo,
		OKBtn: widget.Clickable{},
		Th:    th,
	}
}
