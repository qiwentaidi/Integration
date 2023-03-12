package common

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/spf13/viper"
)

// 修改配置文件，通过directory\method两个字段是否为空判断是执行删除或者是添加操作
func Editor(v *viper.Viper, containername, name, directory, method string) {
	// 读取配置
	v.SetConfigFile("config.json")
	if err := v.ReadInConfig(); err != nil {
		dialog.ShowInformation("error", fmt.Sprintf("Failed to read config.json:%v", err), fyne.CurrentApp().Driver().AllWindows()[0])
		return
	}
	// 查找指定键值对并删除
	if obj := v.GetStringMap(containername); obj != nil {
		delete(obj, name)
	} else {
		dialog.ShowInformation("error", "删除键值对失败", fyne.CurrentApp().Driver().AllWindows()[0])
	}
	if directory != "" || method != "" {
		// 查找指定键值对并删除
		if obj := v.GetStringMap(containername); obj != nil {
			delete(obj, name)
		}
		// 重新写入新的值
		v.Set(fmt.Sprintf("%s.%s.method", containername, name), method)
		v.Set(fmt.Sprintf("%s.%s.path", containername, name), directory)
	}

	// 写入配置文件
	if err := v.WriteConfigAs("./config.json"); err != nil {
		dialog.ShowInformation("", fmt.Sprintf("Failed to write config file: %v\n", err), fyne.CurrentApp().Driver().AllWindows()[0])
		return
	}
}

func Profile() {
	// 设置配置文件名称和路径
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		// 如果配置文件不存在，则创建一个空的配置文件
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := viper.SafeWriteConfig(); err != nil {
				fmt.Println("Failed to create config file:", err)
			}
		} else {
			fmt.Println("Failed to read config file:", err)
		}
	}
}
