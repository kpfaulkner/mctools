// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

//go:generate stringer -type=Id

package item

// List of existing item IDs for version 1.8+
//
// Ref: http://www.minecraftinfo.com/idnamelist.htm
const (
	Air                        Id = 0
	Stone                      Id = 1
	Granite                    Id = 1 | 1<<16
	GranitePolished            Id = 1 | 2<<16
	Diorite                    Id = 1 | 3<<16
	DioritePolished            Id = 1 | 4<<16
	Andesite                   Id = 1 | 5<<16
	AndesitePolished           Id = 1 | 6<<16
	Grass                      Id = 2
	Dirt                       Id = 3
	CoarseDirt                 Id = 3 | 1<<16
	Podzol                     Id = 3 | 2<<16
	Cobblestone                Id = 4
	OakPlanks                  Id = 5
	SprucePlanks               Id = 5 | 1<<16
	BirchPlanks                Id = 5 | 2<<16
	JunglePlanks               Id = 5 | 3<<16
	AcaciaPlanks               Id = 5 | 4<<16
	DarkOakPlanks              Id = 5 | 5<<16
	OakSapling                 Id = 6
	SpruceSapling              Id = 6 | 1<<16
	BirchSapling               Id = 6 | 2<<16
	JungleSapling              Id = 6 | 3<<16
	AcaciaSapling              Id = 6 | 4<<16
	DarkOakSapling             Id = 6 | 5<<16
	Bedrock                    Id = 7
	WaterFlowing               Id = 8
	WaterNoSpread              Id = 9
	LavaFlowing                Id = 10
	LavaNoSpread               Id = 11
	Sand                       Id = 12
	RedSand                    Id = 12 | 1<<16
	Gravel                     Id = 13
	GoldOre                    Id = 14
	IronOre                    Id = 15
	CoalOre                    Id = 16
	OakLog                     Id = 17
	SpruceLog                  Id = 17 | 1<<16
	BirchLog                   Id = 17 | 2<<16
	JungleLog                  Id = 17 | 3<<16
	Oak4Log                    Id = 17 | 4<<16
	Oak5Log                    Id = 17 | 5<<16
	OakLeaves                  Id = 18
	SpruceLeaves               Id = 18 | 1<<16
	BirchLeaves                Id = 18 | 2<<16
	JungleLeaves               Id = 18 | 3<<16
	Sponge                     Id = 19
	SpongeWet                  Id = 19 | 1<<16
	Glass                      Id = 20
	LapisLazuliOre             Id = 21
	LapisLazuliBlock           Id = 22
	Dispenser                  Id = 23
	Sandstone                  Id = 24
	SandstoneChiseled          Id = 24 | 1<<16
	SandstoneSmooth            Id = 24 | 2<<16
	NoteBlock                  Id = 25
	BedBlock                   Id = 26
	RailPowered                Id = 27
	RailDetector               Id = 28
	StickyPiston               Id = 29
	Cobweb                     Id = 30
	TallGrassDeadShrub         Id = 31
	TallGrass                  Id = 31 | 1<<16
	TallGrassFern              Id = 31 | 2<<16
	DeadShrub                  Id = 32
	Piston                     Id = 33
	PistonHead                 Id = 34
	WhiteWool                  Id = 35
	OrangeWool                 Id = 35 | 1<<16
	MagentaWool                Id = 35 | 2<<16
	LightBlueWool              Id = 35 | 3<<16
	YellowWool                 Id = 35 | 4<<16
	LimeWool                   Id = 35 | 5<<16
	PinkWool                   Id = 35 | 6<<16
	GrayWool                   Id = 35 | 7<<16
	LightGrayWool              Id = 35 | 8<<16
	CyanWool                   Id = 35 | 9<<16
	PurpleWool                 Id = 35 | 10<<16
	BlueWool                   Id = 35 | 11<<16
	BrownWool                  Id = 35 | 12<<16
	GreenWool                  Id = 35 | 13<<16
	RedWool                    Id = 35 | 14<<16
	BlackWool                  Id = 35 | 15<<16
	PistonMoving               Id = 36
	Dandelion                  Id = 37
	Poppy                      Id = 38
	BlueOrchid                 Id = 38 | 1<<16
	Allium                     Id = 38 | 2<<16
	AzureBluet                 Id = 38 | 3<<16
	RedTulip                   Id = 38 | 4<<16
	OrangeTulip                Id = 38 | 5<<16
	WhiteTulip                 Id = 38 | 6<<16
	PinkTulip                  Id = 38 | 7<<16
	OxeyeDaisy                 Id = 38 | 8<<16
	MushroomBrown              Id = 39
	MushroomRed                Id = 40
	GoldBlock                  Id = 41
	IronBlock                  Id = 42
	StoneDoubleSlab            Id = 43
	SandstoneDoubleSlab        Id = 43 | 1<<16
	WoodDoubleSlab             Id = 43 | 2<<16
	CobblestoneDoubleSlab      Id = 43 | 3<<16
	BrickDoubleSlab            Id = 43 | 4<<16
	StoneBrickDoubleSlab       Id = 43 | 5<<16
	NetherBrickDoubleSlab      Id = 43 | 6<<16
	QuartzDoubleSlab           Id = 43 | 7<<16
	SmoothStoneDoubleSlab      Id = 43 | 8<<16
	SmoothSandstoneDoubleSlab  Id = 43 | 9<<16
	StoneSlab                  Id = 44
	SandstoneSlab              Id = 44 | 1<<16
	WoodSlab                   Id = 44 | 2<<16
	CobblestoneSlab            Id = 44 | 3<<16
	BrickSlab                  Id = 44 | 4<<16
	StoneBrickSlab             Id = 44 | 5<<16
	NetherBrickSlab            Id = 44 | 6<<16
	QuartzSlab                 Id = 44 | 7<<16
	Brick                      Id = 45
	TNT                        Id = 46
	Bookshelf                  Id = 47
	MossStone                  Id = 48
	Obsidian                   Id = 49
	Torch                      Id = 50
	Fire                       Id = 51
	MobSpawner                 Id = 52
	OakStairs                  Id = 53
	Chest                      Id = 54
	RedstoneWire               Id = 55
	DiamondOre                 Id = 56
	DiamondBlock               Id = 57
	Workbench                  Id = 58
	WheatCrop                  Id = 59
	Farmland                   Id = 60
	Furnace                    Id = 61
	FurnaceSmelting            Id = 62
	SignBlock                  Id = 63
	OakDoorBlock               Id = 64
	Ladder                     Id = 65
	Rail                       Id = 66
	CobblestoneStairs          Id = 67
	Signwall                   Id = 68
	Lever                      Id = 69
	StonePressurePlate         Id = 70
	IronDoorBlock              Id = 71
	WoodPressurePlate          Id = 72
	RedstoneOre                Id = 73
	RedstoneOreGlowing         Id = 74
	RedstoneTorchOff           Id = 75
	RedstoneTorch              Id = 76
	StoneButton                Id = 77
	Snow                       Id = 78
	Ice                        Id = 79
	SnowBlock                  Id = 80
	Cactus                     Id = 81
	ClayBlock                  Id = 82
	SugarcaneBlock             Id = 83
	Jukebox                    Id = 84
	OakFence                   Id = 85
	Pumpkin                    Id = 86
	Netherrack                 Id = 87
	SoulSand                   Id = 88
	Glowstone                  Id = 89
	Portal                     Id = 90
	JackOLantern               Id = 91
	CakeBlock                  Id = 92
	RedstoneRepeaterBlockOff   Id = 93
	RedstoneRepeaterBlockOn    Id = 94
	WhiteStainedGlass          Id = 95
	OrangeStainedGlass         Id = 95 | 1<<16
	MagentaStainedGlass        Id = 95 | 2<<16
	LightBlueStainedGlass      Id = 95 | 3<<16
	YellowStainedGlass         Id = 95 | 4<<16
	LimeStainedGlass           Id = 95 | 5<<16
	PinkStainedGlass           Id = 95 | 6<<16
	GrayStainedGlass           Id = 95 | 7<<16
	LightGrayStainedGlass      Id = 95 | 8<<16
	CyanStainedGlass           Id = 95 | 9<<16
	PurpleStainedGlass         Id = 95 | 10<<16
	BlueStainedGlass           Id = 95 | 11<<16
	BrownStainedGlass          Id = 95 | 12<<16
	GreenStainedGlass          Id = 95 | 13<<16
	RedStainedGlass            Id = 95 | 14<<16
	BlackStainedGlass          Id = 95 | 15<<16
	WoodTrapdoor               Id = 96
	MonsterEggStone            Id = 97
	MonsterEggCobblestone      Id = 97 | 1<<16
	MonsterEggStoneBrick       Id = 97 | 2<<16
	MonsterEggMossyStoneBrick  Id = 97 | 3<<16
	MonsterEggCrackedStone     Id = 97 | 4<<16
	MonsterEggChiseledStone    Id = 97 | 5<<16
	StoneBrick                 Id = 98
	StoneBrickMossy            Id = 98 | 1<<16
	StoneBrickCracked          Id = 98 | 2<<16
	StoneBrickChiseled         Id = 98 | 3<<16
	MushroomBrownBlock         Id = 99
	MushroomRedBlock           Id = 100
	IronBars                   Id = 101
	GlassPane                  Id = 102
	MelonBlock                 Id = 103
	PumpkinVine                Id = 104
	MelonVine                  Id = 105
	Vines                      Id = 106
	OakFenceGate               Id = 107
	BrickStairs                Id = 108
	StoneBrickStairs           Id = 109
	Mycelium                   Id = 110
	LilyPad                    Id = 111
	NetherBrickBlock           Id = 112
	NetherBrickFence           Id = 113
	NetherBrickStairs          Id = 114
	NetherWart                 Id = 115
	EnchantmentTable           Id = 116
	BrewingStandBlock          Id = 117
	CauldronBlock              Id = 118
	EndPortal                  Id = 119
	EndPortalFrame             Id = 120
	EndStone                   Id = 121
	DragonEgg                  Id = 122
	RedstoneLamp               Id = 123
	RedstoneLampOn             Id = 124
	OakDoubleSlab              Id = 125
	SpruceDoubleSlab           Id = 125 | 1<<16
	BirchDoubleSlab            Id = 125 | 2<<16
	JungleDoubleSlab           Id = 125 | 3<<16
	AcaciaDoubleSlab           Id = 125 | 4<<16
	DarkOakDoubleSlab          Id = 125 | 5<<16
	OakSlab                    Id = 126
	SpruceSlab                 Id = 126 | 1<<16
	BirchSlab                  Id = 126 | 2<<16
	JungleSlab                 Id = 126 | 3<<16
	AcaciaSlab                 Id = 126 | 4<<16
	DarkOakSlab                Id = 126 | 5<<16
	CocoaPlant                 Id = 127
	SandstoneStairs            Id = 128
	EmeraldOre                 Id = 129
	EnderChest                 Id = 130
	TripwireHook               Id = 131
	Tripwire                   Id = 132
	EmeraldBlock               Id = 133
	SpruceStairs               Id = 134
	BirchStairs                Id = 135
	JungleStairs               Id = 136
	CommandBlock               Id = 137
	Beacon                     Id = 138
	CobblestoneWall            Id = 139
	MossyCobblestoneWall       Id = 139 | 1<<16
	FlowerPotBlock             Id = 140
	CarrotCrop                 Id = 141
	PotatoCrop                 Id = 142
	WoodButton                 Id = 143
	HeadBlockSkeleton          Id = 144
	HeadBlockWither            Id = 144 | 1<<16
	HeadBlockZombie            Id = 144 | 2<<16
	HeadBlockSteve             Id = 144 | 3<<16
	HeadBlockCreeper           Id = 144 | 4<<16
	Anvil                      Id = 145
	AnvilSlightlyDamaged       Id = 145 | 1<<16
	AnvilVeryDamaged           Id = 145 | 2<<16
	TrappedChest               Id = 146
	WeightedPressurePlateLight Id = 147
	WeightedPressurePlateHeavy Id = 148
	RedstoneComparatorOff      Id = 149
	RedstoneComparatorOn       Id = 150
	DaylightSensor             Id = 151
	RedstoneBlock              Id = 152
	NetherQuartzOre            Id = 153
	Hopper                     Id = 154
	QuartzBlock                Id = 155
	QuartzBlockChiseled        Id = 155 | 1<<16
	QuartzBlockPillar          Id = 155 | 2<<16
	QuartzStairs               Id = 156
	RailActivator              Id = 157
	Dropper                    Id = 158
	WhiteStainedClay           Id = 159
	OrangeStainedClay          Id = 159 | 1<<16
	MagentaStainedClay         Id = 159 | 2<<16
	LightBlueStainedClay       Id = 159 | 3<<16
	YellowStainedClay          Id = 159 | 4<<16
	LimeStainedClay            Id = 159 | 5<<16
	PinkStainedClay            Id = 159 | 6<<16
	GrayStainedClay            Id = 159 | 7<<16
	LightGrayStainedClay       Id = 159 | 8<<16
	CyanStainedClay            Id = 159 | 9<<16
	PurpleStainedClay          Id = 159 | 10<<16
	BlueStainedClay            Id = 159 | 11<<16
	BrownStainedClay           Id = 159 | 12<<16
	GreenStainedClay           Id = 159 | 13<<16
	RedStainedClay             Id = 159 | 14<<16
	BlackStainedClay           Id = 159 | 15<<16
	WhiteStainedGlassPane      Id = 160
	OrangeStainedGlassPane     Id = 160 | 1<<16
	MagentaStainedGlassPane    Id = 160 | 2<<16
	LightBlueStainedGlassPane  Id = 160 | 3<<16
	YellowStainedGlassPane     Id = 160 | 4<<16
	LimeStainedGlassPane       Id = 160 | 5<<16
	PinkStainedGlassPane       Id = 160 | 6<<16
	GrayStainedGlassPane       Id = 160 | 7<<16
	LightGrayStainedGlassPane  Id = 160 | 8<<16
	CyanStainedGlassPane       Id = 160 | 9<<16
	PurpleStainedGlassPane     Id = 160 | 10<<16
	BlueStainedGlassPane       Id = 160 | 11<<16
	BrownStainedGlassPane      Id = 160 | 12<<16
	GreenStainedGlassPane      Id = 160 | 13<<16
	RedStainedGlassPane        Id = 160 | 14<<16
	BlackStainedGlassPane      Id = 160 | 15<<16
	AcaciaLeaves               Id = 161
	DarkOakLeaves              Id = 161 | 1<<16
	AcaciaLog                  Id = 162
	DarkOakLog                 Id = 162 | 1<<16
	AcaciaStairs               Id = 163
	DarkOakStairs              Id = 164
	SlimeBlock                 Id = 165
	Barrier                    Id = 166
	IronTrapdoor               Id = 167
	Prismarine                 Id = 168
	PrismarineBricks           Id = 168 | 1<<16
	PrismarineDark             Id = 168 | 2<<16
	SeaLantern                 Id = 169
	HayBale                    Id = 170
	WhiteCarpet                Id = 171
	OrangeCarpet               Id = 171 | 1<<16
	MagentaCarpet              Id = 171 | 2<<16
	LightBlueCarpet            Id = 171 | 3<<16
	YellowCarpet               Id = 171 | 4<<16
	LimeCarpet                 Id = 171 | 5<<16
	PinkCarpet                 Id = 171 | 6<<16
	GrayCarpet                 Id = 171 | 7<<16
	LightGrayCarpet            Id = 171 | 8<<16
	CyanCarpet                 Id = 171 | 9<<16
	PurpleCarpet               Id = 171 | 10<<16
	BlueCarpet                 Id = 171 | 11<<16
	BrownCarpet                Id = 171 | 12<<16
	GreenCarpet                Id = 171 | 13<<16
	RedCarpet                  Id = 171 | 14<<16
	BlackCarpet                Id = 171 | 15<<16
	HardenedClay               Id = 172
	CoalBlock                  Id = 173
	PackedIce                  Id = 174
	Sunflower                  Id = 175
	Lilac                      Id = 175 | 1<<16
	DoubleTallgrass            Id = 175 | 2<<16
	LargeFern                  Id = 175 | 3<<16
	Rosebush                   Id = 175 | 4<<16
	Peony                      Id = 175 | 5<<16
	BannerStandingBlock        Id = 176
	BannerWallBlock            Id = 177
	DaylightSensorInverted     Id = 178
	RedSandstone               Id = 179
	RedSandstoneChiseled       Id = 179 | 1<<16
	RedSandstoneSmooth         Id = 179 | 2<<16
	RedSandstoneStairs         Id = 180
	RedSandstoneDoubleSlab     Id = 181
	RedSandstoneSlab           Id = 182
	SpruceFenceGate            Id = 183
	BirchFenceGate             Id = 184
	JungleFenceGate            Id = 185
	DarkOakFenceGate           Id = 186
	AcaciaFenceGate            Id = 187
	SpruceFence                Id = 188
	BirchFence                 Id = 189
	JungleFence                Id = 190
	DarkOakFence               Id = 191
	AcaciaFence                Id = 192
	SpruceDoorBlock            Id = 193
	BirchDoorBlock             Id = 194
	JungleDoorBlock            Id = 195
	AcaciaDoorBlock            Id = 196
	DarkOakDoorBlock           Id = 197
	IronShovel                 Id = 256
	IronPickaxe                Id = 257
	IronAxe                    Id = 258
	FlintAndSteel              Id = 259
	Apple                      Id = 260
	Bow                        Id = 261
	Arrow                      Id = 262
	Coal                       Id = 263
	Charcoal                   Id = 263 | 1<<16
	Diamond                    Id = 264
	IronIngot                  Id = 265
	GoldIngot                  Id = 266
	IronSword                  Id = 267
	WoodSword                  Id = 268
	WoodShovel                 Id = 269
	WoodPickaxe                Id = 270
	WoodAxe                    Id = 271
	StoneSword                 Id = 272
	StoneShovel                Id = 273
	StonePickaxe               Id = 274
	StoneAxe                   Id = 275
	DiamondSword               Id = 276
	DiamondShovel              Id = 277
	DiamondPickaxe             Id = 278
	DiamondAxe                 Id = 279
	Stick                      Id = 280
	Bowl                       Id = 281
	MushroomStew               Id = 282
	GoldSword                  Id = 283
	GoldShovel                 Id = 284
	GoldPickaxe                Id = 285
	GoldAxe                    Id = 286
	String                     Id = 287
	Feather                    Id = 288
	Gunpowder                  Id = 289
	WoodHoe                    Id = 290
	StoneHoe                   Id = 291
	IronHoe                    Id = 292
	DiamondHoe                 Id = 293
	GoldHoe                    Id = 294
	WheatSeeds                 Id = 295
	Wheat                      Id = 296
	Bread                      Id = 297
	LeatherHelmet              Id = 298
	LeatherChestplate          Id = 299
	LeatherLeggings            Id = 300
	LeatherBoots               Id = 301
	ChainmailHelmet            Id = 302
	ChainmailChestplate        Id = 303
	ChainmailLeggings          Id = 304
	ChainmailBoots             Id = 305
	IronHelmet                 Id = 306
	IronChestplate             Id = 307
	IronLeggings               Id = 308
	IronBoots                  Id = 309
	DiamondHelmet              Id = 310
	DiamondChestplate          Id = 311
	DiamondLeggings            Id = 312
	DiamondBoots               Id = 313
	GoldHelmet                 Id = 314
	GoldChestplate             Id = 315
	GoldLeggings               Id = 316
	GoldBoots                  Id = 317
	Flint                      Id = 318
	PorkchopRaw                Id = 319
	PorkchopCooked             Id = 320
	Painting                   Id = 321
	GoldenApple                Id = 322
	EnchantedGoldenApple       Id = 322 | 1<<16
	Sign                       Id = 323
	OakDoor                    Id = 324
	Bucket                     Id = 325
	BucketWater                Id = 326
	BucketLava                 Id = 327
	Minecart                   Id = 328
	Saddle                     Id = 329
	IronDoor                   Id = 330
	RedstoneDust               Id = 331
	Snowball                   Id = 332
	Boat                       Id = 333
	Leather                    Id = 334
	BucketMilk                 Id = 335
	ClayBrick                  Id = 336
	Clay                       Id = 337
	Sugarcane                  Id = 338
	Paper                      Id = 339
	Book                       Id = 340
	Slimeball                  Id = 341
	MinecartStorage            Id = 342
	MinecartPowered            Id = 343
	Egg                        Id = 344
	Compass                    Id = 345
	FishingRod                 Id = 346
	Watch                      Id = 347
	GlowstoneDust              Id = 348
	FishRaw                    Id = 349
	SalmonRaw                  Id = 349 | 1<<16
	ClownfishRaw               Id = 349 | 2<<16
	PufferfishRaw              Id = 349 | 3<<16
	FishCooked                 Id = 350
	SalmonCooked               Id = 350 | 1<<16
	ClownfishCooked            Id = 350 | 2<<16
	PufferfishCooked           Id = 350 | 3<<16
	InkSac                     Id = 351
	RoseRedDye                 Id = 351 | 1<<16
	CactusGreenDye             Id = 351 | 2<<16
	CocoaBean                  Id = 351 | 3<<16
	LapisLazuli                Id = 351 | 4<<16
	PurpleDye                  Id = 351 | 5<<16
	CyanDye                    Id = 351 | 6<<16
	LightGrayDye               Id = 351 | 7<<16
	GrayDye                    Id = 351 | 8<<16
	PinkDye                    Id = 351 | 9<<16
	LimeDye                    Id = 351 | 10<<16
	DandelionYellowDye         Id = 351 | 11<<16
	LightBlueDye               Id = 351 | 12<<16
	MagentaDye                 Id = 351 | 13<<16
	OrangeDye                  Id = 351 | 14<<16
	Bonemeal                   Id = 351 | 15<<16
	Bone                       Id = 352
	Sugar                      Id = 353
	Cake                       Id = 354
	Bed                        Id = 355
	RedstoneRepeater           Id = 356
	Cookie                     Id = 357
	Map                        Id = 358
	Shears                     Id = 359
	MelonSlice                 Id = 360
	PumpkinSeeds               Id = 361
	MelonSeeds                 Id = 362
	BeefRaw                    Id = 363
	Steak                      Id = 364
	ChickenRaw                 Id = 365
	ChickenCooked              Id = 366
	RottenFlesh                Id = 367
	EnderPearl                 Id = 368
	BlazeRod                   Id = 369
	GhastTear                  Id = 370
	GoldNugget                 Id = 371
	NetherWartSeeds            Id = 372
	WaterBottle                Id = 373
	AwkwardPotion              Id = 373 | 16<<16
	ThickPotion                Id = 373 | 32<<16
	MundanePotion              Id = 373 | 64<<16
	RegenerationPotion045      Id = 373 | 8193<<16
	SwiftnessPotion300         Id = 373 | 8194<<16
	FireResistancePotion300    Id = 373 | 8195<<16
	PoisonPotion045            Id = 373 | 8196<<16
	HealingPotion              Id = 373 | 8197<<16
	NightVisionPotion300       Id = 373 | 8198<<16
	WeaknessPotion130          Id = 373 | 8200<<16
	StrengthPotion300          Id = 373 | 8201<<16
	SlownessPotion130          Id = 373 | 8202<<16
	HarmingPotion              Id = 373 | 8204<<16
	WaterBreathingPotion300    Id = 373 | 8205<<16
	InvisibilityPotion300      Id = 373 | 8206<<16
	RegenerationPotionII022    Id = 373 | 8225<<16
	SwiftnessPotionII130       Id = 373 | 8226<<16
	PoisonPotionII022          Id = 373 | 8228<<16
	HealingPotionII            Id = 373 | 8229<<16
	StrengthPotionII130        Id = 373 | 8233<<16
	LeapingPotionII130         Id = 373 | 8235<<16
	HarmingPotionII            Id = 373 | 8236<<16
	RegenerationPotion200      Id = 373 | 8257<<16
	SwiftnessPotion800         Id = 373 | 8258<<16
	FireResistancePotion800    Id = 373 | 8259<<16
	PoisonPotion200            Id = 373 | 8260<<16
	NightVisionPotion800       Id = 373 | 8262<<16
	WeaknessPotion400          Id = 373 | 8264<<16
	StrengthPotion800          Id = 373 | 8265<<16
	SlownessPotion400          Id = 373 | 8266<<16
	LeapingPotion300           Id = 373 | 8267<<16
	WaterBreathingPotion800    Id = 373 | 8269<<16
	InvisibilityPotion800      Id = 373 | 8270<<16
	RegenerationPotionII100    Id = 373 | 8289<<16
	SwiftnessPotionII400       Id = 373 | 8290<<16
	PoisonPotionII100          Id = 373 | 8292<<16
	StrengthPotionII400        Id = 373 | 8297<<16
	RegenerationSplash033      Id = 373 | 16385<<16
	SwiftnessSplash215         Id = 373 | 16386<<16
	FireResistanceSplash215    Id = 373 | 16387<<16
	PoisonSplash033            Id = 373 | 16388<<16
	HealingSplash              Id = 373 | 16389<<16
	NightVisionSplash215       Id = 373 | 16390<<16
	WeaknessSplash107          Id = 373 | 16392<<16
	StrengthSplash215          Id = 373 | 16393<<16
	SlownessSplash107          Id = 373 | 16394<<16
	HarmingSplash              Id = 373 | 16396<<16
	BreathingSplash            Id = 373 | 16397<<16
	InvisibilitySplash215      Id = 373 | 16398<<16
	RegenerationSplashII016    Id = 373 | 16417<<16
	SwiftnessSplashII107       Id = 373 | 16418<<16
	PoisonSplashII016          Id = 373 | 16420<<16
	HealingSplashII            Id = 373 | 16421<<16
	StrengthSplashII107        Id = 373 | 16425<<16
	LeapingSplashII107         Id = 373 | 16427<<16
	HarmingSplashII            Id = 373 | 16428<<16
	RegenerationSplash130      Id = 373 | 16449<<16
	SwiftnessSplash600         Id = 373 | 16450<<16
	FireResistanceSplash600    Id = 373 | 16451<<16
	PoisonSplash130            Id = 373 | 16452<<16
	NightVisionSplash600       Id = 373 | 16454<<16
	WeaknessSplash300          Id = 373 | 16456<<16
	StrengthSplash600          Id = 373 | 16457<<16
	SlownessSplash300          Id = 373 | 16458<<16
	LeapingSplash215           Id = 373 | 16459<<16
	BreathingSplash600         Id = 373 | 16461<<16
	InvisibilitySplash600      Id = 373 | 16462<<16
	RegenerationSplashII045    Id = 373 | 16481<<16
	SwiftnessSplashII300       Id = 373 | 16482<<16
	PoisonSplashII045          Id = 373 | 16484<<16
	StrengthSplashII300        Id = 373 | 16489<<16
	GlassBottle                Id = 374
	SpiderEye                  Id = 375
	FermentedSpiderEye         Id = 376
	BlazePowder                Id = 377
	MagmaCream                 Id = 378
	BrewingStand               Id = 379
	Cauldron                   Id = 380
	EyeOfEnder                 Id = 381
	GlisteringMelonSlice       Id = 382
	SpawnEggCreeper            Id = 383 | 50<<16
	SpawnEggSkeleton           Id = 383 | 51<<16
	SpawnEggSpider             Id = 383 | 52<<16
	SpawnEggZombie             Id = 383 | 54<<16
	SpawnEggSlime              Id = 383 | 55<<16
	SpawnEggGhast              Id = 383 | 56<<16
	SpawnEggZombiePigmen       Id = 383 | 57<<16
	SpawnEggEndermen           Id = 383 | 58<<16
	SpawnEggCaveSpider         Id = 383 | 59<<16
	SpawnEggSilverfish         Id = 383 | 60<<16
	SpawnEggBlaze              Id = 383 | 61<<16
	SpawnEggMagmaCube          Id = 383 | 62<<16
	SpawnEggBat                Id = 383 | 65<<16
	SpawnEggWitch              Id = 383 | 66<<16
	SpawnEggEndermite          Id = 383 | 67<<16
	SpawnEggGuardian           Id = 383 | 68<<16
	SpawnEggPig                Id = 383 | 90<<16
	SpawnEggSheep              Id = 383 | 91<<16
	SpawnEggCow                Id = 383 | 92<<16
	SpawnEggChicken            Id = 383 | 93<<16
	SpawnEggSquid              Id = 383 | 94<<16
	SpawnEggWolf               Id = 383 | 95<<16
	SpawnEggMooshroom          Id = 383 | 96<<16
	SpawnEggOcelot             Id = 383 | 98<<16
	SpawnEggHorse              Id = 383 | 100<<16
	SpawnEggRabbit             Id = 383 | 101<<16
	SpawnEggVillager           Id = 383 | 120<<16
	BottleOfEnchanting         Id = 384
	FireCharge                 Id = 385
	BookAndQuill               Id = 386
	WrittenBook                Id = 387
	Emerald                    Id = 388
	ItemFrame                  Id = 389
	FlowerPot                  Id = 390
	Carrot                     Id = 391
	Potato                     Id = 392
	PotatoBaked                Id = 393
	PotatoPoisonous            Id = 394
	EmptyMap                   Id = 395
	GoldenCarrot               Id = 396
	HeadSkeleton               Id = 397
	HeadWither                 Id = 397 | 1<<16
	HeadZombie                 Id = 397 | 2<<16
	HeadSteve                  Id = 397 | 3<<16
	HeadCreeper                Id = 397 | 4<<16
	CarrotOnAStick             Id = 398
	NetherStar                 Id = 399
	PumpkinPie                 Id = 400
	FireworkRocket             Id = 401
	FireworkStar               Id = 402
	EnchantedBook              Id = 403
	RedstoneComparator         Id = 404
	NetherBrick                Id = 405
	NetherQuartz               Id = 406
	MinecartTNT                Id = 407
	MinecartHopper             Id = 408
	PrismarineShard            Id = 409
	PrismarineCrystals         Id = 410
	RabbitRaw                  Id = 411
	RabbitCooked               Id = 412
	RabbitStew                 Id = 413
	RabbitFoot                 Id = 414
	RabbitHide                 Id = 415
	ArmorStand                 Id = 416
	IronHorseArmor             Id = 417
	GoldHorseArmor             Id = 418
	DiamondHorseArmor          Id = 419
	Lead                       Id = 420
	NameTag                    Id = 421
	MinecartCommandBlock       Id = 422
	MuttonRaw                  Id = 423
	MuttonCooked               Id = 424
	BlackBanner                Id = 425
	RedBanner                  Id = 425 | 1<<16
	GreenBanner                Id = 425 | 2<<16
	BrownBanner                Id = 425 | 3<<16
	BlueBanner                 Id = 425 | 4<<16
	PurpleBanner               Id = 425 | 5<<16
	CyanBanner                 Id = 425 | 6<<16
	LightGrayBanner            Id = 425 | 7<<16
	GrayBanner                 Id = 425 | 8<<16
	PinkBanner                 Id = 425 | 9<<16
	LimeBanner                 Id = 425 | 10<<16
	YellowBanner               Id = 425 | 11<<16
	LightBlueBanner            Id = 425 | 12<<16
	MagentaBanner              Id = 425 | 13<<16
	OrangeBanner               Id = 425 | 14<<16
	WhiteBanner                Id = 425 | 15<<16
	SpruceDoor                 Id = 427
	BirchDoor                  Id = 428
	JungleDoor                 Id = 429
	AcaciaDoor                 Id = 430
	DarkOakDoor                Id = 431
	MusicDisk13                Id = 2256
	MusicDiskCat               Id = 2257
	MusicDiskBlocks            Id = 2258
	MusicDiskChirp             Id = 2259
	MusicDiskFar               Id = 2260
	MusicDiskMall              Id = 2261
	MusicDiskMellohi           Id = 2262
	MusicDiskStal              Id = 2263
	MusicDiskStrad             Id = 2264
	MusicDiskWard              Id = 2265
	MusicDisk11                Id = 2266
	MusicDiskWait              Id = 2267
)

// Consumables defines a set with all items one can eat or drink.
// This includes (splash) potions.
var Consumables = []Id{
	EnchantedGoldenApple,
	GoldenApple,
	PorkchopRaw,
	PorkchopCooked,
	Bread,
	Apple,
	FishRaw,
	SalmonRaw,
	ClownfishRaw,
	PufferfishRaw,
	FishCooked,
	SalmonCooked,
	ClownfishCooked,
	PufferfishCooked,
	Cake,
	Cookie,
	MelonSlice,
	BeefRaw,
	Steak,
	ChickenRaw,
	ChickenCooked,
	RottenFlesh,
	BucketMilk,
	Carrot,
	Potato,
	PotatoBaked,
	PotatoPoisonous,
	GoldenCarrot,
	PumpkinPie,
	RabbitRaw,
	RabbitCooked,
	RabbitStew,
	MuttonRaw,
	MuttonCooked,
	MushroomStew,

	RegenerationPotion045,
	SwiftnessPotion300,
	FireResistancePotion300,
	PoisonPotion045,
	HealingPotion,
	NightVisionPotion300,
	WeaknessPotion130,
	StrengthPotion300,
	SlownessPotion130,
	HarmingPotion,
	WaterBreathingPotion300,
	InvisibilityPotion300,
	RegenerationPotionII022,
	SwiftnessPotionII130,
	PoisonPotionII022,
	HealingPotionII,
	StrengthPotionII130,
	LeapingPotionII130,
	HarmingPotionII,
	RegenerationPotion200,
	SwiftnessPotion800,
	FireResistancePotion800,
	PoisonPotion200,
	NightVisionPotion800,
	WeaknessPotion400,
	StrengthPotion800,
	SlownessPotion400,
	LeapingPotion300,
	WaterBreathingPotion800,
	InvisibilityPotion800,
	RegenerationPotionII100,
	SwiftnessPotionII400,
	PoisonPotionII100,
	StrengthPotionII400,
	RegenerationSplash033,
	SwiftnessSplash215,
	FireResistanceSplash215,
	PoisonSplash033,
	HealingSplash,
	NightVisionSplash215,
	WeaknessSplash107,
	StrengthSplash215,
	SlownessSplash107,
	HarmingSplash,
	BreathingSplash,
	InvisibilitySplash215,
	RegenerationSplashII016,
	SwiftnessSplashII107,
	PoisonSplashII016,
	HealingSplashII,
	StrengthSplashII107,
	LeapingSplashII107,
	HarmingSplashII,
	RegenerationSplash130,
	SwiftnessSplash600,
	FireResistanceSplash600,
	PoisonSplash130,
	NightVisionSplash600,
	WeaknessSplash300,
	StrengthSplash600,
	SlownessSplash300,
	LeapingSplash215,
	BreathingSplash600,
	InvisibilitySplash600,
	RegenerationSplashII045,
	SwiftnessSplashII300,
	PoisonSplashII045,
	StrengthSplashII300,
}

// MusicDisks defines a set of all music disks.
var MusicDisks = []Id{
	MusicDisk13,
	MusicDiskCat,
	MusicDiskBlocks,
	MusicDiskChirp,
	MusicDiskFar,
	MusicDiskMall,
	MusicDiskMellohi,
	MusicDiskStal,
	MusicDiskStrad,
	MusicDiskWard,
	MusicDisk11,
	MusicDiskWait,
}

// SpawnEggs defines a set of all mob spawning items.
var Spawners = []Id{
	MobSpawner,
	SpawnEggCreeper,
	SpawnEggSkeleton,
	SpawnEggSpider,
	SpawnEggZombie,
	SpawnEggSlime,
	SpawnEggGhast,
	SpawnEggZombiePigmen,
	SpawnEggEndermen,
	SpawnEggCaveSpider,
	SpawnEggSilverfish,
	SpawnEggBlaze,
	SpawnEggMagmaCube,
	SpawnEggBat,
	SpawnEggWitch,
	SpawnEggEndermite,
	SpawnEggGuardian,
	SpawnEggPig,
	SpawnEggSheep,
	SpawnEggCow,
	SpawnEggChicken,
	SpawnEggSquid,
	SpawnEggWolf,
	SpawnEggMooshroom,
	SpawnEggOcelot,
	SpawnEggHorse,
	SpawnEggRabbit,
	SpawnEggVillager,
}

// Dyes defines a set of all colour dyes.
var Dyes = []Id{
	InkSac,
	RoseRedDye,
	CactusGreenDye,
	CocoaBean,
	LapisLazuli,
	PurpleDye,
	CyanDye,
	LightGrayDye,
	GrayDye,
	PinkDye,
	LimeDye,
	DandelionYellowDye,
	LightBlueDye,
	MagentaDye,
	OrangeDye,
	Bonemeal,
}

// Machinery defines a set of all items related to minecarts, rails,
// redstone and triggers.
var Machinery = []Id{
	TrappedChest,
	Tripwire,
	TripwireHook,
	StoneButton,
	WoodButton,
	StonePressurePlate,
	WoodPressurePlate,
	WeightedPressurePlateLight,
	WeightedPressurePlateHeavy,
	DaylightSensor,
	DaylightSensorInverted,
	RedstoneBlock,
	RedstoneTorchOff,
	RedstoneTorch,
	RedstoneDust,
	RedstoneWire,
	RedstoneComparator,
	RedstoneComparatorOff,
	RedstoneComparatorOn,
	RedstoneRepeater,
	RedstoneRepeaterBlockOff,
	RedstoneRepeaterBlockOn,
	Piston,
	PistonHead,
	PistonMoving,
	StickyPiston,
	Rail,
	RailActivator,
	RailDetector,
	RailActivator,
	RailPowered,
	Minecart,
	MinecartCommandBlock,
	MinecartHopper,
	MinecartPowered,
	MinecartStorage,
	MinecartTNT,
	CommandBlock,
}

// Armor defines all armor sets.
var Armor = []Id{
	LeatherHelmet,
	LeatherChestplate,
	LeatherLeggings,
	LeatherBoots,
	ChainmailHelmet,
	ChainmailChestplate,
	ChainmailLeggings,
	ChainmailBoots,
	IronHelmet,
	IronChestplate,
	IronLeggings,
	IronBoots,
	DiamondHelmet,
	DiamondChestplate,
	DiamondLeggings,
	DiamondBoots,
	GoldHelmet,
	GoldChestplate,
	GoldLeggings,
	GoldBoots,
}

// Tools defines a set of all tools.
// These cover: bows, arrows, swords, axes, hoes, pickaxes, shovels,
// flint&steel, shears, buckets, compass,etc.
var Tools = []Id{
	Compass,
	Watch,
	Map,
	EmptyMap,
	Shears,
	FishingRod,
	CarrotOnAStick,
	Bucket,
	BucketLava,
	BucketWater,
	BucketMilk,
	Bow,
	Arrow,
	FlintAndSteel,
	IronSword,
	WoodSword,
	WoodShovel,
	WoodPickaxe,
	WoodPickaxe,
	WoodAxe,
	StoneSword,
	StoneShovel,
	StonePickaxe,
	StoneAxe,
	DiamondSword,
	DiamondShovel,
	DiamondPickaxe,
	DiamondAxe,
	GoldSword,
	GoldShovel,
	GoldPickaxe,
	GoldAxe,
	WoodHoe,
	StoneHoe,
	IronHoe,
	DiamondHoe,
	GoldHoe,
}

// Carpets defines a set of all carpet colours.
var Carpets = []Id{
	WhiteCarpet,
	OrangeCarpet,
	MagentaCarpet,
	LightBlueCarpet,
	YellowCarpet,
	LimeCarpet,
	PinkCarpet,
	GrayCarpet,
	LightGrayCarpet,
	CyanCarpet,
	PurpleCarpet,
	BlueCarpet,
	BrownCarpet,
	GreenCarpet,
	RedCarpet,
	BlackCarpet,
}

// PrismarineBlocks defines a set of all prismarine blocks.
var PrismarineBlocks = []Id{
	Prismarine,
	PrismarineBricks,
	PrismarineDark,
}

// GlassPanes defines a set of all glass panes.
// This includes regular and stained glass.
var GlassPanes = []Id{
	GlassPane,
	WhiteStainedGlassPane,
	OrangeStainedGlassPane,
	MagentaStainedGlassPane,
	LightBlueStainedGlassPane,
	YellowStainedGlassPane,
	LimeStainedGlassPane,
	PinkStainedGlassPane,
	GrayStainedGlassPane,
	LightGrayStainedGlassPane,
	CyanStainedGlassPane,
	PurpleStainedGlassPane,
	BlueStainedGlassPane,
	BrownStainedGlassPane,
	GreenStainedGlassPane,
	RedStainedGlassPane,
	BlackStainedGlassPane,
}

// ClayBlocks defines a set of all clay blocks.
// This includes regular- and stained clay.
var ClayBlocks = []Id{
	ClayBlock,
	WhiteStainedClay,
	OrangeStainedClay,
	MagentaStainedClay,
	LightBlueStainedClay,
	YellowStainedClay,
	LimeStainedClay,
	PinkStainedClay,
	GrayStainedClay,
	LightGrayStainedClay,
	CyanStainedClay,
	PurpleStainedClay,
	BlueStainedClay,
	BrownStainedClay,
	GreenStainedClay,
	RedStainedClay,
	BlackStainedClay,
}

// QuartzBlocks defines a set of all quarts blocks.
var QuartzBlocks = []Id{
	QuartzBlock,
	QuartzBlockChiseled,
	QuartzBlockPillar,
	QuartzDoubleSlab,
	QuartzSlab,
}

// AnvilBlocks defines a set of all anvil versions.
var AnvilBlocks = []Id{
	Anvil,
	AnvilSlightlyDamaged,
	AnvilVeryDamaged,
}

// Heads defines a set of all head blocks and items.
var Heads = []Id{
	HeadBlockSkeleton,
	HeadBlockWither,
	HeadBlockZombie,
	HeadBlockSteve,
	HeadBlockCreeper,
	HeadSkeleton,
	HeadWither,
	HeadZombie,
	HeadSteve,
	HeadCreeper,
}

// FenceBlocks defines a set of all fence related blocks.
var FenceBlocks = []Id{
	NetherBrickFence,
	OakFence,
	SpruceFence,
	BirchFence,
	JungleFence,
	DarkOakFence,
	AcaciaFence,
	OakFenceGate,
	SpruceFenceGate,
	BirchFenceGate,
	JungleFenceGate,
	DarkOakFenceGate,
	AcaciaFenceGate,
}

// StoneBrickBlocks defines a set of all stone brick blocks.
var StoneBrickBlocks = []Id{
	StoneBrick,
	StoneBrickMossy,
	StoneBrickCracked,
	StoneBrickChiseled,
}

// MonsterEggBlocks defines a set of all monster egg types.
var MonsterEggBlocks = []Id{
	MonsterEggStone,
	MonsterEggCobblestone,
	MonsterEggStoneBrick,
	MonsterEggMossyStoneBrick,
	MonsterEggCrackedStone,
	MonsterEggChiseledStone,
}

// GlassBlocks defines a set of all glass blocks.
// This includes regular and stained glass.
var GlassBlocks = []Id{
	Glass,
	WhiteStainedGlass,
	OrangeStainedGlass,
	MagentaStainedGlass,
	LightBlueStainedGlass,
	YellowStainedGlass,
	LimeStainedGlass,
	PinkStainedGlass,
	GrayStainedGlass,
	LightGrayStainedGlass,
	CyanStainedGlass,
	PurpleStainedGlass,
	BlueStainedGlass,
	BrownStainedGlass,
	GreenStainedGlass,
	RedStainedGlass,
	BlackStainedGlass,
}

// Doors defines a set of all door blocks and items.
// This includes fence gates and trapdoors.
var Doors = []Id{
	OakDoorBlock,
	IronDoorBlock,
	SpruceDoorBlock,
	BirchDoorBlock,
	JungleDoorBlock,
	AcaciaDoorBlock,
	DarkOakDoorBlock,
	IronTrapdoor,
	WoodTrapdoor,
	OakDoor,
	SpruceDoor,
	BirchDoor,
	JungleDoor,
	AcaciaDoor,
	DarkOakDoor,
	OakFenceGate,
	SpruceFenceGate,
	BirchFenceGate,
	JungleFenceGate,
	DarkOakFenceGate,
	AcaciaFenceGate,
}

// Signs defines a set of all sign/banner blocks and items.
var Signs = []Id{
	SignBlock,
	Signwall,
	Sign,
	BannerStandingBlock,
	BannerWallBlock,
	BlackBanner,
	RedBanner,
	GreenBanner,
	BrownBanner,
	BlueBanner,
	PurpleBanner,
	CyanBanner,
	LightGrayBanner,
	GrayBanner,
	PinkBanner,
	LimeBanner,
	YellowBanner,
	LightBlueBanner,
	MagentaBanner,
	OrangeBanner,
	WhiteBanner,
}

// FurnaceBlocks defines a set of all furnace components.
var FurnaceBlocks = []Id{
	Furnace,
	FurnaceSmelting,
}

// StairBlocks defines a set of all stairs.
var StairBlocks = []Id{
	CobblestoneStairs,
	BrickStairs,
	StoneBrickStairs,
	NetherBrickStairs,
	SandstoneStairs,
	RedSandstoneStairs,
	OakStairs,
	SpruceStairs,
	BirchStairs,
	JungleStairs,
	AcaciaStairs,
	DarkOakStairs,
	QuartzStairs,
}

// SlabBlocks defines a set of all single slab blocks.
var SlabBlocks = []Id{
	StoneSlab,
	SandstoneSlab,
	CobblestoneSlab,
	BrickSlab,
	StoneBrickSlab,
	NetherBrickSlab,
	QuartzSlab,
	WoodSlab,
	OakSlab,
	SpruceSlab,
	BirchSlab,
	JungleSlab,
	AcaciaSlab,
	DarkOakSlab,
}

// DoubleSlabBlocks defines a set of all double slab blocks.
var DoubleSlabBlocks = []Id{
	Stone,
	SandstoneDoubleSlab,
	CobblestoneDoubleSlab,
	BrickDoubleSlab,
	StoneBrickDoubleSlab,
	NetherBrickDoubleSlab,
	QuartzDoubleSlab,
	SmoothStoneDoubleSlab,
	SmoothSandstoneDoubleSlab,
	WoodDoubleSlab,
	OakDoubleSlab,
	SpruceDoubleSlab,
	BirchDoubleSlab,
	JungleDoubleSlab,
	AcaciaDoubleSlab,
	DarkOakDoubleSlab,
}

// Plants defines a set of all flowers and plants.
// Covering: flowers, vines, crops, pumpkins/melons, lily pads etc.
// Basically anything that grows with the exception of  tree saplings;
// they have a separate group.
var Plants = []Id{
	Dandelion,
	Poppy,
	BlueOrchid,
	Allium,
	AzureBluet,
	RedTulip,
	OrangeTulip,
	WhiteTulip,
	PinkTulip,
	OxeyeDaisy,
	Sunflower,
	Lilac,
	Rosebush,
	Peony,
	TallGrassDeadShrub,
	TallGrass,
	TallGrassFern,
	DoubleTallgrass,
	LargeFern,
	CarrotCrop,
	PotatoCrop,
	CocoaPlant,
	DeadShrub,
	MushroomBrownBlock,
	MushroomRedBlock,
	SugarcaneBlock,
	WheatCrop,
	Cactus,
	Pumpkin,
	PumpkinVine,
	MelonVine,
	Vines,
	MelonBlock,
	LilyPad,
	NetherWart,
}

// WoolBlocks defines a set of all known wool blocks.
var WoolBlocks = []Id{
	WhiteWool,
	OrangeWool,
	MagentaWool,
	LightBlueWool,
	YellowWool,
	LimeWool,
	PinkWool,
	GrayWool,
	LightGrayWool,
	CyanWool,
	PurpleWool,
	BlueWool,
	BrownWool,
	GreenWool,
	RedWool,
	BlackWool,
}

// SandstoneBlocks defines a set of all known sandstone blocks.
// This includes red sandstone, slabs and stairs.
var SandstoneBlocks = []Id{
	Sandstone,
	SandstoneChiseled,
	SandstoneSmooth,
	SandstoneSlab,
	SandstoneDoubleSlab,
	SandstoneStairs,
	RedSandstone,
	RedSandstoneChiseled,
	RedSandstoneSmooth,
	RedSandstoneSlab,
	RedSandstoneDoubleSlab,
	RedSandstoneStairs,
}

// SpongeBlocks defines a set of all known sponge blocks.
var SpongeBlocks = []Id{
	Sponge,
	SpongeWet,
}

// LeaveBlocks defines a set of all known leave blocks.
var LeaveBlocks = []Id{
	OakLeaves,
	SpruceLeaves,
	BirchLeaves,
	JungleLeaves,
	AcaciaLeaves,
	DarkOakLeaves,
}

// WoodBlocks defines a set of all known wood blocks.
var WoodBlocks = []Id{
	OakLog,
	SpruceLog,
	BirchLog,
	JungleLog,
	Oak4Log,
	Oak5Log,
	AcaciaLog,
	DarkOakLog,
}

// SandBlocks defines a set of all sand blocks.
// This covers regular- and red sand and soul sand.
var SandBlocks = []Id{
	Sand,
	RedSand,
	SoulSand,
}

// Saplings defines a set of all saplings.
var Saplings = []Id{
	OakSapling,
	SpruceSapling,
	BirchSapling,
	JungleSapling,
	AcaciaSapling,
	DarkOakSapling,
}

// WoodPlanks defines a set of all known wooden plank types.
var WoodPlanks = []Id{
	OakPlanks,
	SprucePlanks,
	BirchPlanks,
	JunglePlanks,
	AcaciaPlanks,
	DarkOakPlanks,
}
