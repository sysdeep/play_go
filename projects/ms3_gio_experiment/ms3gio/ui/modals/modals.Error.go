package modals

import (
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func NewError(err error, th *material.Theme) *Message {
	return &Message{
		Title: "Error",
		Body:  err.Error(),
		Type:  MessageTypeErr,
		OKBtn: widget.Clickable{},
		Th:    th,
	}
}
