package start

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

type Start struct {
	closeButton widget.Clickable
}

func NewStart() *Start {
	return &Start{}
}

func (s *Start) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return material.Button(th, &s.closeButton, "Start").Layout(gtx)
}

func (s *Start) NavItem() component.NavItem {

	icon, _ := widget.NewIcon(icons.ActionSettings)

	return component.NavItem{
		Name: "Start",
		Icon: icon,
	}
}

// var SettingsIcon *widget.Icon = func() *widget.Icon {
// 	icon, _ := widget.NewIcon(icons.ActionSettings)
// 	return icon
// }()
