## Recipe

Package recipe defines an interface and listing of all known Minecraft recipes
and facilities to create recipe trees.

The purpose is to combine it with the mcra search facilities, so one may
provide the library with a recipe for N number of X items. It will then
generate a tree of all the recipes needed to create the items, all the way
down to items which must be mined/farmed. And when combined with an existing
world, it will provide a listing of all the locations where these resources
may be found. Essentially not only telling you how to make the item, but
providing you with a shopping list of elementary components you will have to
go out and find.

Finally, it can be combined with a persistent resource database you can create
and maintain for a specific world. It will draw from this database when
determining how many times a given recipe needs to be crafted.
