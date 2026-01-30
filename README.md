# DPM
Development project management tool

## Project Layout

```sh
./dpm       # Implements the actions of dpm
./cli       # CLI interface to execute dpm actions
./internal  # Any non-dpm specific code (e.g. library like packages)
```


## Features

# Clone

Clones git projects into a configurable projects home directory in an organised manner

```
$ dpm clone git@github.com/iainvm/dpm

```

# List

List the git projects found in the configurable projects home directory

```
$ dpm list

```
