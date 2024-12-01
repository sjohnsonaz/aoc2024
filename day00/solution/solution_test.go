package solution

import (
	"log"
	"os"
	"testing"
)

func TestListAnalysis(t *testing.T) {
	file, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}

	la, err := NewListAnalysis(file)
	if err != nil {
		t.Fatal(err)
	}

	if la.Distance != 11 {
		t.Errorf("Distance is incorrect.  Expected: %v, Recieved: %v", 11, la.Distance)
	}

	if la.Similarity != 31 {
		t.Errorf("Similarity is incorrect.  Expected: %v, Recieved: %v", 31, la.Similarity)
	}
}
