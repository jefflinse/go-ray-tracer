package rt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMatrix(t *testing.T) {
	for size := 2; size <= 4; size++ {
		m := NewMatrix(size)
		assert.Len(t, m, size)
		for i := range m {
			assert.Len(t, m[i], size)
		}
	}

	m := Matrix{
		{-3, 5},
		{1, -2},
	}

	assert.Equal(t, -3.0, m[0][0])
	assert.Equal(t, 5.0, m[0][1])
	assert.Equal(t, 1.0, m[1][0])
	assert.Equal(t, -2.0, m[1][1])

	m = Matrix{
		{-3, 5, 0},
		{1, -2, -7},
		{0, 1, 1},
	}

	assert.Equal(t, -3.0, m[0][0])
	assert.Equal(t, -2.0, m[1][1])
	assert.Equal(t, 1.0, m[2][2])

	m = Matrix{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	}

	assert.Equal(t, 1.0, m[0][0])
	assert.Equal(t, 4.0, m[0][3])
	assert.Equal(t, 5.5, m[1][0])
	assert.Equal(t, 7.5, m[1][2])
	assert.Equal(t, 11.0, m[2][2])
	assert.Equal(t, 13.5, m[3][0])
	assert.Equal(t, 15.5, m[3][2])
}

func TestNewIdentityMatrix(t *testing.T) {
	m := NewIdentityMatrix()
	assert.Equal(t, 1.0, m[0][0])
	assert.Equal(t, 1.0, m[1][1])
	assert.Equal(t, 1.0, m[2][2])
	assert.Equal(t, 1.0, m[3][3])
}

func TestMatrix_Equals(t *testing.T) {
	m1 := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}
	m2 := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}
	assert.Equal(t, m1, m2)
	assert.True(t, m1.Equals(m2))
	m2 = Matrix{
		{2, 3, 4, 5},
		{6, 7, 8, 9},
		{8, 7, 6, 5},
		{4, 3, 2, 1},
	}
	assert.NotEqual(t, m1, m2)
	assert.False(t, m1.Equals(m2))
}

func TestMatrix_Multiply(t *testing.T) {
	m1 := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}
	m2 := Matrix{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	}
	m3 := Matrix{
		{20.0, 22.0, 50.0, 48.0},
		{44.0, 54.0, 114.0, 108.0},
		{40.0, 58.0, 110.0, 102.0},
		{16.0, 26.0, 46.0, 42.0},
	}
	assert.Equal(t, m3, m1.Multiply(m2))

	// multiply by identity
	m1 = Matrix{
		{0, 1, 2, 4},
		{1, 2, 4, 8},
		{2, 4, 8, 16},
		{4, 8, 16, 32},
	}
	m2 = NewIdentityMatrix()
	assert.Equal(t, m1, m1.Multiply(m2))
	assert.Equal(t, m1, m2.Multiply(m1))
}

func TestMatrix_MultiplyTuple(t *testing.T) {
	m := Matrix{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	}
	t1 := Tuple{1, 2, 3, 1}
	t2 := Tuple{18.0, 24.0, 33.0, 1.0}
	assert.Equal(t, t2, m.MultiplyTuple(t1))

	// multiple by identity
	m = NewIdentityMatrix()
	t1 = Tuple{1, 2, 3, 4}
	assert.Equal(t, t1, m.MultiplyTuple(t1))
}

func TestMatrix_Transpose(t *testing.T) {
	m1 := Matrix{
		{0, 9, 3, 0},
		{9, 8, 0, 8},
		{1, 8, 5, 3},
		{0, 0, 5, 8},
	}
	m2 := Matrix{
		{0, 9, 1, 0},
		{9, 8, 8, 0},
		{3, 0, 5, 5},
		{0, 8, 3, 8},
	}
	assert.Equal(t, m2, m1.Transpose())

	// transposing the identity matrix gives the identity matrix
	m1 = NewIdentityMatrix()
	assert.Equal(t, m1, m1.Transpose())
}

func TestMatrix_Determinant(t *testing.T) {
	//2x2
	m := Matrix{
		{1, 5},
		{-3, 2},
	}
	assert.Equal(t, 17.0, m.Determinant())

	// 3x3
	m = Matrix{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	}
	assert.Equal(t, 56.0, m.Cofactor(0, 0))
	assert.Equal(t, 12.0, m.Cofactor(0, 1))
	assert.Equal(t, -46.0, m.Cofactor(0, 2))
	assert.Equal(t, -196.0, m.Determinant())

	// 4x4
	m = Matrix{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	}
	assert.Equal(t, 690.0, m.Cofactor(0, 0))
	assert.Equal(t, 447.0, m.Cofactor(0, 1))
	assert.Equal(t, 210.0, m.Cofactor(0, 2))
	assert.Equal(t, 51.0, m.Cofactor(0, 3))
	assert.Equal(t, -4071.0, m.Determinant())
}

func TestMatrix_Submatrix(t *testing.T) {
	m1 := Matrix{
		{1, 5, 0},
		{-3, 2, 7},
		{0, 6, -3},
	}
	m2 := Matrix{
		{-3, 2},
		{0, 6},
	}
	assert.Equal(t, m2, m1.Submatrix(0, 2))

	m1 = Matrix{
		{-6, 1, 1, 6},
		{-8, 5, 8, 6},
		{-1, 0, 8, 2},
		{-7, 1, -1, 1},
	}
	m2 = Matrix{
		{-6, 1, 6},
		{-8, 8, 6},
		{-7, -1, 1},
	}
	assert.Equal(t, m2, m1.Submatrix(2, 1))
}

func TestMatrix_Minor(t *testing.T) {
	m1 := Matrix{
		{3, 5, 0},
		{2, -1, 7},
		{6, -1, 5},
	}
	m2 := m1.Submatrix(1, 0)
	assert.Equal(t, 25.0, m2.Determinant())
	assert.Equal(t, 25.0, m1.Minor(1, 0))
}

func TestMatrix_Cofactor(t *testing.T) {
	m := Matrix{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	}
	assert.Equal(t, -12.0, m.Minor(0, 0))
	assert.Equal(t, -12.0, m.Cofactor(0, 0))
	assert.Equal(t, 25.0, m.Minor(1, 0))
	assert.Equal(t, -25.0, m.Cofactor(1, 0))
}

func TestMatrix_IsInvertable(t *testing.T) {
	m := Matrix{
		{6, 4, 4, 4},
		{5, 5, 7, 6},
		{4, -9, 3, -7},
		{9, 1, 7, -6},
	}
	assert.Equal(t, -2120.0, m.Determinant())
	assert.True(t, m.IsInvertable())

	m = Matrix{
		{-4, 2, -2, -3},
		{9, 6, 2, 6},
		{0, -5, 1, -5},
		{0, 0, 0, 0},
	}
	assert.Equal(t, 0.0, m.Determinant())
	assert.False(t, m.IsInvertable())
}

func TestMatrix_Inverse(t *testing.T) {
	// panics on non-invertable matrix
	m1 := Matrix{
		{-4, 2, -2, -3},
		{9, 6, 2, 6},
		{0, -5, 1, -5},
		{0, 0, 0, 0},
	}
	assert.Panics(t, func() { m1.Inverse() })

	m1 = Matrix{
		{-5, 2, 6, -8},
		{1, -5, 1, 8},
		{7, 7, -6, -7},
		{1, -3, 7, 4},
	}
	m2 := m1.Inverse()
	assert.Equal(t, 532.0, m1.Determinant())
	assert.Equal(t, -160.0, m1.Cofactor(2, 3))
	assert.Equal(t, -160.0/532.0, m2[3][2])
	assert.Equal(t, 105.0, m1.Cofactor(3, 2))
	assert.Equal(t, 105.0/532.0, m2[2][3])
	m3 := Matrix{
		{.21805, .45113, .24060, -.04511},
		{-.80827, -1.45677, -.44361, .52068},
		{-.07895, -.22368, -.05263, .19737},
		{-.52256, -.81391, -.30075, .30639},
	}
	assert.True(t, m2.Equals(m3))

	m1 = Matrix{
		{8, -5, 9, 2},
		{7, 5, 6, 1},
		{-6, 0, 9, 6},
		{-3, 0, -9, -4},
	}
	m2 = Matrix{
		{-.15385, -.15385, -.28205, -.53846},
		{-.07692, .12308, .02564, .03077},
		{.35897, .35897, .43590, .92308},
		{-.69231, -.69231, -.76923, -1.92308},
	}
	assert.True(t, m1.Inverse().Equals(m2))

	m1 = Matrix{
		{9, 3, 0, 9},
		{-5, -2, -6, -3},
		{-4, 9, 6, 4},
		{-7, 6, 6, 2},
	}
	m2 = Matrix{
		{-.04074, -.07778, .14444, -.22222},
		{-.07778, .03333, .36667, -.33333},
		{-.02901, -.14630, -.10926, .12963},
		{.17778, .06667, -.26667, .33333},
	}
	assert.True(t, m1.Inverse().Equals(m2))

	m1 = Matrix{
		{3, -9, 7, 3},
		{3, -8, 2, -9},
		{-4, 4, 4, 1},
		{-6, 5, -1, 1},
	}
	m2 = Matrix{
		{8, 2, 2, 2},
		{3, -1, 7, 0},
		{7, 0, 5, 4},
		{6, -2, 0, 5},
	}
	m3 = m1.Multiply(m2)
	assert.True(t, m3.Multiply(m2.Inverse()).Equals(m1))
}
