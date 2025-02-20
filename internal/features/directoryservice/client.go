//Package directoryservice provides gucumber integration tests suppport.
package directoryservice

import (
	"github.com/gexue/galaxy-ks3-sdk-go/internal/features/shared"
	"github.com/gexue/galaxy-ks3-sdk-go/service/directoryservice"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@directoryservice", func() {
		// FIXME remove custom region
		World["client"] = directoryservice.New(nil)
	})
}
