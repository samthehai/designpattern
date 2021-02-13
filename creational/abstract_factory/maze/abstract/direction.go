package abstract

// Direction enum type
type Direction int

const (
	// North direction
	North Direction = iota
	// East direction
	East
	// South direction
	South
	// West direction
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}
