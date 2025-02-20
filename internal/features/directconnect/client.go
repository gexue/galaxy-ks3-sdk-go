//Package directconnect provides gucumber integration tests suppport.
package directconnect

import (
	"github.com/gexue/galaxy-ks3-sdk-go/internal/features/shared"
	"github.com/gexue/galaxy-ks3-sdk-go/service/directconnect"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@directconnect", func() {
		World["client"] = directconnect.New(nil)
	})
}
