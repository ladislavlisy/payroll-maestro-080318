//NavigationSignIn.go

package components

import (
	"myitcv.io/react"
)

type navigationTabSignIn uint32

//go:generate reactGen

type NavigationSignInDef struct {
	react.ComponentDef
}

type NavigationSignInProps struct {
	BrandName string
}

type NavigationSignInState struct {
	activeNavTab navigationTabSignIn
}

func NavigationSignIn(p NavigationSignInProps) *NavigationSignInElem {
	return buildNavigationSignInElem(p)
}

const (
	navExplore navigationTabSignIn = iota
)

type navigationTabSignInChange struct {
	n NavigationSignInDef
	t navigationTabSignIn
}

func (nc navigationTabSignInChange) OnClick(e *react.SyntheticMouseEvent) {
	nc.n.SetState(NavigationSignInState{nc.t})
	e.PreventDefault()
}

func (r NavigationSignInDef) Render() react.Element {
	return react.Nav(&react.NavProps{ClassName: "navbar navbar-default"},
		react.Div(&react.DivProps{ClassName: "container"},
			react.Div(&react.DivProps{ClassName: "navbar-header"},
				react.A(&react.AProps{ClassName: "navbar-brand", Href: "/"},
					react.S(r.Props().BrandName),
				),
			),
			react.Ul(&react.UlProps{ClassName: "nav navbar-nav"},
				r.buildLink("Explore", navExplore, navigationTabSignInChange{r, navExplore}),
			),
		),
	)
}

func (n NavigationSignInDef) buildLink(prompt string, nt navigationTabSignIn, nc navigationTabSignInChange) *react.LiElem {
	var lip *react.LiProps

	if n.State().activeNavTab == nt {
		lip = &react.LiProps{ClassName: "active"}
	}

	return react.Li(lip,
		react.A(&react.AProps{Href: "/index"}, react.S(prompt)),
	)
}
