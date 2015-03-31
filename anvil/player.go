// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package anvil

// Player defines all properties for a single player.
// For single-player games, this is part of level.dat.
// For servers, this is stored in separate files in the $WORLD/playerdata/ directory.
type Player struct {
	Abilities           Abilities       `nbt:"abilities"`
	Attributes          []Attribute     `nbt:"Attributes"`
	Inventory           []InventorySlot `nbt:"Inventory"`
	EnderItems          []InventorySlot `nbt:"EnderItems"`
	Motion              []float64       `nbt:"Motion"`
	Pos                 []float64       `nbt:"Pos"`
	Rotation            []float64       `nbt:"Rotation"`
	UUIDLeast           int64           `nbt:"UUIDLeast"`
	UUIDMost            int64           `nbt:"UUIDMost"`
	FoodExhaustionLevel float32         `nbt:"foodExhaustionLevel"`
	HealF               float32         `nbt:"HealF"`
	XpP                 float32         `nbt:"XpP"`
	AbsorptionAmount    float32         `nbt:"AbsorptionAmount"`
	FoodSaturationLevel float32         `nbt:"foodSaturationLevel"`
	FallDistance        float32         `nbt:"FallDistance"`
	SelectedItemSlot    int32           `nbt:"SelectedItemSlot"`
	FoodTickTimer       int32           `nbt:"foodTickTimer"`
	XpLevel             int32           `nbt:"XpLevel"`
	XpSeed              int32           `nbt:"XpSeed"`
	XpTotal             int32           `nbt:"XpTotal"`
	PlayerGameType      int32           `nbt:"playerGameType"`
	FoodLevel           int32           `nbt:"foodLevel"`
	Score               int32           `nbt:"Score"`
	HurtByTimestamp     int32           `nbt:"HurtByTimestamp"`
	Dimension           int32           `nbt:"Dimension"`
	PortalCooldown      int32           `nbt:"PortalCooldown"`
	Health              int16           `nbt:"Health"`
	Fire                int16           `nbt:"Fire"`
	DeathTime           int16           `nbt:"DeathTime"`
	SleepTimer          int16           `nbt:"SleepTimer"`
	HurtTime            int16           `nbt:"HurtTime"`
	Air                 int16           `nbt:"Air"`
	Sleeping            bool            `nbt:"Sleeping"`
	Invulnerable        bool            `nbt:"Invulnerable"`
	OnGround            bool            `nbt:"OnGround"`
}
