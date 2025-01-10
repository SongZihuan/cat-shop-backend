package flagparser

import (
	"fmt"
	"io"
	"strings"
	"time"
)

var data flagData

func Help() bool {
	return data.Help()
}

func FprintUseage(writer io.Writer) (int, error) {
	return data.FprintUseage(writer)
}

func PrintUseage() (int, error) {
	return data.PrintUseage()
}

func FprintVersion(writer io.Writer) (int, error) {
	return data.FprintVersion(writer)
}

func PrintVersion() (int, error) {
	return data.PrintVersion()
}

func FprintLicense(writer io.Writer) (int, error) {
	return data.FprintLicense(writer)
}

func PrintLicense() (int, error) {
	return data.PrintLicense()
}

func Version() bool {
	return data.Version()
}

func License() bool {
	return data.License()
}

func NotRunMode() bool {
	return Help() || Version() || License()
}

func NotRunModeOption() string {
	if !NotRunMode() {
		return ""
	}

	var result strings.Builder

	if data.HelpData {
		result.WriteString(fmt.Sprintf("%s%s, ", OptionPrefix, data.HelpName))
	}

	if data.VersionData {
		result.WriteString(fmt.Sprintf("%s%s, ", OptionPrefix, data.VersionName))
	}

	if data.LicenseData {
		result.WriteString(fmt.Sprintf("%s%s, ", OptionPrefix, data.LicenseName))
	}

	return strings.TrimSuffix(result.String(), ", ")
}

func ConfigFile() string {
	return data.ConfigFile()
}

func WaitSec() time.Duration {
	return time.Second * time.Duration(data.Wait())
}
