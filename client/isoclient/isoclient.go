// isoclient.go

package isoclient

import (
	"github.com/ladislavlisy/payroll-maestro-080318/client/isoclient/common"
	"github.com/ladislavlisy/payroll-maestro-080318/client/isoclient/handlers"

	"go.isomorphicgo.org/go/isokit"
)

func RegisterRoutes(env *common.Env) {
	r := isokit.NewRouter()
	r.Handle("/index", handlers.IndexHandler(env))
	r.Listen()
	env.Router = r
}
