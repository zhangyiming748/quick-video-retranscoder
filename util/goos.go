package util

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"log"
	"runtime"
)

const (
	INTEL = iota
	AMD
	Windows
	Linux
	Macintosh
)

func GenerateFFmpegParamsForCurrentSystem() string {
	if brand := getCPUVendor(); brand == INTEL {
		switch runtime.GOOS {
		case "darwin":
			log.Println("当前系统是mac intel平台")
			return "h264_videotoolbox"
		default:
			return "h264_qsv"
		}
	} else if brand == AMD {
		return "h264_amf"
	} else {
		return "h264_videotoolbox"
	}
}
func getCPUVendor() int {
	info, err := cpu.Info()
	if err != nil {
		fmt.Println("Error:", err)
		return -1
	}
	vendorID := info[0].VendorID
	log.Printf("%+v\n", info)
	switch vendorID {
	case "GenuineIntel":
		fmt.Println("当前CPU品牌是Intel")
		return INTEL
	case "AuthenticAMD":
		fmt.Println("当前CPU品牌是AMD")
		return AMD
	default:
		fmt.Println("未知CPU制造商")
		return -1
	}
}
func getOS() int {
	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("当前系统是Windows")
		return Windows
	case "linux":
		fmt.Println("当前系统是Linux")
		return Linux
	case "darwin":
		fmt.Println("当前系统是macOS")
		return Macintosh
	default:
		fmt.Println("未知操作系统")
		return -1
	}
}
