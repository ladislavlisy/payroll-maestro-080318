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
	OnClickDef AppDef
	OnClickFlg bool
}

func ClickButton(p ClickButtonProps) *ClickButtonElem {
	return buildClickButtonElem(p)
}

type clickButtonChange struct {
	a AppDef
	f bool
}

func (bc clickButtonChange) OnClick(e *react.SyntheticMouseEvent) {
	appDef := bc.a
	newFlag := bc.f
	newAppState := appDef.State()
	newAppState.userLoggedIn = newFlag
	appDef.SetState(newAppState)

	e.PreventDefault()
}

func (r ClickButtonDef) Render() react.Element {
	return react.A(&react.AProps{ClassName: "btn btn-primary", OnClick: clickButtonChange{a: r.Props().OnClickDef, f: r.Props().OnClickFlg}, Href: "#"}, react.S(r.Props().Title))
}
