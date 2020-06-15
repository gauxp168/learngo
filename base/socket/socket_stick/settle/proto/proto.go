package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

func Encode(message string) ([]byte, error) {
	// 读取消息长度， 转换成int32   （占4个字节）
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil,err
	}
	return pkg.Bytes(),nil
}

func Decode(reader *bufio.Reader) (string, error) {
	peek, err := reader.Peek(4)
	if err != nil {
		return "", err
	}
	buffer := bytes.NewBuffer(peek)
	var length int32
	err = binary.Read(buffer, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
