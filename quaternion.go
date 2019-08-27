package quaternion

/*
Exprimental code source
Quaternion array package
A Quaternion array is a struct with 4 columns defined as x y z w
it can be created simply using slice.
*/

type quaternion interface {
	ArrayDot(q, []float64)
}

type qMatrix struct {
	qMat []float64
}

//Dot product of a lists of arrays, returns a column array
func (qm *qMatrix) ArrayDot(q, []float64) {}

//Inverse of quaternion array q
func (qm *qMatrix) Inverse(q []float64) []float64 {
	return q * qm.qMat{-1, -1, -1, 1}
}

func (qm *qMatrix) Amplitude(q []float64) []float64 {}

//Normalize quaternion array q or array list to unit quaternions
func (qm *qMatrix) Norm([]float64) []float64 {}

//Rotate vector or array of vectors v by quaternion qm
//func (qm *qMatrix) Rotate(, []float64) []{}
//
func Mul() {}

//Exponential of a quaternion array
func (qm *qMatrix) Exp() {}

//Neprien logarithm of a quaternion array
func (qm *qMatrix) Ln() {}

//Real power of a quaternion array
func (qm *qMatrix) Pow() {}

//Rotate vector or array of vectors v by quaternion q
func Rotate(q []float64) {}

//
func (qm *qMatrix) Slerp([]float64) {}
