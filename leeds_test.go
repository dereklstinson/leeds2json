package leeds2json_test

import (
	"fmt"
	"testing"

	"github.com/dereklstinson/leeds2json"
)

func TestExtended(t *testing.T) {
	slice, err := leeds2json.GetLeedsExtended()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(slice)
}
