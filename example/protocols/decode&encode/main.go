package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func main() {
	// 字符串转为字节
	str:= "test"
	fmt.Println([]byte(str))
	// 整数类型转换为字节
	m := 20
	fmt.Println(IntToBytes(m))
	// 实数类型转换为字节
	f:= 1.345775
	fmt.Println(Float64ToByte(f))
}

func IntToBytes(n int) []byte {
	data := int64(n) // 数据类型转换
	buffer := bytes.NewBuffer([]byte{}) // 字节集合
	binary.Write(buffer, binary.BigEndian, data) // 按照二进制写入字节
	return buffer.Bytes() // 返回字节集合
}

func BytesToInt(bs []byte) int {
	buffer := bytes.NewBuffer(bs)
	var data int64
	binary.Read(buffer, binary.BigEndian, &data)
	return int(data)
}

// 浮点数转化为字节
func Float32ToByte(n float32) []byte {
	bits := math.Float32bits(n)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

func Float64ToByte(n float64) []byte {
	bits := math.Float64bits(n)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func ByteToFloat32(bs []byte) float32 {
	bits := binary.LittleEndian.Uint32(bs)
	return math.Float32frombits(bits)
}

func ByteToFloat64(bs []byte) float64 {
	bits := binary.LittleEndian.Uint64(bs)
	return math.Float64frombits(bits)
}

