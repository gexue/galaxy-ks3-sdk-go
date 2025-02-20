//Package workspaces provides gucumber integration tests suppport.
package workspaces

import (
	"github.com/gexue/galaxy-ks3-sdk-go/internal/features/shared"
	"github.com/gexue/galaxy-ks3-sdk-go/service/workspaces"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@workspaces", func() {
		World["client"] = workspaces.New(nil)
	})
}
