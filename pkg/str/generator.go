package str

import "math/rand"

var chars []int32

func init() {
	add := func(start, stop int32) {
		for i := start; i <= stop; i++ {
			chars = append(chars, i)
		}
	}
	add('0', '9')
	add('a', 'z')
	add('A', 'Z')
}

func Random(length int) string {
	bytes := make([]rune, length)
	for i := 0; i < length; i++ {
		bytes[i] = chars[rand.Int()%len(chars)]
	}
	result := string(bytes)
	println(string(chars))
	return result
}
