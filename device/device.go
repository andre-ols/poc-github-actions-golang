package device

import (
	"fmt"
	"strconv"
	"strings"
)

type InfoDevice struct {
	Vid int
	Pid int
}
type Device struct {
	Class       string
	ClassGuid   string
	Description string
	HardwareID  string
	Info        []InfoDevice
	Name        string
	Status      string
}

// Function responsible for receiving a string containing information about the device and parsing this information
func HardwareIdParser(input string) ([]InfoDevice, error) {
	// Remove foreign braces from string
	input = input[1 : len(input)-1]

	// Break string into more substrings corresponding to each USB device
	devices := strings.Split(input, ", ")

	results := make([]InfoDevice, len(devices))

	// For each device, extract VID, PID and add them to the corresponding map
	for i, device := range devices {
		parts := strings.Split(device, "\\")

		data := strings.Split(parts[1], "\u0026")

		vidPart := strings.TrimPrefix(data[0], "VID_")
		pidPart := strings.TrimPrefix(data[1], "PID_")

		vid, err := strconv.ParseInt(vidPart, 16, 32)
		if err != nil {
			return nil, fmt.Errorf("failed to parse VID: %w", err)
		}
		pid, err := strconv.ParseInt(pidPart, 16, 32)
		if err != nil {
			return nil, fmt.Errorf("failed to parse PID: %w", err)
		}

		result := InfoDevice{
			Vid: int(vid),
			Pid: int(pid),
		}
		results[i] = result
	}

	return results, nil
}
