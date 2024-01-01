package day1

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var data = []item{
	{"mxmldpfsevenpfcvhff9twonineeight", 78},
}

type item struct {
	input  string
	result int
}

func TestDay2(t *testing.T) {
	for _, i := range data {
		r, err := getCalibrationValuePart2V1(i.input)
		if assert.NoError(t, err) {
			assert.Equal(t, i.result, r)
		}
	}
}

func TestDay2FullInput(t *testing.T) {
	fp, err := os.Open("../data/day1.txt")
	assert.NoError(t, err)
	defer fp.Close()
	reader := bufio.NewScanner(fp)

	for reader.Scan() {
		line := reader.Text()
		v1, err := getCalibrationValuePart2V1(line)
		assert.NoError(t, err)
		v2, err := getCalibrationValuePart2V2(line)
		assert.NoError(t, err)
		assert.Equal(t, v2, v1, line)
	}
	assert.NoError(t, reader.Err())
}
