## NBT (Named Binary Tag)

NBT (Named Binary Tag) is a tag based binary format designed to carry large
amounts of binary data with smaller amounts of meta data. An NBT file
consists of a single compressed, named tag of type `TAG_Compound`, which
holds an arbitrary amount of nested tags.


### Usage

Unmarshal parses the uncompressed, NBT-encoded data and stores the result
in the value pointed to by v. This is most commonly a struct, which reflects
the NBT structure in its fields. Field tags can be used to map these fields to
appropriate tag names in the NBT data. For example:

	type Level struct {
		Name             string    `nbt:"levelName"`
		Generator        string    `nbt:"generatorName"`
		GeneratorOptions string    `nbt:"generatorOptions"`
		LastPlayed       time.Time `nbt:"LastPlayed", time:"2006-01-02T15:04:05"`
		Seed             int64     `nbt:"RandomSeed"`
		SpawnX           int32     `nbt:"SpawnX"`
		SpawnY           int32     `nbt:"SpawnY"`
		SpawnZ           int32     `nbt:"SpawnZ"`
		MapFeatures      bool      `nbt:"MapFeatures"`
		AllowCommands    bool      `nbt:"allowCommands"`
	}

	gz, err := gzip.NewReader(r)
	...

	var level Level
	err = nbt.Unmarshal(gz, &level)
	...

Data can be re-encoded with the `Marshal` call, or directly through the
`Encoder` type:

	var level Level
	level.Name = "test"
	...

	gz := gzip.NewWriter(w)
	err = nbt.Marshal(gz, &level)
	...


### Type compatibility

Some implicit type conversions can be achieved by adhering to the following
guidelines.

```
    Tag Id         | Assignable to field | Note
    =======================================================================
    TAG_Byte       | int8, uint8         |
                   | int16, uint16       |
                   | int32, uint32       |
                   | int64, uint64       |
                   | bool                | true if the tag value is not 0
    -----------------------------------------------------------------------
    TAG_Short      | int16, uint16       |
                   | int32, uint32       |
                   | int64, uint64       |
    -----------------------------------------------------------------------
    TAG_Int        | int32, uint32       |
                   | int64, uint64       |
    -----------------------------------------------------------------------
    TAG_Long       | int64, uint64       |
                   | time.Time           | Expects a Unix timestamp.
    -----------------------------------------------------------------------
    TAG_Float      | float32             |
                   | float64             |
    -----------------------------------------------------------------------
    TAG_Double     | float64             |
    -----------------------------------------------------------------------
    TAG_Byte_Array | []int8, []uint8     |
    -----------------------------------------------------------------------
    TAG_Int_Array  | []int32, []uint32   |
    -----------------------------------------------------------------------
    TAG_String     | string              |
                   | bool                | Parsed using strconv.ParseBool()
    -----------------------------------------------------------------------
    TAG_List       | []T, []*T           |
    -----------------------------------------------------------------------
    Tag_Compound   | T, *T               |
    -----------------------------------------------------------------------
```

Any other, incompatible assignment will result in a parse error.
Exceptions are type 'aliases' which are directly convertible to any of
the types listed above. For instance:

	type GameMode int32

	type Level struct {
		...
		Mode   GameMode  `nbt:"gameType"` // == TAG_Int("gameType")
		...
	}

If the encoder should not output a tag for an empty field, append the
`omitempty` value to the struct field tag. For example:

	type T struct {
		Data []byte  `nbt:"data,omitempty"
	}

If `len(T.Data) == 0`, the encoder will ignore this field and no tag is
emitted.
