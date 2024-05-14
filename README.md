## 前言
开发此程序的初衷是，市面上类似的工具都是命令行的，对本人来说用起来不爽，而且比如：python ,nodejs运行条件较为苛刻，环境配置麻烦，综上开发此应用。


## 功能
- [x] 小程序包解密解压
- [x] 敏感信息扫描
- [x] 自定义正则
- [x] 文件浏览
- [x] 分类筛选

## 技术栈
golang+wails(TS+vue3+vite+ant-design)

## 依赖
使用前必须先安装webview2 https://developer.microsoft.com/en-us/microsoft-edge/webview2

## 用法
1.小程序包解压，选择【*.wxapkg】包的上级目录即可，如果路径中包括wxid，那么程序会自动识别并填入，如果路径中没有请自行填入wxid   
2.小程序批量解压，如下，请以wxid为目录名，把小程序包放到对应目录下即可，程序会自动识别wxid
![image](https://user-images.githubusercontent.com/53891640/215651971-35f954c7-34b4-4e19-88be-c5257126e3e5.png)

3.敏感信息扫描，程序自带了部分正则，也可以自定义正则进行扫描，扫描完成后点击路径，可以自动打开文件

## 图示
![image](https://github.com/o0x1024/IMiniCrack/assets/53891640/91e38ec5-1203-49b6-807e-2ee8702ae81a)
![image](https://github.com/o0x1024/IMiniCrack/assets/53891640/75f47c08-1b47-4c77-b577-607c557cab52)
![image](https://github.com/o0x1024/IMiniCrack/assets/53891640/a4f660e1-be6e-4796-a3ba-b34069de93aa)




## 说明
本程序仅用于学习交流，请勿用于非法用途
