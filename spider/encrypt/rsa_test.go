package encrypt

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	fmt.Println(len(RandomStr(16)))
}
