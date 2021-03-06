package huawei_routers_telemetry

import (
	"github.com/influxdata/telegraf/testutil"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandleTelemetryHuawei(t *testing.T) {
	h := &HuaweiRoutersTelemetry{Log: testutil.Logger{}, ServicePort: "5600"}
	acc := &testutil.Accumulator{}
	err := h.Start(acc)
	if err != nil {
		require.Error(t, err)
	}

	packet := []byte{0x00, 0x00, 0x01, 0xe2, 0x5e, 0x44, 0x60, 0x80, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x0b, 0x48, 0x75, 0x61, 0x77, 0x65, 0x69, 0x54, 0x65, 0x73, 0x74, 0x31, 0x12, 0x1e, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x6e, 0x64, 0x6f, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x61, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x63, 0x69, 0x6f, 0x6e, 0x1a, 0x21, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x2d, 0x64, 0x65, 0x76, 0x6d, 0x3a, 0x64, 0x65, 0x76, 0x6d, 0x2f, 0x63, 0x70, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x2f, 0x63, 0x70, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x20, 0xe9, 0xcf, 0x02, 0x28, 0x87, 0xab, 0xa2, 0xd9, 0x83, 0x2e, 0x30, 0xbf, 0xab, 0xa2, 0xd9, 0x83, 0x2e, 0x3a, 0xda, 0x02, 0x0a, 0x1d, 0x08, 0x87, 0xab, 0xa2, 0xd9, 0x83, 0x2e, 0x5a, 0x14, 0x2a, 0x12, 0x0a, 0x10, 0x22, 0x01, 0x39, 0x08, 0x81, 0x80, 0xa4, 0x08, 0x28, 0x0a, 0x18, 0x5a, 0x30, 0x4b, 0x10, 0x08, 0x0a, 0x1e, 0x08, 0x87, 0xab, 0xa2, 0xd9, 0x83, 0x2e, 0x5a, 0x15, 0x2a, 0x13, 0x0a, 0x11, 0x22, 0x02, 0x31, 0x30, 0x08, 0x81, 0x80, 0xa8, 0x08, 0x28, 0x06, 0x18, 0x5a, 0x30, 0x4b, 0x10, 0x08, 0x0a, 0x1e, 0x08, 0x87, 0xab, 0xa2, 0xd9, 0x83, 0x2e, 0x5a, 0x15, 0x2a, 0x13, 0x0a, 0x11, 0x22, 0x02, 0x31, 0x31, 0x08, 0x81, 0x80, 0xac, 0x08, 0x28, 0x03, 0x18, 0x5a, 0x30, 0x4b, 0x10, 0x08, 0x0a, 0x1e, 0x08, 0x87, 0xab, 0xa2, 0xd9, 0x83, 0x2e, 0x5a, 0x15, 0x2a, 0x13, 0x0a, 0x11, 0x22, 0x02, 0x31, 0x32, 0x08, 0x81, 0x80, 0xb0, 0x08, 0x28, 0x02, 0x18, 0x5a, 0x30, 0x4b, 0x10, 0x08, 0x0a, 0x1e, 0x08, 0x87, 0xab, 0xa2, 0xd9, 0x83, 0x2e, 0x5a, 0x15, 0x2a, 0x13, 0x0a, 0x11, 0x22, 0x02, 0x31, 0x33, 0x08, 0x81, 0x80, 0xb4, 0x08, 0x28, 0x03, 0x18, 0x5a, 0x30, 0x4b, 0x10, 0x08, 0x0a, 0x1e, 0x08, 0x87, 0xab, 0xa2, 0xd9, 0x83, 0x2e, 0x5a, 0x15, 0x2a, 0x13, 0x0a, 0x11, 0x22, 0x02, 0x31, 0x34, 0x08, 0x81, 0x80, 0xb8, 0x08, 0x28, 0x03, 0x18, 0x5a, 0x30, 0x4b, 0x10, 0x08, 0x0a, 0x1d, 0x08, 0x87, 0xab, 0xa2, 0xd9, 0x83, 0x2e, 0x5a, 0x14, 0x2a, 0x12, 0x0a, 0x10, 0x22, 0x01, 0x35, 0x08, 0x81, 0x80, 0x94, 0x08, 0x28, 0x12, 0x18, 0x5a, 0x30, 0x4b, 0x10, 0x08, 0x0a, 0x1d, 0x08, 0x87, 0xab, 0xa2, 0xd9, 0x83, 0x2e, 0x5a, 0x14, 0x2a, 0x12, 0x0a, 0x10, 0x22, 0x01, 0x34, 0x08, 0x81, 0x80, 0x90, 0x08, 0x28, 0x10, 0x18, 0x5a, 0x30, 0x4b, 0x10, 0x08, 0x0a, 0x1d, 0x08, 0x87, 0xab, 0xa2, 0xd9, 0x83, 0x2e, 0x5a, 0x14, 0x2a, 0x12, 0x0a, 0x10, 0x22, 0x01, 0x31, 0x08, 0x81, 0x80, 0x84, 0x08, 0x28, 0x2a, 0x18, 0x5a, 0x30, 0x4b, 0x10, 0x08, 0x0a, 0x1d, 0x08, 0x87, 0xab, 0xa2, 0xd9, 0x83, 0x2e, 0x5a, 0x14, 0x2a, 0x12, 0x0a, 0x10, 0x22, 0x01, 0x37, 0x08, 0x81, 0x80, 0x9c, 0x08, 0x28, 0x29, 0x18, 0x5a, 0x30, 0x4b, 0x10, 0x08, 0x0a, 0x1d, 0x08, 0x87, 0xab, 0xa2, 0xd9, 0x83, 0x2e, 0x5a, 0x14, 0x2a, 0x12, 0x0a, 0x10, 0x22, 0x01, 0x38, 0x08, 0x81, 0x80, 0xa0, 0x08, 0x28, 0x0e, 0x18, 0x5a, 0x30, 0x4b, 0x10, 0x08, 0x40, 0x87, 0xab, 0xa2, 0xd9, 0x83, 0x2e, 0x48, 0x90, 0x4e, 0x52, 0x02, 0x4f, 0x4b, 0x5a, 0x05, 0x4e, 0x45, 0x34, 0x30, 0x45, 0x60, 0x00}

	grouper := HuaweiTelemetryDecoder(packet)
	acc.AddMetric(grouper.Metrics()[0])

	require.Empty(t, acc.Errors)

	tags := map[string]string{"path": "huawei-devm:devm/cpuInfos/cpuInfo", "position": "\"9\"", "source": "HuaweiTest1", "subscription": "ProbandoTelemetriaSubscripcion"}
	fields := map[string]interface{}{"entIndex": uint32(17367041), "interval": uint32(8), "ovloadThreshold": uint32(90), "systemCpuUsage": uint32(10), "unovloadThreshold": uint32(75)}
	acc.AssertDoesNotContainsTaggedFields(t, "huawei-devm:devm/cpuInfos/cpuInfo", fields, tags)
}
