package twodimensionpacking

type Point struct {
	X int
	Y int
}

func PointNew(x, y int) *Point {
	return &Point{x, y}
}

type Line struct {
	P1, P2 *Point
}

func LineNew(p1, p2 *Point) *Line {
	return &Line{p1, p2}
}

// VerticalIntersect checks if the two horizontal lines intersect in the vertical
//  direction, and returns their vertical distance (l - l1).
func (l *Line) VerticalIntersect(l1 *Line) (isInter bool, dist int) {
	isInter = l.P1.X < l1.P2.X && l1.P1.X < l.P2.X
	dist = l.P1.Y - l1.P1.Y
	return
}

// HorizontalIntersect checks if the two vertical lines intersect in the horizontal
//  direction, and returns their vertical distance (l - l1).
func (l *Line) HorizontalIntersect(l1 *Line) (isInter bool, dist int) {
	isInter = l.P1.Y < l1.P2.Y && l1.P1.Y < l.P2.Y
	dist = l.P1.X - l1.P1.X
	return
}

type Size struct {
	Height int
	Width  int
}

func SizeNew(height, width int) *Size {
	return &Size{height, width}
}

func (s *Size) Area() int {
	return s.Height * s.Width
}

// CanHold returns if r can hold r1 in w and h.
func (s *Size) CanHold(s1 *Size) bool {
	return s.Width >= s1.Width && s.Height >= s1.Height
}

type Rectangle struct {
	*Size
	// Sp is the bottom left point
	Sp *Point
}

func RectangleNew(sz *Size, sp *Point) *Rectangle {
	return &Rectangle{sz, sp}
}

// BottomLine returns the bottom line of r.
func (r *Rectangle) BottomLine() *Line {
	bl := *r.Sp
	br := PointNew(r.Sp.X+r.Width, r.Sp.Y)
	return LineNew(&bl, br)
}

// TopLine returns the top line of r.
func (r *Rectangle) TopLine() *Line {
	tl := PointNew(r.Sp.X, r.Sp.Y+r.Height)
	tr := PointNew(r.Sp.X+r.Width, r.Sp.Y+r.Height)
	return LineNew(tl, tr)
}

// LeftLine returns the line on the left side of r.
func (r *Rectangle) LeftLine() *Line {
	tl := PointNew(r.Sp.X, r.Sp.Y+r.Height)
	bl := *r.Sp
	return LineNew(&bl, tl)
}

// RightLine returns the top line of r.
func (r *Rectangle) RightLine() *Line {
	br := PointNew(r.Sp.X+r.Width, r.Sp.Y)
	tr := PointNew(r.Sp.X+r.Width, r.Sp.Y+r.Height)
	return LineNew(br, tr)
}

// Intersect checks if two rectangle interesect.
func (r *Rectangle) Intersect(r1 *Rectangle) bool {
	isInterX := r.Sp.X < r1.Sp.X+r1.Width && r1.Sp.X < r.Sp.X+r.Width
	isInterY := r.Sp.Y < r1.Sp.Y+r1.Height && r1.Sp.Y < r.Sp.Y+r.Height
	return isInterX && isInterY
}

type Box struct {
	*Rectangle
	StoredItem []*Rectangle
}

// BoxNew returns a box with given size and the starting point at (0,0).
func BoxNew(height, width int) *Box {
	return &Box{
		Rectangle: &Rectangle{Size: SizeNew(height, width), Sp: PointNew(0, 0)},
	}
}

type Item struct {
	Size
}

func ItemNew(height, width int) *Item {
	return &Item{Size: *SizeNew(height, width)}
}
