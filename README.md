Lamewrapper version (ii)
========================

This program wraps the LAME encoder with my standard options.

Dr Owain Kenway

Where it is distributed, it is done so under a 4 clause,
BSD-style license (see LICENSE.txt)

Install
-------

To install, make sure you have go installed, modify the Makefile with
your options and then issue:

```
make && make install
```

Usage
-----

The lamewrapper executable takes input from stdin.

```
./lamewrapper
```

You can also pass lamewrapper arguments as follows:

```
./lamewrapper <file> <song> <artist> <album> <genre> 
```
