// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package recipe

import (
	"bytes"
	"fmt"
)

// RecipeList defines a list of recipes, sortable by ingredient count.
type RecipeList []*Recipe

func (r RecipeList) Len() int           { return len(r) }
func (r RecipeList) Less(i, j int) bool { return len(r[i].Ingredients) < len(r[j].Ingredients) }
func (r RecipeList) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func (r RecipeList) String() string {
	var w bytes.Buffer

	for _, v := range r {
		fmt.Fprint(&w, v.String())
	}

	return w.String()
}
