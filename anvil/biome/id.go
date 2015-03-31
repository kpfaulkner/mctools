// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

//go:generate stringer -type=Id

package biome

// ref: http://minecraft.gamepedia.com/Biome

// Id defines a biome id.
type Id uint8

// Known biomes.
const (
	Ocean               Id = 0
	Plains              Id = 1
	SunflowerPlains     Id = 129
	Desert              Id = 2
	DesertM             Id = 130
	ExtremeHills        Id = 3
	ExtremeHillsM       Id = 131
	Forest              Id = 4
	FlowerForest        Id = 132
	Taiga               Id = 5
	TaigaM              Id = 133
	Swampland           Id = 6
	SwamplandM          Id = 134
	River               Id = 7
	Hell                Id = 8
	TheEnd              Id = 9
	FrozenOcean         Id = 10
	FrozenRiver         Id = 11
	IcePlains           Id = 12
	IcePlainsSpikes     Id = 140
	IceMountains        Id = 13
	MushroomIsland      Id = 14
	MushroomIslandShore Id = 15
	Beach               Id = 16
	DesertHills         Id = 17
	ForestHills         Id = 18
	TaigaHills          Id = 19
	ExtremeHillsEdge    Id = 20
	Jungle              Id = 21
	JungleM             Id = 149
	JungleHills         Id = 22
	JungleEdge          Id = 23
	JungleEdgeM         Id = 151
	DeepOcean           Id = 24
	StoneBeach          Id = 25
	ColdBeach           Id = 26
	BirchForest         Id = 27
	BirchForestM        Id = 155
	BirchForestHills    Id = 28
	BirchForestHillsM   Id = 156
	RoofedForest        Id = 29
	RoofedForestM       Id = 157
	ColdTaiga           Id = 30
	ColdTaigaM          Id = 158
	ColdTaigaHills      Id = 31
	MegaTaiga           Id = 32
	MegaSpruceTaiga     Id = 160
	MegaTaigaHills      Id = 33
	RedwoodTaigaHillsM  Id = 161
	ExtremeHillsPlus    Id = 34
	ExtremeHillsPlusM   Id = 162
	Savanna             Id = 35
	SavannaM            Id = 163
	SavannaPlateau      Id = 36
	SavannaPlateauM     Id = 164
	Mesa                Id = 37
	MesaBryce           Id = 165
	MesaPlateauF        Id = 38
	MesaPlateauFM       Id = 166
	MesaPlateau         Id = 39
	MesaPlateauM        Id = 167
)
