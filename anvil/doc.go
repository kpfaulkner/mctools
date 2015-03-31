// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

/*
Package anvil reads level.dat and .mca region files which make up Minecraft
worlds.

This library allows loading, modifying and saving of all the information.
This means it can be used to analyze the contents of a world, but also change
it to create new worlds.


Warning

The code defines read->write->read roundtrip tests to ensure the written data
is consistent and correct. Even so, you should __always create backups__ of any
world you are about to change. I am __not__ responsible for any damage you
may do to your world data or Minecraft itself.

This package has been tested with world data created in Vanilla Minecraft 1.8.3+.
There are no guarantees that older worlds can be read or that changes to such a
world will still load up in Minecraft. Test your changes thoroughly. And when
you're done testing, test some more.

Additionally, don't read/write world data which is currently being used by
Minecraft itself. The behaviour of this undefined and will very likely end
up in tears.

*/
package anvil
