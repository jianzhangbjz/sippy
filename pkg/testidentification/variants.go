package testidentification

import (
	"github.com/openshift/sippy/pkg/db/models"
	"github.com/openshift/sippy/pkg/util/sets"
)

// VariantManager identifies and describes different variants
type VariantManager interface {
	// AllVariants returns a set of all known variants
	AllVariants() sets.String

	// AllPlatforms returns a set of all known platform variants
	AllPlatforms() sets.String

	// IdentifyVariants takes a job name and returns the list of variants that job belongs to.
	IdentifyVariants(jobName string, release string, jobVariants models.ClusterData) []string

	// IsJobNeverStable returns true if the job has been curated as never having passed more than 50ish% of the time.
	// This is used sparingly for jobs that are persistently failing and never taken stable.
	IsJobNeverStable(jobName string) bool
}
