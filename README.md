# Heck!

[![Build Status](https://travis-ci.org/bfontaine/heck.svg?branch=master)](https://travis-ci.org/bfontaine/heck)
<!--
[![Coverage Status](https://coveralls.io/repos/bfontaine/heck/badge.svg?branch=master)](https://coveralls.io/r/bfontaine/heck?branch=master)
-->

**Heck** is an hexadecimal calculator CLI tool.

## Install

    go get github.com/bfontaine/heck

### Dependencies

* Go 1.1 or higher

## Usage

    $ heck

You’ll get a prompt from which you can use the usual operations, but all
numbers are in hexadecimal.

    $ heck
    > 5+5
    A         (10)
    > 15*2
    2A        (42)
    > (CA + FE) * (BA + BE)
    29DC0     (171456)

Heck prints both the hexadecimal and the decimal representation of the result.
If both are the same, only one is printed:

    > 2+2
    4

### Operators

* Values: `int64` for now, both positive and negative
* Binary operators: `+`, `-`, `/`, `*`
* Parentheses to group operations
