package main

import (
	"CLMusicPlayer/Methods"
	"CLMusicPlayer/cmd"
	"fmt"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"os"
)

func main() {
	Methods.EnsureDir(Methods.TEMPDIR)
	if !Methods.FileExists(Methods.TEMPDIR + "counts.json") {
		os.Create(Methods.TEMPDIR + "counts.json")
	}

	if !Methods.FileExists(Methods.TEMPDIR + "list.json") {
		_, err := os.Create(Methods.TEMPDIR + "list.json")
		cjson := jsonvalue.NewObject()
		cjson.SetArray().At("MusicList")
		marshal, err := cjson.MarshalString()

		openFile, err := os.OpenFile(Methods.TEMPDIR+"list.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
		if err != nil {
			fmt.Println(err)
			return
		}
		openFile.WriteString(marshal)
		openFile.Close()
	}
	cmd.Execute()
}
