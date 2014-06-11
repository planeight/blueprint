/*

The Blueprint bootstrapping mechanism is intended to enable building a source
tree using a Blueprint-based build system that is embedded (as source) in that
source tree.  The only prerequisites for performing such a build are:
	1. A Ninja binary
	2. A script interpreter (e.g. Bash or Python)
	3. A Go toolchain

The bootstrapping process is intended to be customized for the source tree in
which it is embedded.  As an initial example, we'll examine the bootstrapping
system used to build the core Blueprint packages.  Bootstrapping the core
Blueprint packages involves two files:

	bootstrap.bash:		When this script is run it initializes the current
						working directory as a build output directory.  It does
						this by first automatically determining the root source
						directory and Go build environment.  It then uses those
						values to do a simple string replacement over the
						build.ninja.in file contents, and places the result into
						the current working directory.

	build.ninja.in:		This file is generated by passing all the Blueprint

						files needed to construct the primary builder


*/
package bootstrap