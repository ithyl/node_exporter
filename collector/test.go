package collector

import (
	"github.com/prometheus/procfs"
	"testing"
)

var proc = "D:\\work\\source code\\monitor\\node_exporter-master\\node_exporter-master\\collector\\fixtures\\proc"

func test1(t *testing.T) {
	_, err := procfs.NewFS(proc)
	if err != nil {
	}
}
