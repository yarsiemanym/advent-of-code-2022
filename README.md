# Advent of Code 2022

Using the [Advent of Code 2022](https://adventofcode.com/2022/) puzzles to become more proficient in [Go](https://golang.org/).

## Environment Variables

### AOC_SESSION_TOKEN

Required for this application to automatically download your user-specific input files for each puzzle. You can find your session token in adventofcode.com's session cookie in your browser.

If not set, the application will not be able to automatically download puzzle inputs. You will have to download each manually and save it to the correct location, e.g. `./day01/input.txt`.

### AOC_LOG_LEVEL

Sets the [logrus](https://github.com/Sirupsen/logrus) log level. Valid values are `panic`, `fatal`, `error`, `warn`, `info`, `debug`, and `trace`. Defaults to `warn` if not set.

## Makefile

This repo's [`Makefile`](./Makefile) automates many useful operations.

### make setup

Downloads and installs [Go 1.19.3](https://go.dev/dl/) and downloads all of this application's package dependencies.

### make test

Runs all unit tests.

### make build

Builds the `advent-of-code-2022` executable in the current directory.

### make clean

Deletes all build artifacts from the current directory, i.e. the `advent-of-code-2022` executable.

### make run

Builds the `advent-of-code-2022` executable and runs a single day's solution. Specify the day by setting the `DAY` parameter in the command line. Optionally, you can also set the `LOG_LEVEL` variable, otherwise it will default to `warn`. You can also specify an alternate `INPUT_FILE`, otherwise it will default to `day${DAY}/input.txt`.

```
make run DAY=1 [LOG_LEVEL=trace] [INPUT_FILE={path/to/input/file.txt}]
```

### make run-all

Builds the `advent-of-code-2022` executable and runs all solutions for all days. Optionally, you can also set the `LOG_LEVEL` variable, otherwise it will default to `warn`.

```
make run-all [LOG_LEVEL=trace]
```

### make install

Builds the `advent-of-code-2022` executable and installs it in the `$GOPATH/bin` directory.

### make uninstall

Deletes the `advent-of-code-2022` executable from the `$GOPATH/bin` directory.

## Day 0 Mock Puzzle

This repo also contains a day 0 mock puzzle to do some integration testing of all things int he common package prior to the day 1 puzzle unlocking. It reads in a list of users from the input file and calculates their ages from their birthdays. Part 1's answer is the name of the oldest user and part 2's answer is that person's current age.

```
$ make run DAY=0
go build
AOC_LOG_LEVEL=warn ./advent-of-code-2022 0
Part 1 Answer: Joe Schmoe
Part 2 Answer: 38
```