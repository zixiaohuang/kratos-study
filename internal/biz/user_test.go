package biz

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

// go test -v 可以输出
func TestHashPassword(t *testing.T) {
	s := hashPassword("abc")
	spew.Dump(s)
}

// T 测试 B brenchmark 性能测试
func TestVerifyPassword(t *testing.T) {
	a := assert.New(t)
	a.False(verifyPassword("$2a$10$9EUnbbMEb7WoCYercfKYLOj1o1GEEhru07ynotk0/HMac7VgNeFtm", "abc"))
	a.False(verifyPassword("$2a$10$9EUnbbMEb7WoCYercfKYLOj1o1GEEhru07ynotk0/HMac7VgNeFtm", "abc1"))
}