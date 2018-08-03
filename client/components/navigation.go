//Navigation.go

package components

import (
	"myitcv.io/react"
)

type navTab uint32

//go:generate reactGen

type NavigationDef struct {
	react.ComponentDef
}

type NavigationProps struct {
	BrandName string
}

type NavigationState struct {
	activeNavTab navTab
}

func Navigation(p NavigationProps) *NavigationElem {
	return buildNavigationElem(p)
}

const (
	navCompanies navTab = iota
)

type navTabChange struct {
	n NavigationDef
	t navTab
}

func (nc navTabChange) OnClick(e *react.SyntheticMouseEvent) {
	nc.n.SetState(NavigationState{nc.t})
	e.PreventDefault()
}

func (r NavigationDef) Render() react.Element {
	return react.Nav(&react.NavProps{ClassName: "navbar navbar-default"},
		react.Div(&react.DivProps{ClassName: "container"},
			react.Div(&react.DivProps{ClassName: "navbar-header"},
				react.A(&react.AProps{ClassName: "navbar-brand", Href: "/"},
					react.S(r.Props().BrandName),
				),
			),
			react.Ul(&react.UlProps{ClassName: "nav navbar-nav"},
				r.buildLink("Companies", navCompanies, navTabChange{r, navCompanies}),
			),
		),
	)
}

func (n NavigationDef) buildLink(prompt string, nt navTab, nc navTabChange) *react.LiElem {
	var lip *react.LiProps

	if n.State().activeNavTab == nt {
		lip = &react.LiProps{ClassName: "active"}
	}

	return react.Li(lip,
		react.A(&react.AProps{Href: "/index"}, react.S(prompt)),
	)
}
