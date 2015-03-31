// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

/*
package mctools holds a collection of tools used to read, edit and write Minecraft
world data. Refer to the respective sub-packages for details on their purpose.

The code targets Minecraft 1.8+. Older worlds may work or they may not.
Use at your own risk.


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
package mctools
