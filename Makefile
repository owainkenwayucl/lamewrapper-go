# This program wraps the LAME encoder with my standard options.
# Dr Owain Kenway

# Where it is distributed, it is done so under a 4 clause,
# BSD-style license (see LICENSE)

sources = src/lamewrapper.go
installdir = $(HOME)/bin

all: lamewrapper

lamewrapper: $(sources) Makefile
	go build $(sources)

install: lamewrapper
	cp lamewrapper $(installdir)/lamewrapper

clean:
	rm lamewrapper
