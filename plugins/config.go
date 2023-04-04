package plugins

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	T      *widget.Tree // 全局主节点
	appMap = map[string][]string{
		"": {},
	} // 一级节点
)

// 将config.json的内容序列化到config结构体中
func GetConfig() map[string]interface{} {
	var config map[string]interface{}
	f, _ := os.Open("config.json")
	defer f.Close()
	decoder := json.NewDecoder(f)
	decoder.Decode(&config)
	return config
}

// 初始化右边的界面
func InitTrailingNode(firstnode string) (items []*container.TabItem) {
	// GetConfig()[firstnode].(map[string]interface{}) 获取每个一级节点对应的二级节点名称
	for secoundnode := range GetConfig()[firstnode].(map[string]interface{}) {
		b := NewGetParentButton(firstnode, secoundnode, nil) // 每个容器需要一个初始化按钮,作用为添加新的按钮
		content := ContentTertiary(b)                        // 将初始化按钮添加到容器中
		// GetConfig()[firstnode].(map[string]interface{})[secoundnode].(map[string]interface{}) 获取每个二级节点对应的三级节点名称
		for buttonname := range GetConfig()[firstnode].(map[string]interface{})[secoundnode].(map[string]interface{}) {
			content.Add(NewSuperButton(buttonname, content, b.NodeName, b.ContainerName)) // 初始化config中的按钮
		}
		b.OnTapped = func() {
			p := AddContentPopUp(content, b.NodeName, b.ContainerName)
			p.Show()
		}
		// 按一级节点的名称,初始化二级节点
		item := container.NewTabItem(secoundnode, content)
		items = append(items, item)
	}
	return items
}

// 返回AppMap由HostNode接收,并创建成widget.Tree
func InitFirstNode() map[string][]string {
	config := GetConfig()
	// 获取所有第一级子节点的 key
	for key := range config {
		appMap[""] = append(appMap[""], key) // 将一级子节点内容添加到AppMap的第一行
	}
	return appMap
}

func SaveConfig(data map[string]interface{}) {
	// 将修改后的数据写入json文件
	file, err := os.Create("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateCofigFile() {
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		config := make(map[string]interface{})   // create an empty map to hold the config data
		configJSON, err1 := json.Marshal(config) // convert the map to JSON
		if err1 != nil {
			log.Fatalln(err1)
			return
		}
		err1 = ioutil.WriteFile("config.json", configJSON, 0644) // write the JSON to a file named "config.json"
		if err1 != nil {
			log.Fatalln(err1)
			return
		}
	}
}
