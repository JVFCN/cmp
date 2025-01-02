package Play

import (
	"CLMusicPlayer/Methods"
	"CLMusicPlayer/MusicList"
	"fmt"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"github.com/dhowden/tag"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/flac"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	sPath "path"
	"strconv"
	"strings"
	"time"
)

var Play = &cobra.Command{
	Use:   "play <path>",
	Short: "play music from a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		listNumber, err := cmd.Flags().GetInt16("list")
		if err != nil {
			fmt.Println(err)
			return
		}
		filePath, err := cmd.Flags().GetString("path")
		if err != nil {
			fmt.Println(err)
			return
		}

		caseJudge(filePath, listNumber)
	},
}

func caseJudge(inputPath string, listNumber int16) {
	speaker.Init(44100, 44100/10)
	if listNumber == -1 {
		if inputPath == "" {
			fmt.Println("请输入有效路径")
			return
		} else {
			playMusic(inputPath)
		}
	} else if listNumber < 1 {
		fmt.Println("请输入有效序号")
		return
	} else {
		fJson, _ := os.ReadFile(Methods.TEMPDIR + "list.json")
		jsonData, _ := jsonvalue.Unmarshal(fJson)

		musicList, _ := jsonData.GetArray("MusicList")
		arr := Methods.JsonArrToArr(musicList)

		var counts = 0

		path := strings.TrimSpace(arr[listNumber-1])
		files, _ := ioutil.ReadDir(path)

		//fmt.Println(path)

		var musicFilesPath []string

		for _, file := range files {
			if MusicList.In(sPath.Ext(file.Name()), MusicList.Format) {
				musicFilesPath = append(musicFilesPath, path+"\\"+file.Name())
				counts++
			}
		}
		fmt.Println("共" + strconv.Itoa(counts) + "首")

		for i, file := range musicFilesPath {
			fmt.Println("正在播放第" + strconv.Itoa(i+1) + "首")
			playMusic(file)
			fmt.Println("\n")
		}
	}
}

func playMusic(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var streamer beep.StreamSeekCloser
	var format beep.Format

	ext := sPath.Ext(path)

	// 解码 FLAC 文件
	if ext == ".flac" {
		streamer, format, err = flac.Decode(f)
		if err != nil {
			log.Fatal(err)
		}
		defer streamer.Close()
	} else if ext == ".mp3" {
		streamer, format, err = mp3.Decode(f)
		if err != nil {
			log.Fatal(err)
		}
		defer streamer.Close()
	}

	acmd := exec.Command("ffmpeg", "-i", path, "./output.mp3")
	out, err := acmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	file, _ := os.Open(path)
	m, err := tag.ReadFrom(file)
	if err == nil {
		fmt.Println(m.Title() + "\n专辑:" + m.Album() + "\n歌手:" + m.Artist())
	}
	play(format, streamer, out)

	//duration := Methods.DurationToSeconds(string(out))
	//for currentSecond := 0; currentSecond <= duration; currentSecond++ {
	//	time.Sleep(1 * time.Second)
	//}
}

func play(format beep.Format, streamer beep.StreamSeekCloser, out []byte) {
	// 初始化音频设备
	//err := speaker.Init(format.SampleRate*1, format.SampleRate.N(time.Second/10))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//speaker.Lock()
	//speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	//speaker.Unlock()
	//// 播放音频
	speaker.Play(streamer)
	// 等待播放结束
	duration := Methods.DurationToSeconds(string(out))
	//duration := streamer.Len() / format.SampleRate.N(time.Second) // 计算时长

	formatDur := Methods.FormatTime(duration)
	for currentSecond := 0; currentSecond <= duration; currentSecond++ {
		// 输出格式：分:秒 [进度条]
		fmt.Printf("\r%s %s", Methods.FormatTime(currentSecond)+"/"+formatDur, progressBar(currentSecond, duration, 50))
		time.Sleep(1 * time.Second)
	}
}

func progressBar(current, total, barWidth int) string {
	progress := float64(current) / float64(total)
	bar := int(progress * float64(barWidth))
	if bar > 0 {
		return "[" + strings.Repeat("=", bar-1) + ">" + strings.Repeat(" ", barWidth-bar) + "]"
	}
	return "[" + ">" + strings.Repeat(" ", barWidth-1) + "]"
}
