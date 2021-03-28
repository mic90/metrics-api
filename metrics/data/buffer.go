package data

const bufferIncreaseStep = 1000

type Buffer struct {
	points []Point
}

func NewBuffer() *Buffer {
	return &Buffer{
		make([]Point, 0, bufferIncreaseStep),
	}
}

func (b *Buffer) Grow(increase int) {
	old := make([]Point, 0, len(b.points))
	copy(old, b.points)
	b.points = make([]Point, 0, len(b.points)+increase)
	copy(b.points, old)
}

func (b *Buffer) Cap() int {
	return cap(b.points)
}

func (b *Buffer) Data() []Point {
	return b.points
}

func (b *Buffer) Add(value Point) {
	b.points = append(b.points, value)

	// grow internal buffer if we have reached it's cap
	if len(b.points) == cap(b.points) {
		b.Grow(bufferIncreaseStep)
	}
}
