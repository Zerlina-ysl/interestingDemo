package base_struct

import (
	"bytes"
	"strings"
	"testing"
)

func Test_Buffer(t *testing.T) {
	var b bytes.Buffer
	for i := 0; i < 10000; i++ {
		b.WriteByte('a')
	}
	_ = b.String()
}
func Test_Builder(t *testing.T) {
	var b strings.Builder
	for i := 0; i < 10000; i++ {
		b.WriteByte('a')
	}
	_ = b.String()
}
