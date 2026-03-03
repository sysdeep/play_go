package containers

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

type Containers struct {
	closeButton widget.Clickable
}

func NewContainers() *Containers {
	return &Containers{}
}

func (s *Containers) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return material.Button(th, &s.closeButton, "Containers close").Layout(gtx)
}

func (s *Containers) NavItem() component.NavItem {

	icon, _ := widget.NewIcon(icons.ActionSettings)

	return component.NavItem{
		Name: "Containers",
		Icon: icon,
	}
}
