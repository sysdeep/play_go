package ui

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type SideNav struct{}

var list = &widget.List{
	List: layout.List{
		Axis: layout.Vertical,
	},
}

func (s SideNav) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {

	widgets := []layout.Widget{
		func(gtx layout.Context) layout.Dimensions {
			l := material.H3(th, "topLabel")
			// l.State = topLabelState
			return l.Layout(gtx)
		},

		func(gtx layout.Context) layout.Dimensions {
			l := material.H4(th, "topLabel")
			// l.State = topLabelState
			return l.Layout(gtx)
		},
	}

	return material.List(th, list).Layout(gtx, len(widgets), func(gtx layout.Context, i int) layout.Dimensions {
		return layout.UniformInset(unit.Dp(16)).Layout(gtx, widgets[i])
	})

}
