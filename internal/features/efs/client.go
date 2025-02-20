//Package efs provides gucumber integration tests suppport.
package efs

import (
	"github.com/gexue/galaxy-ks3-sdk-go/aws"
	"github.com/gexue/galaxy-ks3-sdk-go/internal/features/shared"
	"github.com/gexue/galaxy-ks3-sdk-go/service/efs"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@efs", func() {
		// FIXME remove custom region
		World["client"] = efs.New(&aws.Config{Region: "us-west-2"})
	})
}
