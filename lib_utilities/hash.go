package utilities

import (
	"io"
	"strings"
	"github.com/OneOfOne/xxhash"
	"hash/crc32"
	 b64 "encoding/base64"
)

func XXHash(input string) int64 {
	h := xxhash.New64()
	r := strings.NewReader(input)
	io.Copy(h, r)
	return int64(h.Sum64())
}

func Hash32(input string) int {
	crc32Uint32 := crc32.ChecksumIEEE([]byte(input))
	return int(crc32Uint32)
}

func B64Encode(input string) string {
	sEnc := b64.StdEncoding.EncodeToString([]byte(input))
	return sEnc
}

func B64Decode(input string) (string, error) {
	decode, err := b64.StdEncoding.DecodeString(input)
	return string(decode), err
}

func URLEncode(input string) string {
	uEnc := b64.URLEncoding.EncodeToString([]byte(input))
	return uEnc
}

func URLDecode(input string) (string, error) {
	decode, err := b64.URLEncoding.DecodeString(input)
	return string(decode), err
}