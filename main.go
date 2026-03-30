package main

import (
	"fmt"

	"github.com/rjchen8/torrent-client/predownload"
)

func main() {
	content, err := predownload.ReadFile("sample/LibreOffice_26.2.2_MacOS_aarch64.dmg.torrent")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(content)
}
