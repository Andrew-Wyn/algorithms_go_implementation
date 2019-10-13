package main

// SquareGrid
type SquareGrid struct {
	width  int
	height int
}

func (s *SquareGrid) inbounds(p point) bool {
	return 0 <= p.x && p.x < s.width && 0 <= p.y && p.y < s.height
}

func (s *SquareGrid) neighbors(p point) []point {

	candidate := make([]point, 4)
	var out []point

	candidate[0] = point{p.x - 1, p.y}
	candidate[1] = point{p.x + 1, p.y}
	candidate[2] = point{p.x, p.y - 1}
	candidate[3] = point{p.x, p.y + 1}

	for i := 0; i < len(candidate); i++ {
		if s.inbounds(candidate[i]) {
			out = append(out, candidate[i])
		}
	}

	return out
}
