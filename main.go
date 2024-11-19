package main

// date: 2024/10/30
// email: brach@lssin.com

import (
	"github.com/spf13/cobra"
	"github.com/wanggaolin/go_lib/w"
	"guitar/vender"
)

var (
	mod_debug                 string = "debug"
	flag_video_scan_directory string // 扫描目录
	GoVersion                 string
	ReleaseTime               string
	Auchar                    string = ""
	AppVersion                string = "guitar-tools/0.1"
	run_args                  *vender.Body_args
)

var rootCmd = &cobra.Command{
	Use:   "guitar-tools",
	Short: "guitar tools",
	Run: func(cmd *cobra.Command, args []string) {
		if run_args.Version {
			w.Show_version(GoVersion, Auchar, AppVersion)
		} else {
			vender.Guitar.Main(run_args)
		}
	},
}

func main() {
	run_args = &vender.Body_args{}
	rootCmd.PersistentFlags().BoolVarP(&run_args.Version, "version", "v", false, "查询版本号")
	rootCmd.PersistentFlags().BoolVarP(&run_args.Chord_Level, "level", "l", false, "和弦级数问答测试")
	rootCmd.PersistentFlags().BoolVarP(&run_args.Chord_Scale, "scale", "s", false, "和弦够成音问答测试")
	rootCmd.PersistentFlags().BoolVarP(&run_args.Chord_Empty, "empty", "e", false, "吉他空弦音测试")
	//rootCmd.PersistentFlags().BoolVarP(&run_args.Query, "query", "q", false, "生成在线问答")
	rootCmd.PersistentFlags().BoolVarP(&run_args.ShowAll, "all", "a", true, "题库随机问答测试")
	rootCmd.Execute()
}
