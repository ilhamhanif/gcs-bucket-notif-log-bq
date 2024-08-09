#!/bin/sh

go build -gccgoflags "-s -w" sequencer.go | strip sequencer