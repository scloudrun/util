//
// type_conv.go
//
// Distributed under terms of the MIT license.
//

package util

import (
	"fmt"
	"strconv"
)

const (
	digits     = "0123456789"
	uintbuflen = 20
)

func Int64(val interface{}) int64 {
	switch t := val.(type) {

	case int:
		return int64(t)
	case int8:
		return int64(t)
	case int16:
		return int64(t)
	case int32:
		return int64(t)
	case int64:
		return int64(t)

	case uint:
		return int64(t)
	case uint8:
		return int64(t)
	case uint16:
		return int64(t)
	case uint32:
		return int64(t)
	case uint64:
		return int64(t)

	case bool:
		if t == true {
			return int64(1)
		}
		return int64(0)

	case float32:
		return int64(t)
	case float64:
		return int64(t)
	default:
		i, _ := strconv.ParseInt(String(val), 10, 64)
		return i
	}

	panic("Invalid")

}

func Uint64(val interface{}) uint64 {
	switch t := val.(type) {

	case int:
		return uint64(t)
	case int8:
		return uint64(t)
	case int16:
		return uint64(t)
	case int32:
		return uint64(t)
	case int64:
		return uint64(t)

	case uint:
		return uint64(t)
	case uint8:
		return uint64(t)
	case uint16:
		return uint64(t)
	case uint32:
		return uint64(t)
	case uint64:
		return uint64(t)

	case float32:
		return uint64(t)
	case float64:
		return uint64(t)

	case bool:
		if t == true {
			return uint64(1)
		}
		return uint64(0)
	default:
		i, _ := strconv.ParseUint(String(val), 10, 64)
		return i
	}

	panic("Invalid")

}

func Float64(val interface{}) float64 {

	switch t := val.(type) {

	case int:
		return float64(t)
	case int8:
		return float64(t)
	case int16:
		return float64(t)
	case int32:
		return float64(t)
	case int64:
		return float64(t)

	case uint:
		return float64(t)
	case uint8:
		return float64(t)
	case uint16:
		return float64(t)
	case uint32:
		return float64(t)
	case uint64:
		return float64(t)

	case float32:
		return float64(t)
	case float64:
		return float64(t)

	case bool:
		if t == true {
			return float64(1)
		}
		return float64(0)
	case string:
		f, _ := strconv.ParseFloat(val.(string), 64)
		return f
	}

	panic("Invalid")
}

func Bytes(val interface{}) []byte {

	if val == nil {
		return []byte{}
	}

	switch t := val.(type) {
	case int:
		return int64ToBytes(int64(t))
	case int8:
		return int64ToBytes(int64(t))
	case int16:
		return int64ToBytes(int64(t))
	case int32:
		return int64ToBytes(int64(t))
	case int64:
		return int64ToBytes(int64(t))

	case uint:
		return uint64ToBytes(uint64(t))
	case uint8:
		return uint64ToBytes(uint64(t))
	case uint16:
		return uint64ToBytes(uint64(t))
	case uint32:
		return uint64ToBytes(uint64(t))
	case uint64:
		return uint64ToBytes(uint64(t))

	case float32:
		return float32ToBytes(t)
	case float64:
		return float64ToBytes(t)

	case complex128:
		return complex128ToBytes(t)
	case complex64:
		return complex128ToBytes(complex128(t))

	case bool:
		if t == true {
			return []byte("true")
		}
		return []byte("false")

	case string:
		return []byte(t)

	case []byte:
		return t

	default:
		return []byte(fmt.Sprintf("%v", val))
	}

	panic("Invalid")
}

func String(val interface{}) string {
	var buf []byte

	if val == nil {
		return ""
	}

	switch t := val.(type) {

	case int:
		buf = int64ToBytes(int64(t))
	case int8:
		buf = int64ToBytes(int64(t))
	case int16:
		buf = int64ToBytes(int64(t))
	case int32:
		buf = int64ToBytes(int64(t))
	case int64:
		buf = int64ToBytes(int64(t))

	case uint:
		buf = uint64ToBytes(uint64(t))
	case uint8:
		buf = uint64ToBytes(uint64(t))
	case uint16:
		buf = uint64ToBytes(uint64(t))
	case uint32:
		buf = uint64ToBytes(uint64(t))
	case uint64:
		buf = uint64ToBytes(uint64(t))

	case float32:
		buf = float32ToBytes(t)
	case float64:
		buf = float64ToBytes(t)

	case complex128:
		buf = complex128ToBytes(t)
	case complex64:
		buf = complex128ToBytes(complex128(t))

	case bool:
		if val.(bool) == true {
			return "true"
		}

		return "false"

	case string:
		return t

	case []byte:
		return string(t)

	default:
		return fmt.Sprintf("%v", val)
	}

	return string(buf)
}

func Bool(value interface{}) bool {
	b, _ := strconv.ParseBool(String(value))
	return b
}

func uint64ToBytes(v uint64) []byte {
	buf := make([]byte, uintbuflen)

	i := len(buf)

	for v >= 10 {
		i--
		buf[i] = digits[v%10]
		v = v / 10
	}

	i--
	buf[i] = digits[v%10]

	return buf[i:]
}

func int64ToBytes(v int64) []byte {
	negative := false

	if v < 0 {
		negative = true
		v = -v
	}

	uv := uint64(v)

	buf := uint64ToBytes(uv)

	if negative {
		buf2 := []byte{'-'}
		buf2 = append(buf2, buf...)
		return buf2
	}

	return buf
}

func float32ToBytes(v float32) []byte {
	slice := strconv.AppendFloat(nil, float64(v), 'g', -1, 32)
	return slice
}

func float64ToBytes(v float64) []byte {
	slice := strconv.AppendFloat(nil, v, 'g', -1, 64)
	return slice
}

func complex128ToBytes(v complex128) []byte {
	buf := []byte{'('}

	r := strconv.AppendFloat(buf, real(v), 'g', -1, 64)

	im := imag(v)
	if im >= 0 {
		buf = append(r, '+')
	} else {
		buf = r
	}

	i := strconv.AppendFloat(buf, im, 'g', -1, 64)

	buf = append(i, []byte{'i', ')'}...)

	return buf
}
