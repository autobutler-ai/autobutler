//go:build windows

package util

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func getAvailableSpaceInBytes(fileDir string) uint64 {
	var freeBytesAvailable uint64
	var totalNumberOfBytes uint64
	var totalNumberOfFreeBytes uint64

	err := windows.GetDiskFreeSpaceEx(
		windows.StringToUTF16Ptr(fileDir),
		&freeBytesAvailable,
		&totalNumberOfBytes,
		&totalNumberOfFreeBytes,
	)
	if err != nil {
		fmt.Printf("Error getting disk space: %v\n", err)
		return 0
	}
	return freeBytesAvailable
}
