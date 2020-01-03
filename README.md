# goflatten
> `goflatten` is a very small utility that flattens folders into a single file.

`goflatten` is a utility that flattens a nested directory into a flat one. 

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

- **rootDir** - Specifies the starting root directory. By default, the current working directory is the starting directory.

- **skipHidden** - If set to `false`, hidden files will be moved to the root directory. By default, this is set to `true`.

- **dryRun** - If set to `false`, `goflatten` will physically move nested files to the root directory. This is set to `true` by default.



