# goflatten
[![Go Report Card](https://goreportcard.com/badge/github.com/ShiraazMoollatjie/goflatten?style=flat-square)](https://goreportcard.com/report/github.com/ShiraazMoollatjie/goflatten)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/ShiraazMoollatjie/goflatten)
[![Build status](https://ci.appveyor.com/api/projects/status/v0x22haovyo1y396/branch/master?svg=true)](https://ci.appveyor.com/project/ShiraazMoollatjie/goflatten/branch/master)

> `goflatten` is a utility that flattens a nested directory of files into a flat one. 

For example, suppose that you have this structure:

```
pricelist
  -> 2019-01-01
    -> 2019-01-01.csv
  -> 2019-01-02
    -> 2019-01-02.csv
```

After running `goflatten` the resulting structure is:

```
pricelist
  -> 2019-01-01.csv
  -> 2019-01-02.csv
  -> 2019-01-01
  -> 2019-01-02
```

# Installing and running

To install `goflatten` run:

```sh
go get -u github.com/ShiraazMoollatjie/goflatten
```

Thereafter, you can simply run:

```
goflatten
```

# Flags

The following flags exist for `goflatten`:

- **dryRun** - If set to `false`, `goflatten` will physically move nested files to the root directory. This is set to `true` by default.

- **rootDir** - Specifies the starting root directory. By default, the current working directory is the starting directory.

- **skipDupes** - If set to `true`, duplicate files will not be overwritten when moving them to the root directory. By default, this is set to `true`.
