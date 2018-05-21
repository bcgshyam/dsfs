// +build linux

package svtfs

import (
	"testing"

	"github.com/docker/docker/daemon/graphdriver/graphtest"

	"github.com/docker/docker/pkg/reexec"

)

func init() {
	reexec.Init()
}

// This avoids creating a new driver for each test if all tests are run
// Make sure to put new tests between TestVfsSetup and TestVfsTeardown
func TestsvtfsSetup(t *testing.T) {
	graphtest.GetDriver(t, "svtfs")
}

func TestsvtfsCreateEmpty(t *testing.T) {
	graphtest.DriverTestCreateEmpty(t, "svtfs")
}

func TestsvtfsCreateBase(t *testing.T) {
	graphtest.DriverTestCreateBase(t, "svtfs")
}

func TestsvtfsCreateSnap(t *testing.T) {
	graphtest.DriverTestCreateSnap(t, "svtfs")
}

func TestsvtfsSetQuota(t *testing.T) {
	graphtest.DriverTestSetQuota(t, "svtfs", false)
}

func TestsvtfsTeardown(t *testing.T) {
	graphtest.PutDriver(t)
}
