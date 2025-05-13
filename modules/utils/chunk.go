package utils

import "math"

func CalculateChunkNumber(fileSize int64) int16 {
	mb := fileSize / 1024 / 1024
	chunks := math.Ceil(float64(mb) / 5)
	return int16(chunks)
}
