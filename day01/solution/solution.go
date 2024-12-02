package solution

import (
	"bufio"
	"encoding/csv"
	"io"
	"slices"
	"strconv"
	"strings"
)

type SafeData struct {
	Safe         int
	SafeDampened int
}

func NewSafeData(r io.Reader) (SafeData, error) {
	scanner := bufio.NewScanner(r)
	reader := csv.NewReader(r)
	reader.Comma = ' '
	safe := 0
	safeDampened := 0
	for scanner.Scan() {
		record := strings.Fields(scanner.Text())
		report, err := newReport(record)
		if err != nil {
			return SafeData{}, err
		}

		if s, i := report.safe(); s {
			safe++
			safeDampened++
		} else if report.safeDampened(i) {
			safeDampened++
		}
	}

	return SafeData{
		Safe:         safe,
		SafeDampened: safeDampened,
	}, nil
}

func newReport(record []string) (Report, error) {
	data := make([]int, 0, len(record))
	for _, a := range record {
		x, err := strconv.Atoi(a)
		if err != nil {
			return Report{}, err
		}

		data = append(data, x)
	}
	return Report(data), nil
}

type Report []int

func (r Report) safe() (bool, int) {
	switch len(r) {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		return true, 0
	default:
		direction := 0
		for i := 1; i < len(r); i++ {
			a, b := r[i-1], r[i]
			c := b - a
			if c == 0 {
				return false, i - 1
			} else if c > 0 {
				if c > 3 {
					return false, i - 1
				} else if direction == 0 {
					direction = 1
				} else if direction == -1 {
					return false, i - 1
				}
			} else {
				if c < -3 {
					return false, i - 1
				} else if direction == 0 {
					direction = -1
				} else if direction == 1 {
					return false, i - 1
				}
			}
		}
		return true, 0
	}
}

func (r Report) safeDampened(i int) bool {
	r = slices.Delete(r, i, i+1)
	safe, _ := r.safe()
	return safe
}
