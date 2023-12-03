package grid

// Rect2D represents a rectangle in 2D space
type Rect2D struct {
	Min, Max Point2D
}

// R2 returns a 2D rectangle
func R2(min, max Point2D) Rect2D {
	return Rect2D{min, max}
}

// R2c returns a 2D rectangle from coordinates
func R2c(x1, y1, x2, y2 int) Rect2D {
	return Rect2D{Point2D{x1, y1}, Point2D{x2, y2}}
}

// R2dim returns a 2D rectangle from coordinates and dimensions
func R2dim(x, y, w, h int) Rect2D {
	return Rect2D{Point2D{x, y}, Point2D{x + w, y + h}}
}

// Rect3D represents a rectangle in 3D space
type Rect3D struct {
	Min, Max Point3D
}

// R3 returns a 3D rectangle
func R3(min, max Point3D) Rect3D {
	return Rect3D{min, max}
}

// R3c returns a 3D rectangle from coordinates
func R3c(x1, y1, z1, x2, y2, z2 int) Rect3D {
	return Rect3D{Point3D{x1, y1, z1}, Point3D{x2, y2, z2}}
}

// R3dim returns a 3D rectangle from coordinates and dimensions
func R3dim(x, y, z, w, h, d int) Rect3D {
	return Rect3D{Point3D{x, y, z}, Point3D{x + w, y + h, z + d}}
}
