package main

import (
	"fmt"
	"github.com/zhangyiming748/FastMediaInfo"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"quick-video-retranscoder/util"
	"strings"
)

func init() {
	if util.GetRoot() == "" {
		util.SetRoot()
	}
	util.SetLog()
	util.ExitAfterRun()
}
func main() {
	err := filepath.Walk(util.GetRoot(), func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			absPath, err := filepath.Abs(p)
			if err != nil {
				return err
			}
			fmt.Printf("准备处理的文件夹%v\n", info.Name())
			files := util.GetAllFiles(absPath)
			for _, file := range files {
				if strings.Contains(file, "quick") {
					continue
				}
				mi := FastMediaInfo.GetStandMediaInfo(file)
				path, fname := filepath.Split(file)
				baseName := strings.Replace(fname, filepath.Ext(fname), "", 1)
				extension := strings.Replace(filepath.Ext(fname), ".", "", 1)
				afterName := strings.Join([]string{baseName, "quick"}, "_")
				afterName = strings.Join([]string{afterName, "mp4"}, ".")
				afterName = strings.Join([]string{path, afterName}, "")
				/*
					ffmpeg -i RKI116.avi -ss 20:38.480 -to 23:43.401 -c:v libvpx-vp9 -crf 31 -c:a libopus -vbr on -ac 1 -map_chapters -1 RKI116p1.mp4
				*/
				cmd := exec.Command("ffmpeg", "-i", file, "-c:v", util.GenerateFFmpegParamsForCurrentSystem(), "-c:a", "copy", afterName)
				log.Printf("mi:%+v\npath:%+v\nbase:%+v\next:%+v\naftername:%+v\ncmd:%+v\n", mi, path, baseName, extension, afterName, cmd.String())
				err = util.ExecCommand(cmd, fmt.Sprintf("正在处理快照的视频:%v", baseName))
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
