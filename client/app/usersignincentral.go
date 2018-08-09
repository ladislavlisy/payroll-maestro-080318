//UserSignInCentral.go

package app

import (
	"github.com/ladislavlisy/payroll-maestro-080318/client/components"

	"myitcv.io/react"
)

//go:generate reactGen

type UserSignInCentralDef struct {
	react.ComponentDef
}

type UserSignInCentralProps struct {
	BrandName string
	Title     string
	SignInApp AppDef
}

func UserSignInCentral(p UserSignInCentralProps) *UserSignInCentralElem {
	return buildUserSignInCentralElem(p)
}

func (r UserSignInCentralDef) Render() react.Element {
	return react.Div(nil,
		components.NavigationSignIn(components.NavigationSignInProps{BrandName: r.Props().BrandName}),
		react.Div(&react.DivProps{ClassName: "container", ID: "container"},
			react.H1(nil,
				react.S("Welcome to "+r.Props().BrandName),
			),
			react.P(nil,
				react.S("This is sign in page."),
			),
			ClickButton(ClickButtonProps{OnClickDef: r.Props().SignInApp, OnClickFlg: true, Title: r.Props().Title}),
		),
	)
}
