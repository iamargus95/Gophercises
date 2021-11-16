package kindergarten

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Garden map[string][]string

func NewGarden(diagram string, children []string) (*Garden, error) {
	if children == nil {

		return nil, errors.New("no children")
	}

	orderedChildren := make([]string, len(children))
	copy(orderedChildren, children)
	sort.Strings(orderedChildren)
	garden := make(Garden)
	plantMap := map[rune]string{
		'V': "violets",
		'R': "radishes",
		'C': "clover",
		'G': "grass",
	}

	rows := strings.Split(diagram, "\n")
	if len(rows) != 3 {

		return nil, fmt.Errorf("expected 3 lines in diagram bug got %v", len(rows))
	}

	row0 := []rune(rows[1])
	row1 := []rune(rows[2])
	if len(row0) != 2*len(orderedChildren) || len(row1) != 2*len(orderedChildren) {

		return nil, errors.New("number of plants in row is illegal")
	}

	for i, child := range orderedChildren {
		plant0, ok0 := plantMap[row0[2*i]]
		plant1, ok1 := plantMap[row0[2*i+1]]
		plant2, ok2 := plantMap[row1[2*i]]
		plant3, ok3 := plantMap[row1[2*i+1]]

		if !(ok0 && ok1 && ok2 && ok3) {

			return nil, errors.New("one or more cups have invalid codes")
		}

		plants := []string{plant0, plant1, plant2, plant3}
		garden[child] = plants
	}

	if len(garden) != len(children) {

		return nil, errors.New("duplicate name found in children names")
	}

	return &garden, nil
}
func (g *Garden) Plants(child string) (plants []string, ok bool) {
	plants, ok = (*g)[child]
	return
}
