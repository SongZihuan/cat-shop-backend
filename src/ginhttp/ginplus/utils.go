package ginplus

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func getGoVersion() (int64, int64, int64, error) {
	version := runtime.Version()
	v := version

	if len(v) < 2 {
		return 0, 0, 0, fmt.Errorf("invalid version: %q", version)
	}

	if strings.HasPrefix(v, "go") {
		v = v[2:]
	}

	if len(v) == 0 {
		return 0, 0, 0, fmt.Errorf("invalid version: %q", version)
	}

	vLstStr := strings.Split(v, ".")
	vLst := make([]int64, len(vLstStr))

	for i, j := range vLstStr {
		var err error
		vLst[i], err = strconv.ParseInt(j, 10, 64)
		if err != nil {
			return 0, 0, 0, fmt.Errorf("invalid version: %q", version)
		}
	}

	if len(vLst) == 0 {
		return 0, 0, 0, fmt.Errorf("invalid version: %q", version)
	} else if len(vLst) == 1 {
		return vLst[0], 0, 0, nil
	} else if len(vLst) == 2 {
		return vLst[0], vLst[1], 0, nil
	} else if len(vLst) == 3 {
		return vLst[0], vLst[1], vLst[2], nil
	} else {
		return 0, 0, 0, fmt.Errorf("invalid version: %q", version)
	}
}

func getGoVersionMajor() (int64, error) {
	major, _, _, err := getGoVersion()
	if err != nil {
		return 0, err
	}

	return major, nil
}

func getGoVersionMinor() (int64, error) {
	_, minor, _, err := getGoVersion()
	if err != nil {
		return 0, err
	}

	return minor, nil
}

func getGoVersionPatch() (int64, error) {
	_, _, patch, err := getGoVersion()
	if err != nil {
		return 0, err
	}

	return patch, nil
}

var DebugPrintFunc func(format string, values ...interface{})
var DefaultWriter io.Writer = os.Stdout
var DefaultErrorWriter io.Writer = os.Stderr

func debugPrint(format string, values ...any) {
	if !IsDebugging() {
		return
	}

	if DebugPrintFunc != nil {
		DebugPrintFunc(format, values...)
		return
	}

	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	_, _ = fmt.Fprintf(DefaultWriter, "[GIN-PLUS-debug] "+format, values...)
}

func processURL(url string, defaultUrl ...string) string {
	if len(url) == 0 && len(defaultUrl) == 1 {
		url = defaultUrl[0]
	}

	url = strings.TrimSpace(url)

	if !strings.HasPrefix(url, "/") {
		url = "/" + url
	}

	url = strings.TrimRight(url, "/")

	return url
}
