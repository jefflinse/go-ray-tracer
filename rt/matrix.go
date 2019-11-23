package rt

// A Matrix is a 4x4 matrix of numbers.
type Matrix [][]float64

// NewMatrix creates a new zero-filled Matrix.
func NewMatrix(size int) Matrix {
	m := make(Matrix, size)
	for i := range m {
		m[i] = make([]float64, size)
	}

	return m
}

// NewIdentityMatrix creates a new identity matrix.
func NewIdentityMatrix() Matrix {
	return Matrix{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

// Equals returns true of all of the elements in this matrix are equal to all of the elements in the other.
func (m Matrix) Equals(other Matrix) bool {
	size := len(m)
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			if !eq(m[r][c], other[r][c]) {
				return false
			}
		}
	}

	return true
}

// Multiply creates a new Matrix by multiplying this matrix with another.
func (m Matrix) Multiply(other Matrix) Matrix {
	size := len(m)
	new := NewMatrix(size)
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			for i := 0; i < size; i++ {
				new[r][c] += m[r][i] * other[i][c]
			}
		}
	}

	return new
}

// MultiplyTuple creates a new Tuple by multiplying this matrix with the given tuple.
func (m Matrix) MultiplyTuple(t Tuple) Tuple {
	size := len(m)
	new := NewTuple()
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			new[r] += m[r][c] * t[c]
		}
	}

	return new
}

// Transpose creates a new Matrix whose rows and columns are transposed from this one.
func (m Matrix) Transpose() Matrix {
	new := NewMatrix(len(m))
	for r := range m {
		for c := range m[r] {
			new[c][r] = m[r][c]
		}
	}

	return new
}

// Determinant computes the determinant of a matrix.
func (m Matrix) Determinant() float64 {
	if len(m) == 2 {
		return m[0][0]*m[1][1] - m[0][1]*m[1][0]
	}

	var d float64
	for c := 0; c < len(m); c++ {
		d += m[0][c] * m.Cofactor(0, c)
	}

	return d
}

// Submatrix creates a new matrix by removing the specifed row and column from this one.
func (m Matrix) Submatrix(row int, col int) Matrix {
	newLen := len(m) - 1
	new := NewMatrix(newLen)
	var nr, nc int
	for mr := range m {
		if mr != row {
			for mc := range m[mr] {
				if mc != col {
					new[nr][nc] = m[mr][mc]
					nc = (nc + 1) % newLen
				}
			}

			nr = (nr + 1) % newLen
		}
	}

	return new
}

// Minor computes the minor of a matrix at the specified row and column.
func (m Matrix) Minor(row int, col int) float64 {
	return m.Submatrix(row, col).Determinant()
}

// Cofactor computes the cofactor of a matrix at the specified row and column.
func (m Matrix) Cofactor(row int, col int) float64 {
	minor := m.Minor(row, col)
	if (row+col)%2 == 1 {
		minor = -minor
	}

	return minor
}

// IsInvertable returns true if the matrix is invertable.
func (m Matrix) IsInvertable() bool {
	return m.Determinant() != 0
}

// Inverse creates a new matrix representing the inverse of this one.
func (m Matrix) Inverse() Matrix {
	if !m.IsInvertable() {
		panic("attempted to invert non-invertable matrix")
	}

	size := len(m)
	new := NewMatrix(size)
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			cofactor := m.Cofactor(r, c)
			new[c][r] = cofactor / m.Determinant()
		}
	}

	return new
}
