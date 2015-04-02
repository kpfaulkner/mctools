// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package recipe

import (
	"fmt"
	"testing"

	"github.com/jteeuwen/mctools/anvil/item"
)

func TestRecipe(t *testing.T) {
	fmt.Println(len(recipes))

	set := Find(item.AndesitePolished)
	fmt.Println(set)
}
