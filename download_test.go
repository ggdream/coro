package main

import "testing"


func TestDownload(t *testing.T) {
	url := "https://video.pearvideo.com/mp4/adshort/20210210/cont-1720061-15604794_adpkg-ad_hd.mp4"
	if err := downer(url); err != nil {
		panic(err)
	}
}
