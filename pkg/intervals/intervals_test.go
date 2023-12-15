package intervals

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

const test = "[400 500 ] 505   [111 450] (505 600] 505 400 200 [502 503)"

const testClean = "[111 500] (505 600] [502 503) 505"

var testTokens = []string{"[", "400", "500", "]", "505", "[", "111", "450", "]", "(", "505", "600", "]", "505", "400", "200", "[", "502", "503", ")"}

func TestSplitTokens(t *testing.T) {
	reader := strings.NewReader(test)

	scanner := bufio.NewScanner(reader)

	scanner.Split(splitTokens)

	assert := func(i int, exp string, act string) {
		if exp != act {
			t.Errorf("expected %s got %s", exp, act)
		}
	}

	for i := 0; scanner.Scan(); i++ {
		assert(i, testTokens[i], scanner.Text())
	}
}

func TestStingRepresentation(t *testing.T) {
	r, err := NewInterval(test)

	if err != nil {
		t.Error(err)
	}

	assert := func(exp string, act string) {
		if exp != act {
			t.Errorf("expected %s got %s", exp, act)
		}
	}

	assert(testClean, fmt.Sprint(r))
}

func TestErr(t *testing.T) {
	_, err := NewInterval("[300 500] abc [300 400]")

	t.Logf("%v", err)

	if err == nil {
		t.Errorf("no err")
	}

	_, err = NewInterval("[300 500] abc [300 ]")

	t.Logf("%v", err)

	if err == nil {
		t.Errorf("no err")
	}
}

func TestInterval_Includes(t *testing.T) {
	r, _ := NewInterval(test)

	if !r.Includes(200) {
		t.Errorf("should include")
	}

	if !r.Includes(505) {
		t.Errorf("should include")
	}

	if r.Includes(601) {
		t.Errorf("should not include")
	}
}
