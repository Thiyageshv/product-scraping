package utilities

import (
	"io"
	"strings"
	"github.com/OneOfOne/xxhash"
)

func XXHash(input string) int64 {
	h := xxhash.New64()
	r := strings.NewReader(input)
	io.Copy(h, r)
	return int64(h.Sum64())
}