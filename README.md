# Integration

目前渗透工具资料繁多，有时候经常找个工具找半天，为了应对这种情况，才出现了这个启动器

其他项目推荐：https://www.aiviy.com/item/rolan（但是收费，可以找破解版）

```
1、按钮功能完全支持自定义，配置自己所需要用到的工具即可，支持
	a.以文件所在的cmd窗口启动  	// 命令行操作工具
	b.exe启动 			  	 // exe窗口工具
	c.jar工具启动		   	    // Java GUI窗口工具
	d.bat工具启动 				// burp之类的

2、具备托盘功能，点击X号关闭程序并不会完全退出，需要在系统任务栏里进行关闭。

3、可以自定义或删除分页

4、按钮右键支持自定义配置路径，打开文件目录或删除按钮（同时会将配置文件中的结果删除）
```

![image-20230305155012101](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230305155012101.png)

## 使用说明：

1. 点击设置-添加分页

   ![image-20230312093240969](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230312093240969.png)

2. 点击添加按钮，编写按钮名称点击save保存，退出这个小窗口需要点cancle

   ![image-20230312093545032](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230312093545032.png)

3. 按钮右键-编辑配置启动方式与工具路径，可以手动复制路径或者点击文件夹按钮进行工具路径选择（启动方法或者路径二者之中有一项不为空，点击保存才会写入配置）

   ![image-20230312093639335](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230312093639335.png)

   ![image-20230312093730404](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230312093730404.png)

4. 按钮右键-删除，可以将不需要的按钮进行删除，删除分页会将当前分页进行删除。类似于列表的按钮是用于展开或收缩手风琴按钮。

5. 样式-分页样式    修改分页是水平排列或者垂直排列（锤子排列比较占空间）

   ![image-20230312094738807](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230312094738807.png)

   ![image-20230312094753292](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230312094753292.png)

6. 其他支持远程桌面一键启动，或者注册表一键启动（没软用，但是注册表需要以管理员权限运行才能启动），备忘录功能

   ![image-20230312094823400](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230312094823400.png)

## 配置文件：

1. 修改config.json

   Integration可以根据config.json目录文件的内容自动生成按钮，通过method配置启动方式，path配置工具路径即可完成启动任务，如果在启动的时候配置，需要重启应用，编写格式为：

   ```
   {
   	"info":{ // 分页名称
   		"dirsearch":{ // 工具名称
   			"method":"folder", // 启动方法
   			"path":"路径" 
   		}
   	}
   }
   ```

2. 按钮右键配置

## 注意事项：

```
1、请检查你的工具权限是否为最高，若不是可能导致某些程序无法启动。

2、类似于冰蝎、哥斯拉等工具会在当前工具目录生成文件，若丢失则会导致无法启动，也可能会有生成的文件冲突的问题。

3、该工具启动jar程序默认启动命令为 java -jar 若你本地环境java的名称，不是这个会导致启动失败,使用 java -version 查看配置是否正确,若不想更改环境变量可以通过修改 widget/rightclickbutton.go文件中的 java -jar 字段修改成 本地存在的 java jdk8变量。

4、在添加burp bat程序指向时无法启动，
bat修改-javaagent:%burp-loader-keygen-2_1_06.jar为下面代码
	  -javaagent:%~dp0burp-loader-keygen-2_1_06.jar        ，%~dp0为指定bat当前路径

5、注册表无法启动,需要以管理员权限运行程序
```

## 缺点：

1. 暂不支持热键呼出或隐藏（本来做了，但是有bug）
2. 不支持贴边隐藏（写不来，这个GUI少了很多鼠标拖动事件）
3. 不支持文件直接拖入生成路径（做不来）
4. 配置写入之后名称都会转为小写（不知道为什么，可能跟viper库有关）可以通过修改配置文件改成大写开头（只支持分页大写）
5. 暂不支持名称重命名
6. 由于fyne框架本身不支持中文，所以需要打包tff字体文件，如果追求轻量化又不需要中文，可以删除字体文件进行并删除main.go的init函数重新打包

## 项目地址：

https://github.com/qiwentaidi/Integration/

## 打包：

```
windows 无命令行窗口打包 ： go build -ldflags -H=windowsgui main.go 

使用upx压缩工具(github上有)，可以使工具大小到11Mb
```

## 联系方式：

![image-20230312101259727](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230312101259727.png)
