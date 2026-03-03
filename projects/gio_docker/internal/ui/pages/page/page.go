package page

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

type Page interface {
	Layout(gtx layout.Context, th *material.Theme) layout.Dimensions
	NavItem() component.NavItem
}
