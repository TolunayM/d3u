package tools

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

var home, _ = os.UserHomeDir()
var d3uHomeDLSS = home + "\\d3u\\dlss"

func DownloadDLSS(version string) {

	downloadLink := "https://github.com/TolunayM/dlss-repo/releases/download/" + version + "/nvngx_dlss.dll"
	customDLL := d3uHomeDLSS + "\\nvngx_dlss_" + version + ".dll"

	fmt.Println("Downloading dlss files.\nThis may take a minute based on your connection speed.")
	download := exec.Command(
		"curl",
		"-o",
		customDLL,
		"-L",
		downloadLink)
	_, err := download.Output()

	if err != nil {
		_ = fmt.Errorf("something happened %s", err)
	}
}

func CheckDlssVersion(location string) string {

	loca := strings.ReplaceAll(location, "\\", "\\\\")
	trying := loca + "\\\\nvngx_dlss.dll"

	cmd := exec.Command("wmic")
	cmdLine := "datafile where name=" + "\"" + trying + "\"" + " get Version /value"
	cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: "/c " + os.ExpandEnv(cmdLine)}
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(err)
	}

	version := strings.Split(string(out), "=")[1]
	version = strings.TrimSpace(version)
	version = version[:len(version)-2]

	return version
}
