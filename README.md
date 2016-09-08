plausible
=========

[![Build Status](https://travis-ci.org/alrs/plausible.svg?branch=master)](https://travis-ci.org/alrs/plausible)
[![Go Report Card](https://goreportcard.com/badge/github.com/alrs/plausible)](https://goreportcard.com/report/github.com/alrs/plausible)

The plausible program uses the database of MAC vendor prefixes from the
Wireshark manuf file to generate MAC addresses with the first three
octets corresponding to a desired vendor, and the second three octets
generated at random.

plausible assumes that Wireshark is installed via package management on
a Debian system. If you aren't on a Debian system, you can download the
manuf DB with:

`make manuf`

You can choose the location of your manuf file with the `-d` flag.
The `-l` flag lists all available manufacturer codes.
The `-m` flag selects a manufacturer to use for generating a MAC address.
