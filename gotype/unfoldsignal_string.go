// Code generated by "stringer -type=unfoldSignal"; DO NOT EDIT

package gotype

import "fmt"

const _unfoldSignal_name = "sigObjectStartsigObjectFinishedsigObjectKeysigArrayStartsigArrayFinishedsigNilsigBoolsigStringsigInt8sigInt16sigInt32sigInt64sigIntsigBytesigUint8sigUint16sigUint32sigUint64sigUintsigFloat32sigFloat64"

var _unfoldSignal_index = [...]uint8{0, 14, 31, 43, 56, 72, 78, 85, 94, 101, 109, 117, 125, 131, 138, 146, 155, 164, 173, 180, 190, 200}

func (i unfoldSignal) String() string {
	if i >= unfoldSignal(len(_unfoldSignal_index)-1) {
		return fmt.Sprintf("unfoldSignal(%d)", i)
	}
	return _unfoldSignal_name[_unfoldSignal_index[i]:_unfoldSignal_index[i+1]]
}
