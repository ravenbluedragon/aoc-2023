package grid

// Point2D stores the x and y coordinates of a point
type Point2D struct {
	X, Y int
}

// P2D returns a 2D point
func P2(x, y int) Point2D {
	return Point2D{x, y}
}

// Point3D stores the x, y and z coordinates of a point
type Point3D struct {
	X, Y, Z int
}

// P3D returns a 3D point
func P3(x, y, z int) Point3D {
	return Point3D{x, y, z}
}
