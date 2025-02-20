//Package elastictranscoder provides gucumber integration tests suppport.
package elastictranscoder

import (
	"github.com/gexue/galaxy-ks3-sdk-go/internal/features/shared"
	"github.com/gexue/galaxy-ks3-sdk-go/service/elastictranscoder"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@elastictranscoder", func() {
		World["client"] = elastictranscoder.New(nil)
	})
}
