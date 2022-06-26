package utils

import (
	"github.com/bwmarrin/snowflake"
)

var chars []byte = []byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
	'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
	'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
var charsLen = int64(len(chars))
var node *snowflake.Node

func init() {
	tmp, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	node = tmp
}

func GenerateID() (string, error) {
	id := node.Generate()
	id_num := id.Int64()
	res := make([]byte, 0, 6)
	for i := 0; i < 6; i++ {
		cur := id_num % charsLen
		id_num /= charsLen
		res = append(res, chars[cur])
	}
	return string(res), nil
}
