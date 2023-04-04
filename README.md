# Integration

其他项目推荐：https://www.aiviy.com/item/rolan（好用且轻量，内存占用也少，但是收费，可以找破解版。）

```
1、按钮功能支持自定义，配置自己所需要用到的工具即可，支持
	cmd -- 以文件所在的cmd窗口启动  	// 命令行操作工具
	exe -- exe启动 			  	 // exe窗口工具
	jar -- jar的GUI工具启动		   // Java GUI窗口工具
	bat -- bat工具启动 				// burp之类的,也可以就是个脚本

2、一、二级节点支持添加删除，一级节点支持重命名

3、按钮右键支持自定义配置路径，打开文件目录或删除按钮（同时会将配置文件中的结果删除）
```

## 使用说明：

1. 首先打开EXE程序，什么都没有

   ![image-20230404173223844](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230404173223844.png)

2. 点击左上角设置添加分页，创建一个新的主节点。

   ![image-20230404173307182](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230404173307182.png)

   ![image-20230404173406588](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230404173406588.png)

3. 主节点位置标签，分为左右两个部分，**左边对应的是左键（左键点击实现翻页，要靠左边，点歪不会响应），右边对应的是右键（实现重命名和删除主节点的功能）**。

   ![image-20230404173505402](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230404173505402.png)

4. 左键点击后，**点击下方输入需要添加二级节点的名称（不输入无法添加二级节点），然后在点击左上角+号，新增二级节点**，出现添加按钮控件即创建成功。

   ![image-20230404173835511](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230404173835511.png)

   ![image-20230404174017681](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230404174017681.png)

   ![image-20230404174050594](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230404174050594.png)

5. 点击添加按钮，完成添加，在添加的按钮上右键存在三个功能（编辑、删除、打开文件目录），编辑设置正确的启动方式和路径，点击保存完成。

   ![image-20230404174306085](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230404174306085.png)

   ![image-20230404174411577](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230404174411577.png)

   ![image-20230404182255673](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230404182255673.png)

6. 左键Fscan,以工具所在路径的CMD窗口打开。

   ![image-20230404182556922](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230404182556922.png)

7. 其他支持远程桌面一键启动，或者注册表一键启动（没软用，但是注册表需要以管理员权限运行才能启动），备忘录功能

   ![image-20230404182730119](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230404182730119.png)

## 配置文件：

![image-20230404182839262](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230404182839262.png)

```
scanner为一级节点（主节点），内网一把梭为二级节点，Fscan为按钮，method为启动方法，path为启动路径
```

## 注意事项：

```
1、请检查你的工具权限是否为最高，若不是可能导致某些程序无法启动（例如蚁剑）

2、类似于冰蝎、哥斯拉等工具会在当前工具目录生成文件，若丢失则会导致无法启动，也可能会有生成的文件冲突的问题。

3、该工具启动jar程序默认启动命令为 java -jar 若你本地环境java的名称，不是这个会导致启动失败,使用 java -version 查看配置是否正确,若不想更改环境变量可以通过修改 widget/rightclickbutton.go文件中的 java -jar 字段修改成 本地存在的 java jdk8变量。

4、在添加burp bat程序指向时无法启动，
bat修改-javaagent:%burp-loader-keygen-2_1_06.jar为下面代码
	  -javaagent:%~dp0burp-loader-keygen-2_1_06.jar        ，%~dp0为指定bat当前路径

5、注册表无法启动,需要以管理员权限运行程序
```

## 缺点：

1. 暂不支持热键呼出或隐藏
2. 不支持贴边隐藏（写不来，这个GUI少了很多鼠标拖动事件）
3. 不支持文件直接拖入生成路径（做不来）

## 项目地址：

https://github.com/qiwentaidi/Integration/

## 打包：

```
windows 无命令行窗口打包 ： go build -ldflags -H=windowsgui main.go 

使用upx压缩工具(github上有)，可以使工具大小到10Mb
```

## 联系方式：

![image-20230312101259727](https://qwtd-image.oss-cn-hangzhou.aliyuncs.com/img/image-20230312101259727.png)
