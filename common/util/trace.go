package util

import (
	"context"
	"encoding/binary"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

// GenerateSpanId 将十进制的 int64 转换为十六进制
// 变成 16位的字符串
func GenerateSpanId(addr string) string {
	strAddr := strings.Split(addr, ":")
	ip, _ := ipToInt32(strAddr[0])
	times := uint64(time.Now().Unix())
	spanId := ((uint64(ip) ^ times) << 32) | uint64(rand.Int31())
	return strconv.FormatUint(spanId, 16)
}

func ipToInt32(ip string) (uint32, error) {
	result, err := net.ResolveIPAddr("ip", ip)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(result.IP.To4()), nil
}

func GetTraceIdFromContext(ctx context.Context) (traceId, spanId, parentId string) {
	if ctx.Value("traceId") != nil {
		traceId = ctx.Value("traceId").(string)
	}
	if ctx.Value("spanId") != nil {
		spanId = ctx.Value("spanId").(string)
	}
	if ctx.Value("parentId") != nil {
		parentId = ctx.Value("parentId").(string)
	}
	return
}
