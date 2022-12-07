package day07

import (
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

const totalDiskSpace = 70000000
const freeSpaceRequired = 30000000

func Solve(puzzle *common.Puzzle) common.Answer {

	input := common.ReadFile(puzzle.InputFile)
	lines := common.Split(input, "\n")

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(lines),
		Part2: solvePart2(lines),
	}
}

func solvePart1(lines []string) string {
	log.Debug("Solving part 1.")

	root := parseTerminalOutput(lines)
	directories := root.FindDirectoriesOfSizeLessThanOrEqualTo(100000)

	sum := 0
	for _, directory := range directories {
		sum += directory.Size()
	}

	log.Debug("Part 1 solved.")
	return strconv.Itoa(sum)
}

func solvePart2(lines []string) string {
	log.Debug("Solving part 2.")

	root := parseTerminalOutput(lines)
	usedDiskSpace := root.Size()
	freeSpace := totalDiskSpace - usedDiskSpace
	spaceToFreeUp := freeSpaceRequired - freeSpace

	deletionCandidates := root.FindDirectoriesOfSizeGreaterThanOrEqualTo(spaceToFreeUp)

	directoryToDelete := deletionCandidates[0]

	for _, deletionCandidate := range deletionCandidates[1:] {
		if deletionCandidate.Size() < directoryToDelete.Size() {
			directoryToDelete = deletionCandidate
		}
	}

	log.Debug("Part 2 solved.")
	return strconv.Itoa(directoryToDelete.Size())
}

func parseTerminalOutput(lines []string) *directory {

	root := NewDirectory("/", nil)
	var cwd *directory = nil

	commandRegexp := regexp.MustCompile(`^\$ (\S*)( \S*)?$`)
	outputRegexp := regexp.MustCompile(`^(\d*|dir) ([\S\.]*)$`)

	for index := 0; index < len(lines); index++ {
		line := lines[index]

		if line == "" {
			continue
		}

		log.Debugf("Processing command \"%s\".", line)

		if !commandRegexp.MatchString(line) {
			log.Panicf("Line \"%s\" is not a command.", line)
		}

		groups := commandRegexp.FindStringSubmatch(line)
		command := groups[1]
		parameter := strings.Trim(groups[2], " ")

		switch command {
		case "cd":
			log.Debug("Changing directory.")

			if parameter == "/" {
				cwd = root
			} else if parameter == ".." {
				cwd = cwd.Parent
			} else {
				var targetDirectory *directory = nil

				for _, subdirectory := range cwd.Subdirectories {
					if subdirectory.Name == parameter {
						targetDirectory = subdirectory
						break
					}
				}

				if targetDirectory == nil {
					log.Panicf("Directory \"%s\" does not contain a subdirectory \"%s\".", cwd.Name, parameter)
				} else {
					cwd = targetDirectory
				}
			}

			log.Debugf("Current Working Directory is \"%s\".", cwd.Name)

		case "ls":
			log.Debug("Listing contents of Current Working Directory.")

			index++
			for ; index < len(lines); index++ {
				line := lines[index]

				if outputRegexp.MatchString(line) {
					groups := outputRegexp.FindStringSubmatch(line)

					if groups[1] == "dir" {
						log.Debugf("Directory \"%s\" found.", groups[2])
						newDirectory := NewDirectory(groups[2], cwd)
						cwd.Subdirectories = append(cwd.Subdirectories, newDirectory)
					} else {
						log.Debugf("File \"%s\" found.", groups[2])
						size, err := strconv.Atoi(groups[1])
						common.Check(err)
						newFile := NewFile(groups[2], size)
						cwd.Files = append(cwd.Files, newFile)
					}

				} else {
					break
				}
			}

			index--
		default:
			log.Panicf("Unsupported command \"%s\".", command)
		}
	}

	return root
}
