//LinkButton.go

package components

import (
	"myitcv.io/react"
)

//go:generate reactGen

type LinkButtonDef struct {
	react.ComponentDef
}

type LinkButtonProps struct {
	Title string
	Href  string
}

func LinkButton(p LinkButtonProps) *LinkButtonElem {
	return buildLinkButtonElem(p)
}

func (r LinkButtonDef) Render() react.Element {
	return react.A(&react.AProps{ClassName: "btn btn-primary", Href: r.Props().Href}, react.S(r.Props().Title))
}
