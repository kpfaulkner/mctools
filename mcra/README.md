## mcra

Package mcra defines an API to query Minecraft world data.
Specifically intended to tally and locate sets of specific resources.

This code has been run on a few really old worlds, as well as a brand
new one generated in 1.8.3 and found to be working in all of them.

It should be noted that no attempts have been made at speed optimization'
as far as block search performance goes. Specifically when searching all
regions in a world, things can get rather slow.

The library operates directly on the world's region data. Nothing is
cached or indexed in any way.


### Usage

Loading a world and a specific region:

	world, err := mcra.Open(WorldPath)
	if err != nil {
		log.Fatal(err)
	}

	region, err = world.Region(mcra.DimensionOverworld, -1, 0)
	defer region.Close()


Finding blocks in a region:

	// Find only sand and sandstone blocks.
	result := FindInRegion(region, NewInclusionQuery(
		item.Sand,
		item.Sandstone,
	))

	...

	// Find everything, except Air and bedrock.
	result := FindInRegion(region, NewExclusionQuery(
		item.Air,
		item.Bedrock,
	))

	// Find all redstone ore within a radius from a given point.
	result := FindInRegion(region, NewRadiusQuery(
		Location{...},
		100,
		item.RedstoneOre,
		item.RedstoneOreGlowing,
	))


Tallying redstone and diamond ores in a region:

	tally := TallyInRegion(
		region,
		item.RedstoneOre,
		item.RedstoneOreGlowing,
		item.DiamondOre,
	)

	for k, v := range tally {
		fmt.Println(k, v)
	}

	// yields:
	//
	//     RedstoneOre 108
	//     DiamondOre 25


Tally all resources in a region:

	tally := TallyInRegion(region)
	...

	// yields:
	//
	//     Air 65184
	//     Sandstone 29709
	//     Sand 13629
	//     Granite 3655
	//     CoalOre 832
	//     Andesite 3975
	//     DiamondOre 25
	//     DeadShrub 3
	//     Bedrock 16043
	//     LavaNoSpread 1200
	//     Dirt 1765
	//     IronOre 512
	//     Grass 274
	//     Diorite 3183
	//     GoldOre 35
	//     TallGrass 8
	//     Stone 289207
	//     WaterNoSpread 44
	//     RedstoneOre 108
	//     Gravel 671
	//     LapisLazuliOre 16
	//     Cactus 2


### See it at work

This is a screenshot of a new (1.8.3) world in Vanilla MC. The output in the
console shows the locations of diamond ores in region `-1, 0` for a world with
seed `2856811980450068479`. This is the world data in the `testdata/newworld`
directory.

![screenshot](https://raw.githubusercontent.com/jteeuwen/mctools/master/testdata/closeenough.png)

