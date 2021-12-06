package conf

import (
	"fmt"
	"my-singo/util"
	"testing"
)

func init() {
	if err := InitLocales("zh", "/Users/chenjia/Documents/code/go_code/my-singo/conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("中文翻译文件加载失败", err)
	}
	if err := InitLocales("en", "/Users/chenjia/Documents/code/go_code/my-singo/conf/locales/us-en.yaml"); err != nil {
		util.Log().Panic("英文翻译文件加载失败", err)
	}
}

func TestLoadLocales(t *testing.T) {
	fmt.Println(Message("en", "Tag.required"))
	fmt.Println(Message("zh", "Tag.required"))
}
