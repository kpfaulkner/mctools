// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package anvil

import (
	"github.com/jteeuwen/mctools/anvil/item"
)

// Modifier defines an attribute modifier.
type Modifier struct {
	Name      string  `nbt:"Name"`
	UUIDLeast int64   `nbt:"UUIDLeast"`
	UUIDMost  int64   `nbt:"UUIDMost"`
	Amount    float64 `nbt:"Amount"`
	Operation int32   `nbt:"Operation"`
}

// Attribute defines an active attribute for an entity.
type Attribute struct {
	Modifiers []Modifier `nbt:"Modifiers"`
	Name      string     `nbt:"Name"`
	Base      float64    `nbt:"Base"`
}

// InventorySlot defines an inventory slot.
type InventorySlot struct {
	Id     string `nbt:"id"`
	Damage int16  `nbt:"Damage"`
	Count  int8   `nbt:"Count"`
	Slot   int8   `nbt:"Slot"`
}

// Abilities describes entity abilities.
type Abilities struct {
	FlySpeed     float32 `nbt:"flySpeed"`
	WalkSpeed    float32 `nbt:"walkSpeed"`
	Flying       bool    `nbt:"flying"`
	Instabuild   bool    `nbt:"instabuild"`
	Invulnerable bool    `nbt:"invulnerable"`
	Mayfly       bool    `nbt:"mayfly"`
	MayBuild     bool    `nbt:"mayBuild"`
}

type EntityTag struct {
	CanDestroy  item.Id `nbt:"CanDestroy"`
	Unbreakable bool    `nbt:"Damage"`
}

// Items are used both in the player's inventory, Ender inventory,
// and in chest tile entities, dropped item entities, furnace tile
// entities, brewing stand tile entities, and Villager trading recipes.
//
// Sometimes a Slot tag is used to specify the slot the item is in,
// such was with chests; other times there is no Slot tag, such as with
// dropped items.
type Item struct {
	Tag    EntityTag `nbt:"tag"`
	Id     string    `nbt:"id"`
	Count  int8      `nbt:"Count"`
	Slot   int8      `nbt:"Slot"`
	Damage int8      `nbt:"Damage"`
}

// CommandStats defines information identifying scoreboard parameters
// to modify relative to the last command run.
type CommandStats struct {
	SuccessCountObjective     string `nbt:"SuccessCountObjective"`
	SuccessCountName          string `nbt:"SuccessCountName"`
	AffectedBlocksObjective   string `nbt:"AffectedBlocksObjective"`
	AffectedBlocksName        string `nbt:"AffectedBlocksName"`
	AffectedEntitiesObjective string `nbt:"AffectedEntitiesObjective"`
	AffectedEntitiesName      string `nbt:"AffectedEntitiesName"`
	AffectedItemsObjective    string `nbt:"AffectedItemsObjective"`
	AffectedItemsName         string `nbt:"AffectedItemsName"`
	QueryResultObjective      string `nbt:"QueryResultObjective"`
	QueryResultName           string `nbt:"QueryResultName"`
}

// TileEntity describes a tile entity.
//
// These are part of Chunk descriptors and cover things like mob spawners
// and chests.
type TileEntity struct {
	Id string `nbt:"id"`
	X  int32  `nbt:"x"`
	Y  int32  `nbt:"y"`
	Z  int32  `nbt:"z"`

	// Chest fields.
	Lock  string `nbt:"Lock"`
	Items []Item `nbt:"Items"`

	// Mob spawner fields.
	EntityId            string `nbt:"EntityId"`
	Delay               int16  `nbt:"Delay"`
	RequiredPlayerRange int16  `nbt:"RequiredPlayerRange"`
	MaxNearbyEntities   int16  `nbt:"MaxNearbyEntities"`
	MinSpawnDelay       int16  `nbt:"MinSpawnDelay"`
	MaxSpawnDelay       int16  `nbt:"MaxSpawnDelay"`
	SpawnRange          int16  `nbt:"SpawnRange"`
	SpawnCount          int16  `nbt:"SpawnCount"`
}

// Entity defines a single entity with fields shared by all entity types.
//
// FIXME: This listing is far from complete. There are a huge amount of
// fields we do not account for here as they belong to individual entity
// sub-types like Horses, chickens, etc.
type Entity struct {
	Riding            *Entity       `nbt:"Riding"`
	CommandStats      *CommandStats `nbt:"CommandStats"`
	Id                string        `nbt:"id"`
	UUID              string        `nbt:"UUID"`
	CustomName        string        `nbt:"CustomName"`
	OwnerUUID         string        `nbt:"OwnerUUID"`
	Owner             string        `nbt:"Owner"`
	Pos               []float64     `nbt:"Pos"`
	Motion            []float64     `nbt:"Motion"`
	Rotation          []float32     `nbt:"Rotation"`
	UUIDMost          int64         `nbt:"UUIDMost"`
	UUIDLeast         int64         `nbt:"UUIDLeast"`
	FallDistance      float32       `nbt:"FallDistance"`
	Dimension         int32         `nbt:"Dimension"`
	PortalCooldown    int32         `nbt:"PortalCooldown"`
	BreedingInLove    int32         `nbt:"InLove"`
	Fire              int16         `nbt:"Fire"`
	Air               int16         `nbt:"Air"`
	Health            uint8         `nbt:"Health"`
	OnGround          bool          `nbt:"OnGround"`
	Invulnerable      bool          `nbt:"Invulnerable"`
	CustomNameVisible bool          `nbt:"CustomNameVisible"`
	Silent            bool          `nbt:"Silent"`
}
