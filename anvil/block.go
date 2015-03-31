// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package anvil

import "github.com/jteeuwen/mctools/anvil/item"

// Block represents a single block.
type Block struct {
	Id         item.Id // The complete block id.
	BlockLight uint8   // Amount of block-emitted light.
	SkyLight   uint8   // Amount of sunlight or moonlight hitting the block.
}
