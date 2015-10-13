package collector

import (
	"github.com/lowstz/mongodb_exporter/shared"
	"testing"
)

func Test_NetworkCollectData(t *testing.T) {
	stats := &NetworkStats{}

	groupName := "network"
	stats.Export(groupName)

	if shared.Groups[groupName+"_bytes_total"] == nil {
		t.Error("Group not created")
	}
	if shared.Groups[groupName+"_metrics"] == nil {
		t.Error("Group not created")
	}
}
