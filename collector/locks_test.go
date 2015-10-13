package collector

import (
	"github.com/lowstz/mongodb_exporter/shared"
	"testing"
)

func Test_LocksCollectData(t *testing.T) {
	stats := &LockStatsMap{
		".": LockStats{
			TimeLockedMicros:    ReadWriteLockTimes{},
			TimeAcquiringMicros: ReadWriteLockTimes{},
		},
	}

	groupName := "locks"
	stats.Export(groupName)

	if shared.Groups["locks_time_locked_microseconds_global"] == nil {
		t.Error("Group not created")
	}
	if shared.Groups["locks_time_acquiring_microseconds_global"] == nil {
		t.Error("Group not created")
	}
}
