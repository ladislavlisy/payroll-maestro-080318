//UserSignedCentral.go

package app

import (
	"github.com/ladislavlisy/payroll-maestro-080318/client/components"

	"myitcv.io/react"
)

//go:generate reactGen

type UserSignedCentralDef struct {
	react.ComponentDef
}

type UserSignedCentralProps struct {
	BrandName string
	Title     string
	Handler   AppSignHandler
}

func UserSignedCentral(p UserSignedCentralProps) *UserSignedCentralElem {
	return buildUserSignedCentralElem(p)
}

func (r UserSignedCentralDef) Render() react.Element {
	return react.Div(nil,
		components.NavigationSigned(components.NavigationSignedProps{BrandName: r.Props().BrandName}),
		react.Div(&react.DivProps{ClassName: "container", ID: "container"},
			react.H1(nil,
				react.S("Welcome to "+r.Props().BrandName),
			),
			react.P(nil,
				react.S("This is central page."),
			),
			ClickButton(ClickButtonProps{Title: r.Props().Title, Handler: r.Props().Handler, HandlerFlg: false}),
		),
	)
}
