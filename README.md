## 前言
开发此程序的初衷是，市面上类似的工具都是命令行的，对本人来说用起来不爽，而且比如：python ,nodejs运行条件较为苛刻，环境配置麻烦，综上开发起程序。


## 功能
- [x] 小程序包解密解压
- [x] 敏感信息扫描
- [x] 自定义正则
- [x] 文件浏览

## 技术栈
golang+wails(TS+vue3+vite+ant-design)

## 用法
1.小程序包解压，选择【*.wxapkg】包的上级目录即可，如果路径中包括wxid，那么程序会自动识别并填入，如果路径中没有请自行填入wxid
2.小程序批量解压，如下，请以wxid为目录名，把小程序包放到对应目录下即可，程序会自动识别wxid
![image](https://user-images.githubusercontent.com/53891640/215651971-35f954c7-34b4-4e19-88be-c5257126e3e5.png)

3.敏感信息扫描，程序自动了部分正则，也可以自定义正则进行扫描，扫描完成后点击路径，可以自动打开文件

## 图示
![image](https://user-images.githubusercontent.com/53891640/215632865-4c186384-ba65-4fd1-b6a5-1eaeb022c2ae.png)
![image](https://user-images.githubusercontent.com/53891640/215632919-fed52d14-c744-48ab-8b10-1230cf6a6a11.png)
![image](https://user-images.githubusercontent.com/53891640/215632939-8a47cc71-c1cd-4bbf-9666-e2a37d0808a1.png)



## 说明
本程序仅用于学习交流，请勿用于非法用途
