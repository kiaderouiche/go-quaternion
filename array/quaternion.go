package quaternion

/*
Quaternion array package
A Quaternion array is a struct with 4 columns defined as x y z w
it can be created simply using slice.
*/

import "math"

type qMatrix struct{
	qMat []float
}

//
func (qm *qMatrix) ArrayDot(){

}

//Inverse of quaternion array q
func (qm *qMatrix) Inverse() {

}


func (qm *qMatrix) Amplitude(){

}

//Normalize quaternion array q or array list to unit quaternions
func (qm *qMatrix) Norm(){

}

//
func Division(){

}

//
func Mul(){

}

//Exponential of a quaternion array
func (qm *qMatrix) Exp(){

}

//Neprien logarithm of a quaternion array
func (qm *qMatrix) Ln(){

}

//Real power of a quaternion array
func (qm *qMatrix) Pow(){

}

//Rotate vector or array of vectors v by quaternion q
func Rotate(q []float64) {

}

//
func Toaxisangle(){

}