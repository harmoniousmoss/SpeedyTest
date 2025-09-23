package handlers

import (
	"fmt"
	"time"

	"github.com/showwin/speedtest-go/speedtest"
)

// RunSpeedTest runs the download, upload, and ping tests on the provided server
func RunSpeedTest(s *speedtest.Server) (*SpeedTestResult, error) {
	// Perform ping test with a callback for logging latency
	s.PingTest(func(latency time.Duration) {
		fmt.Printf("Ping: %v\n", latency)
	})

	// Perform download and upload tests
	err := s.DownloadTest()
	if err != nil {
		return nil, fmt.Errorf("failed to perform download test")
	}

	err = s.UploadTest()
	if err != nil {
		return nil, fmt.Errorf("failed to perform upload test")
	}

	// Create result structure with proper unit conversion
	// The speedtest-go library returns speeds in Bps (bytes per second)
	// Convert to Mbps: divide by 1,000,000 and multiply by 8 (bits/bytes)
	downloadMbps := (s.DLSpeed * 8) / 1000000
	uploadMbps := (s.ULSpeed * 8) / 1000000

	return &SpeedTestResult{
		DownloadSpeed: fmt.Sprintf("%.2f Mbps", downloadMbps),
		UploadSpeed:   fmt.Sprintf("%.2f Mbps", uploadMbps),
	}, nil
}
