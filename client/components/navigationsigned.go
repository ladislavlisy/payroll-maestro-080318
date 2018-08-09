//NavigationSigned.go

package components

import (
	"myitcv.io/react"
)

type navigationTabSigned uint32

//go:generate reactGen

type NavigationSignedDef struct {
	react.ComponentDef
}

type NavigationSignedProps struct {
	BrandName string
}

type NavigationSignedState struct {
	activeNavTab navigationTabSigned
}

func NavigationSigned(p NavigationSignedProps) *NavigationSignedElem {
	return buildNavigationSignedElem(p)
}

const (
	navCompanies navigationTabSigned = iota
)

type navigationTabSignedChange struct {
	n NavigationSignedDef
	t navigationTabSigned
}

func (nc navigationTabSignedChange) OnClick(e *react.SyntheticMouseEvent) {
	nc.n.SetState(NavigationSignedState{nc.t})
	e.PreventDefault()
}

func (r NavigationSignedDef) Render() react.Element {
	return react.Nav(&react.NavProps{ClassName: "navbar navbar-default"},
		react.Div(&react.DivProps{ClassName: "container"},
			react.Div(&react.DivProps{ClassName: "navbar-header"},
				react.A(&react.AProps{ClassName: "navbar-brand", Href: "/"},
					react.S(r.Props().BrandName),
				),
			),
			react.Ul(&react.UlProps{ClassName: "nav navbar-nav"},
				r.buildLink("Companies", navCompanies, navigationTabSignedChange{r, navCompanies}),
			),
		),
	)
}

func (n NavigationSignedDef) buildLink(prompt string, nt navigationTabSigned, nc navigationTabSignedChange) *react.LiElem {
	var lip *react.LiProps

	if n.State().activeNavTab == nt {
		lip = &react.LiProps{ClassName: "active"}
	}

	return react.Li(lip,
		react.A(&react.AProps{Href: "/index"}, react.S(prompt)),
	)
}
