package day14

import "golang.org/x/exp/slices"

type Map struct {
	locations []Location
	current   Location
	Sand      int
	bottom    int
}

func (m *Map) AddLocation(loc Location) {
	m.locations = append(m.locations, loc)
}

func (m *Map) SetCurrentLocation(loc Location) {
	m.current = loc
	m.AddLocation(loc)

	if loc.Y > m.bottom {
		m.bottom = loc.Y
	}
}

func (m *Map) BuildTo(loc Location) {
	if loc.X == m.current.X && m.current.Y == loc.Y {
		return
	}

	stepx := m.current.X
	stepy := m.current.Y

	if loc.X > stepx {
		stepx++
	} else if loc.X < stepx {
		stepx--
	}

	if loc.Y > stepy {
		stepy++
	} else if loc.Y < stepy {
		stepy--
	}

	m.SetCurrentLocation(Location{X: stepx, Y: stepy})
	m.BuildTo(loc)
}

func (m *Map) FillWithSand(start Location, hasFloor bool) {
	m.Sand = 0

	for loc := start; loc.Y < m.bottom+2; {
		if hasFloor {
			if loc.Y == m.bottom + 1 {
				goto stopped
			}
		}
		
		if !slices.Contains(m.locations, Location{X: loc.X, Y: loc.Y + 1}) {
			loc = Location{X: loc.X, Y: loc.Y + 1}
			continue
		} else if !slices.Contains(m.locations, Location{X: loc.X - 1, Y: loc.Y + 1}) {
			loc = Location{X: loc.X - 1, Y: loc.Y + 1}
			continue
		} else if !slices.Contains(m.locations, Location{X: loc.X + 1, Y: loc.Y + 1}) {
			loc = Location{X: loc.X + 1, Y: loc.Y + 1}
			continue
		}		

		stopped:
		m.AddLocation(loc)
		m.Sand++
		if loc.X == 500 && loc.Y == 0 {
			break
		}
		loc = start
	}
}
