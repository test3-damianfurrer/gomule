/*
 * Copyright (C) 2013 Deepin, Inc.
 *               2013 Leslie Zhai <zhaixiang@linuxdeepin.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package emule

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func ByteToInt32(data []byte) (ret int32) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &ret)
	return
}

func ByteToUint32(data []byte) (ret uint32) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &ret)
	return
}

func ByteToInt16(data []byte) (ret int16) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &ret)
	return
}
func ByteToUint16(data []byte) (ret uint16) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &ret)
	return
}

func ByteToFloat32(data []byte) (ret float32) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &ret)
	return
}
func Float32ToByte(data float32) (ret []byte) {
	ret = []byte{}
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, data)
	ret = buf.Bytes()
	return
}

func Int16ToByte(data int16) (ret []byte) {
	ret = []byte{}
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, data)
	ret = buf.Bytes()
	return
}
func Uint16ToByte(data uint16) (ret []byte) {
	ret = []byte{}
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, data)
	ret = buf.Bytes()
	return
}

func Int32ToByte(data int32) (ret []byte) {
	ret = []byte{}
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, data)
	ret = buf.Bytes()
	return
}
func Uint32ToByte(data uint32) (ret []byte) {
	ret = []byte{}
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, data)
	ret = buf.Bytes()
	return
}

func HighId(addr string) (ret uint32) {
	ret = 0
	var ip1, ip2, ip3, ip4, port uint32
	fmt.Sscanf(addr, "%d.%d.%d.%d:%d", &ip1, &ip2, &ip3, &ip4, &port)
	ret = ip1 + uint32(math.Pow(2, 8))*ip2 + uint32(math.Pow(2, 16))*ip3 +
		uint32(math.Pow(2, 24))*ip4
	return
}
