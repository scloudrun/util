//
// enc.go
//
// Distributed under terms of the MIT license.
//

package util

import (
	"crypto/md5"
	"fmt"
	"hash/crc32"
	"hash/crc64"
)

func Md5(str string) (res string) {
	aMd5 := md5.New()
	aMd5.Write([]byte(str))
	res = fmt.Sprintf("%x", aMd5.Sum(nil))
	return
}

func Crc32(str string) uint32 {
	table := crc32.MakeTable(crc32.IEEE)
	return crc32.Checksum([]byte(str), table)
}

func Crc64(str string) uint64 {
	table := crc64.MakeTable(crc64.ECMA)
	return crc64.Checksum([]byte(str), table)
}
