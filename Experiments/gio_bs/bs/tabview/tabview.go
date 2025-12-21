package tabview

import (
	"giobs/bs/misc"
	"giobs/bs/theme"
	"image"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

var (
	horizontalInset = layout.Inset{Left: unit.Dp(2)}
	verticalInset   = layout.Inset{Top: unit.Dp(2)}
	horizontalFlex  = layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}
	verticalFlex    = layout.Flex{Axis: layout.Horizontal, Alignment: layout.Start}
)

type TabView struct {
	Axis        layout.Axis
	list        layout.List
	tabItems    []*TabItem
	currentView int
	headerSize  int
	bodySize    int
}

func (tv *TabView) Layout(gtx C, th *theme.Theme) D {
	tv.Update(gtx)

	if len(tv.tabItems) <= 0 {
		return layout.Dimensions{}
	}

	maxTabSize := tv.calculateWidth(gtx, th)
	var direction layout.Direction
	var flex layout.Flex
	var tabAlign layout.Direction
	if tv.Axis == layout.Horizontal {
		direction = layout.Center
		flex = horizontalFlex
		tabAlign = layout.Center
	} else {
		direction = layout.N
		flex = verticalFlex
		tabAlign = layout.W
	}

	return flex.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return direction.Layout(gtx, func(gtx C) D {
				tv.list.Axis = tv.Axis
				tv.list.Alignment = layout.Start
				listDims := tv.list.Layout(gtx, len(tv.tabItems), func(gtx C, index int) D {
					gtx.Constraints.Min = maxTabSize
					item := tv.tabItems[index]
					item.alignment = tabAlign

					if index == 0 {
						return item.LayoutTab(gtx, th)
					}

					if tv.Axis == layout.Horizontal {
						return horizontalInset.Layout(gtx, func(gtx C) D {
							return item.LayoutTab(gtx, th)
						})
					} else {
						return verticalInset.Layout(gtx, func(gtx C) D {
							return item.LayoutTab(gtx, th)
						})
					}

				})

				if tv.Axis == layout.Horizontal {
					tv.headerSize = listDims.Size.X
				} else {
					tv.headerSize = listDims.Size.Y
				}
				return listDims
			})
		}),
		layout.Rigid(func(gtx C) D {
			if tv.Axis == layout.Horizontal {
				return layout.Spacer{Height: unit.Dp(2)}.Layout(gtx)
			} else {
				return layout.Spacer{Width: unit.Dp(24)}.Layout(gtx)
			}
		}),

		layout.Rigid(func(gtx C) D {
			if tv.Axis == layout.Horizontal {
				gtx.Constraints.Min.X = tv.headerSize
			} else {
				gtx.Constraints.Min.Y = max(tv.headerSize, tv.bodySize)
			}
			return misc.Divider(tv.Axis, unit.Dp(0.5)).Layout(gtx, th)
		}),

		layout.Rigid(func(gtx C) D {
			if tv.Axis == layout.Horizontal {
				return layout.Spacer{Height: unit.Dp(24)}.Layout(gtx)
			} else {
				return layout.Spacer{Width: unit.Dp(24)}.Layout(gtx)
			}
		}),

		layout.Rigid(func(gtx C) D {
			dims := tv.tabItems[tv.currentView].LayoutWidget(gtx, th)
			if tv.Axis == layout.Vertical {
				tv.bodySize = dims.Size.Y
				gtx.Execute(op.InvalidateCmd{})
			}

			return dims
		}),
	)
}

func (tv *TabView) Update(gtx C) {
	for idx, item := range tv.tabItems {
		if item.Update(gtx) {
			// unselect last item
			lastItem := tv.tabItems[tv.currentView]
			if lastItem != nil && idx != tv.currentView {
				lastItem.selected = false
			}

			tv.currentView = idx
		}

		if tv.currentView == idx && !item.selected {
			item.selected = true
		}
	}
}

func (tv *TabView) CurrentTab() int {
	return tv.currentView
}

func (tv *TabView) calculateWidth(gtx C, th *theme.Theme) image.Point {
	fakeOps := new(op.Ops)
	current := gtx.Ops
	gtx.Ops = fakeOps
	maxSize := image.Point{}

	gtx.Constraints.Min = image.Point{}
	for _, item := range tv.tabItems {
		dims := item.layoutTab(gtx, th)
		if dims.Size.X > maxSize.X {
			maxSize.X = dims.Size.X
		}
		// if dims.Size.Y > maxSize.Y {
		// 	maxSize.Y = dims.Size.Y
		// }
	}

	gtx.Ops = current
	return maxSize
}

func NewTabView(axis layout.Axis, item ...*TabItem) *TabView {
	return &TabView{
		Axis:     axis,
		tabItems: item,
	}
}
