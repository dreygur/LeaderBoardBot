// Package scoring holds the rules that turn activity into points and levels.
// It stays free of config and database access so it can be tested alone.
package scoring

import "math"

// levelBase is the points needed for level 1. Level n needs levelBase * n^2.
const levelBase = 100

// LevelFor returns the level a user holds with the given total points.
// Each level costs more than the last: 100, 400, 900, ... points.
func LevelFor(points int) int {
	if points < levelBase {
		return 0
	}
	return int(math.Sqrt(float64(points / levelBase)))
}
