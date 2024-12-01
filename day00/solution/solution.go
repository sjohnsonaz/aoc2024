package solution

import (
	"encoding/csv"
	"errors"
	"io"
	"slices"
	"strconv"
)

type ListAnalysis struct {
	Distance   int
	Similarity int
}

func NewListAnalysis(r io.Reader) (ListAnalysis, error) {
	a, b, err := loadData(r)
	if err != nil {
		return ListAnalysis{}, err
	}

	return ListAnalysis{
		Distance:   listDistance(a, b),
		Similarity: listSimilarity(a, b),
	}, nil
}

func loadData(r io.Reader) ([]int, []int, error) {
	reader := csv.NewReader(r)
	reader.Comma = ' '

	a := make([]int, 0)
	b := make([]int, 0)

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, nil, err
			}
		}

		if len(record) < 4 {
			return nil, nil, errors.New("incorrect record length")
		}

		x, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, nil, err
		}

		y, err := strconv.Atoi(record[3])
		if err != nil {
			return nil, nil, err
		}

		a = append(a, x)
		b = append(b, y)
	}

	slices.Sort(a)
	slices.Sort(b)

	return a, b, nil
}

func listDistance(a []int, b []int) int {
	d := 0

	for i := range a {
		x := a[i]
		y := b[i]
		d += distance(x, y)
	}

	return d
}

func distance(a int, b int) int {
	c := b - a
	if c < 0 {
		return -c
	}
	return c
}

func listSimilarity(a []int, b []int) int {
	s := 0

	for _, x := range a {
		for _, y := range b {
			if x == y {
				s += x
			}
		}
	}

	return s
}
