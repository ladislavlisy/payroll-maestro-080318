package handlers

import (
	"context"

	"go.isomorphicgo.org/go/isokit"
	"honnef.co/go/js/dom"
	"myitcv.io/react"

	"github.com/ladislavlisy/payroll-maestro-080318/client/isoclient/common"
)

var document = dom.GetWindow().Document()

func IndexHandler(env *common.Env) isokit.Handler {
	return isokit.HandlerFunc(func(ctx context.Context) {
		messageTarget := document.GetElementByID("container")
		react.Render(react.H1(nil, react.S("Index")), messageTarget)
	})
}
