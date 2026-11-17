package scoring

import "testing"

func TestLevelFor(t *testing.T) {
	cases := []struct {
		points int
		want   int
	}{
		{-50, 0},
		{0, 0},
		{99, 0},
		{100, 1},
		{399, 1},
		{400, 2},
		{899, 2},
		{900, 3},
		{1_000_000, 100},
	}

	for _, c := range cases {
		if got := LevelFor(c.points); got != c.want {
			t.Errorf("LevelFor(%d) = %d, want %d", c.points, got, c.want)
		}
	}
}
