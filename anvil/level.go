// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

//go:generate stringer -type=GameMode,Difficulty -output=level_string.go

package anvil

import (
	"compress/gzip"
	"os"

	"github.com/kpfaulkner/mctools/anvil/nbt"
)

// GameMode defines the current game mode.
type GameMode int8

// Known game modes.
const (
	Survival GameMode = iota
	Creative
	Adventure
	Spectator
)

// Difficulty defines the world's difficulty level
type Difficulty int8

// Known game modes.
const (
	Peaceful Difficulty = iota
	Easy
	Normal
	Hard
)

// GameRules describes the current rules for a world.
type GameRules struct {
	RandomTickSpeed     string `nbt:"randomTickSpeed"`
	CommandBlockOutput  bool   `nbt:"commandBlockOutput"`
	DaylightCycle       bool   `nbt:"doDaylightCycle"`
	FireTick            bool   `nbt:"doFireTick"`
	TileDrops           bool   `nbt:"doTileDrops"`
	KeepInventory       bool   `nbt:"keepInventory"`
	LogAdminCommands    bool   `nbt:"logAdminCommands"`
	MobLoot             bool   `nbt:"doMobLoot"`
	MobSpawning         bool   `nbt:"doMobSpawning"`
	MobGriefing         bool   `nbt:"mobGriefing"`
	NaturalRegeneration bool   `nbt:"naturalRegeneration"`
	SendCommandFeedback bool   `nbt:"sendCommandFeedback"`
	ShowDeathMessages   bool   `nbt:"showDeathMessages"`
	ReducedDebugInfo    bool   `nbt:"reducedDebugInfo"`
	EntityDrops         bool   `nbt:"doEntityDrops"`
}

// Level describes the level.dat file for a Minecraft world.
// It holds general information about a world, like the name,
// the generator and seed and other things.
type Level struct {
	Player               *Player    `nbt:"Player"`
	Rules                GameRules  `nbt:"GameRules"`
	Name                 string     `nbt:"LevelName"`
	GeneratorName        string     `nbt:"generatorName"`
	GeneratorOptions     string     `nbt:"generatorOptions"`
	LastPlayed           int64      `nbt:"LastPlayed"`
	Seed                 int64      `nbt:"RandomSeed"`
	Time                 int64      `nbt:"Time"`
	DayTime              int64      `nbt:"DayTime"`
	SizeOnDisk           int64      `nbt:"SizeOnDisk"`
	BorderSizeLerpTime   int64      `nbt:"BorderSizeLerpTime"`
	BorderCenterX        float64    `nbt:"BorderCenterX"`
	BorderCenterZ        float64    `nbt:"BorderCenterZ"`
	BorderSize           float64    `nbt:"BorderSize"`
	BorderSizeLerpTarget float64    `nbt:"BorderSizeLerpTarget"`
	BorderWarningBlocks  float64    `nbt:"BorderWarningBlocks"`
	BorderWarningTime    float64    `nbt:"BorderWarningTime"`
	BorderDamagePerBlock float64    `nbt:"BorderDamagePerBlock"`
	BorderSafeZone       float64    `nbt:"BorderSafeZone"`
	GeneratorVersion     int32      `nbt:"generatorVersion"`
	Version              int32      `nbt:"version"`
	SpawnX               int32      `nbt:"SpawnX"`
	SpawnY               int32      `nbt:"SpawnY"`
	SpawnZ               int32      `nbt:"SpawnZ"`
	RainTime             int32      `nbt:"rainTime"`
	ClearWeatherTime     int32      `nbt:"clearWeatherTime"`
	ThunderTime          int32      `nbt:"thunderTime"`
	GameMode             GameMode   `nbt:"GameType"`
	Difficulty           Difficulty `nbt:"Difficulty"`
	Initialized          bool       `nbt:"initialized"`
	MapFeatures          bool       `nbt:"MapFeatures"`
	AllowCommands        bool       `nbt:"allowCommands"`
	Hardcore             bool       `nbt:"hardcore"`
	DifficultyLocked     bool       `nbt:"DifficultyLocked"`
	Raining              bool       `nbt:"raining"`
	Thundering           bool       `nbt:"thundering"`
}

// LoadLevel loads level data from the given level.dat file.
func LoadLevel(file string) (*Level, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer fd.Close()

	gz, err := gzip.NewReader(fd)
	if err != nil {
		return nil, err
	}

	defer gz.Close()

	var v struct {
		Data Level
	}

	err = nbt.Unmarshal(gz, &v)
	if err != nil {
		return nil, err
	}

	return &v.Data, nil
}

// Save saves level data to the given file.
func (l *Level) Save(file string) error {
	fd, err := os.Create(file)
	if err != nil {
		return err
	}

	defer fd.Close()

	var v struct {
		Data *Level
	}

	v.Data = l

	gz := gzip.NewWriter(fd)
	err = nbt.Marshal(gz, v)
	gz.Close()
	return err
}
