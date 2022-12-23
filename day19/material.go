package day19

type material int

const (
	none     material = 0
	ore      material = 1
	clay     material = 2
	obsidian material = 3
	geode    material = 4
)

func (material material) String() string {
	switch material {
	case ore:
		return "ore"
	case clay:
		return "clay"
	case obsidian:
		return "obsidian"
	case geode:
		return "geode"
	default:
		return "unknown"
	}
}
