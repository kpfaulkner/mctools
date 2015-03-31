// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

//go:generate stringer -type=TagId

package nbt

// tagId describes a type oftag.
type tagId uint8

// Known tag types
const (
	tagEnd       tagId = 0x0
	tagByte      tagId = 0x1
	tagShort     tagId = 0x2
	tagInt       tagId = 0x3
	tagLong      tagId = 0x4
	tagFloat     tagId = 0x5
	tagDouble    tagId = 0x6
	tagByteArray tagId = 0x7
	tagString    tagId = 0x8
	tagList      tagId = 0x9
	tagCompound  tagId = 0xa
	tagIntArray  tagId = 0xb
	tagUnknown   tagId = 0xc
)
