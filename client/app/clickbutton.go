//ClickButton.go

package app

import (
	"myitcv.io/react"
)

//go:generate reactGen

type ClickButtonDef struct {
	react.ComponentDef
}

type ClickButtonProps struct {
	Title      string
	Handler    AppSignHandler
	HandlerFlg bool
}

func ClickButton(p ClickButtonProps) *ClickButtonElem {
	return buildClickButtonElem(p)
}

type clickButtonChange struct {
	h AppSignHandler
	f bool
}

func (bc clickButtonChange) OnClick(e *react.SyntheticMouseEvent) {
	handler := bc.h
	newFlag := bc.f

	handler.HandleClickButton(newFlag)

	e.PreventDefault()
}

func (r ClickButtonDef) Render() react.Element {
	return react.A(&react.AProps{ClassName: "btn btn-primary", OnClick: clickButtonChange{h: r.Props().Handler, f: r.Props().HandlerFlg}, Href: "#"}, react.S(r.Props().Title))
}
