package solution

import (
	"log"
	"os"
	"testing"
)

func TestSafeData(t *testing.T) {
	file, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}

	sd, err := NewSafeData(file)
	if err != nil {
		log.Fatal(err)
	}

	if sd.Safe != 2 {
		t.Errorf("Safe count is incorrect.  Expected: %v, Recieved: %v", 2, sd.Safe)
	}

	if sd.SafeDampened != 4 {
		t.Errorf("Safe count is incorrect.  Expected: %v, Recieved: %v", 4, sd.SafeDampened)
	}
}

func TestReportSafe(t *testing.T) {
	for _, test := range []struct {
		r    Report
		safe bool
	}{{
		r:    Report{1, 2, 3},
		safe: true,
	}, {
		r:    Report{3, 2, 1},
		safe: true,
	}, {
		r:    Report{3, 1, 2},
		safe: false,
	}, {
		r:    Report{1, 1, 2},
		safe: false,
	}, {
		r:    Report{1, 2, 6},
		safe: false,
	}, {
		r:    Report{6, 5, 1},
		safe: false,
	}} {
		t.Run("", func(t *testing.T) {
			if s, _ := test.r.safe(); s != test.safe {
				t.Error("incorrect")
			}
		})
	}
}
