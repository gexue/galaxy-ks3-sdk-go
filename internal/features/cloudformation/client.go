//Package cloudformation provides gucumber integration tests suppport.
package cloudformation

import (
	"github.com/gexue/galaxy-ks3-sdk-go/internal/features/shared"
	"github.com/gexue/galaxy-ks3-sdk-go/service/cloudformation"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@cloudformation", func() {
		World["client"] = cloudformation.New(nil)
	})
}
