package crack

import (
	"IMiniCrack/pkg/jsbeautifier/jsbeautifier"
	"IMiniCrack/pkg/model"
	"IMiniCrack/pkg/scan"
	"IMiniCrack/pkg/util"
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/wailsapp/wails"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/crypto/pbkdf2"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Crack struct {
	ctx context.Context
	rt  *wails.Runtime
	log *wails.CustomLogger
}

type WxapkgFile struct {
	NameLen int
	Name    string
	Offset  int
	Size    int
}

type SliceFile struct {
	Name string
	Data string
}

type PackList struct {
	File string
	WxId string
	data []byte
}

func (c *Crack) GetCtx(ctx context.Context) {
	c.ctx = ctx
	scan.Sc.Ctx = ctx
}

// WailsInit .
func (c *Crack) WailsInit(runtime *wails.Runtime) error {
	fmt.Println("Crack WailsInit")
	c.log = c.rt.Log.New("Crack")
	c.rt = runtime
	return nil
}

func (c *Crack) GetFileList(outpath string) (resp model.Response) {
	var pList []string
	err := filepath.Walk(outpath, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() {
			pList = append(pList, path)
		}

		return nil
	})
	if pList != nil {
		resp.FileList = pList
	}
	if err != nil {
		resp.Err = err.Error()
		return resp
	}
	return resp
}

func (c *Crack) Unpack(wxpkgPath, wxid, outPath string) (resp model.Response) {
	if wxpkgPath == "" {
		resp.Err = "参数为空"
		return resp
	}
	//f, err := os.OpenFile("d:\\profile", os.O_CREATE|os.O_RDWR, 0666)
	//if err != nil {
	//	resp.Err = "pprof error"
	//	return resp
	//}
	//pprof.StartCPUProfile(f)

	packlist, err := c.decWxApkg(wxpkgPath, wxid)
	if err != nil {
		resp.Err = "解密失败,decWxApkg：" + err.Error()
		runtime.EventsEmit(c.ctx, "log", "解密失败,decWxApkg："+err.Error())
		return resp
	}

	for _, v := range packlist {
		tpath := outPath
		if wxid == "" {
			tpath = outPath + "\\" + v.WxId
		}
		err = c.unPackFile(v.File, v.data, tpath)
		if err != nil {
			resp.Err = "解密失败,unPackFile：" + err.Error()
			runtime.EventsEmit(c.ctx, "log", "解密失败,unPackFile："+err.Error())
			return resp
		}
	}

	//pprof.StopCPUProfile()
	resp.Msg = "解密导出成功"
	return resp
}

func (c *Crack) decWxApkg(wxapkgPath string, wxid string) ([]PackList, error) {
	salt := "saltiest"
	iv := "the iv: 16 bytes"
	//files, _ := os.ReadDir(wxapkgPath)
	haveWxapkg := false
	packList := []PackList{}
	err := filepath.Walk(wxapkgPath, func(path string, info os.FileInfo, err error) error {
		_wxid := ""
		if strings.Contains(path, "wxapkg") {
			haveWxapkg = true
		}
		if !info.IsDir() && strings.Contains(path, "wxapkg") {
			if wxid == "" {
				reg := regexp.MustCompile(`(?m)\\(wx\w+)\\`)
				wxids := reg.FindAllStringSubmatch(path, -1)
				if len(wxids) <= 0 {
					return nil
				}
				_wxid = wxids[0][1]
			} else {
				_wxid = wxid
			}

			//fmt.Println("i:", _wxid, "path:", path)

			pack := PackList{}
			pack.File = path
			originData := make([]byte, 1024)
			dataByte, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			dk := pbkdf2.Key([]byte(_wxid), []byte(salt), 1000, 32, sha1.New)
			block, _ := aes.NewCipher(dk)
			blockMode := cipher.NewCBCDecrypter(block, []byte(iv))

			blockMode.CryptBlocks(originData, dataByte[6:1024+6])

			afData := make([]byte, len(dataByte)-1024-6)
			var xorKey = byte(0x66)
			if len(_wxid) >= 2 {
				xorKey = _wxid[len(_wxid)-2]
			}
			for i, b := range dataByte[1024+6:] {
				afData[i] = b ^ xorKey
			}

			originData = append(originData[:1023], afData...)
			pack.data = originData
			pack.WxId = _wxid
			packList = append(packList, pack)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	if !haveWxapkg {
		return nil, errors.New("无小程序包")
	}

	return packList, nil
}

func (c *Crack) unPackFile(wxapkgPath string, data []byte, outRoot string) error {

	wxPackName := util.GetFileName(wxapkgPath)
	//fmt.Println()
	r := bytes.NewReader(data)
	firstMark := make([]byte, 1)
	_, err := r.Read(firstMark)
	if err != nil {
		return err
	}
	infoTable := make([]byte, 4)
	_, err = r.Read(infoTable)
	if err != nil {
		return err
	}
	indexInfoLength := make([]byte, 4)
	_, err = r.Read(indexInfoLength)
	if err != nil {
		return err
	}
	bodyInfoLength := make([]byte, 4)
	_, err = r.Read(bodyInfoLength)
	if err != nil {
		return err
	}
	lastMark := make([]byte, 1)
	_, err = r.Read(lastMark)
	if err != nil {
		return err
	}
	if bytes.Compare(firstMark, []byte{0xBE}) != 0 || bytes.Compare(lastMark, []byte{0xED}) != 0 {

		fmt.Println("err:", wxapkgPath)
		fmt.Printf("err  firstMark:%x  lastMark:%x\n", firstMark, lastMark)
		return errors.New("文件无效或wxid错误, PATH:" + wxapkgPath)
	}
	//fmt.Printf("firstMark:%x  lastMark:%x", firstMark, lastMark)

	fileCount := make([]byte, 4)
	_, err = r.Read(fileCount)
	if err != nil {
		return err
	}

	//read index
	fileList := []WxapkgFile{}
	var i uint32 = 0
	for ; i < binary.BigEndian.Uint32(fileCount); i++ {
		line := WxapkgFile{}
		nameLen := make([]byte, 4)
		r.Read(nameLen)
		line.NameLen = int(binary.BigEndian.Uint32(nameLen))

		name := make([]byte, line.NameLen)
		r.Read(name)
		line.Name = string(name)

		offset := make([]byte, 4)
		r.Read(offset)
		line.Offset = int(binary.BigEndian.Uint32(offset))

		size := make([]byte, 4)
		r.Read(size)
		line.Size = int(binary.BigEndian.Uint32(size))

		fileList = append(fileList, line)
	}

	//save file
	nameList := []string{}
	for _, v := range fileList {
		outFileName := v.Name
		fmt.Println("outRoot:", outRoot, "wxPackName:", wxPackName, "outFileName:", outFileName)
		outFilePath := outRoot + "\\" + wxPackName + "\\" + outFileName
		//fmt.Println("outFilePath:", outFilePath)

		nameList = append(nameList, outFilePath)
		parentDir := util.GetParentDirectory(outFilePath)
		if !util.PathExists(parentDir) {
			err := os.MkdirAll(parentDir, 0666)
			if err != nil {
				return err
			}
		}

		out, err := os.OpenFile(outFilePath, os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			fmt.Println("------")
			return err
		}

		runtime.EventsEmit(c.ctx, "log", outFilePath)

		r.Seek(int64(v.Offset), 0)
		buf := make([]byte, v.Size)
		r.Read(buf)
		str := string(buf)

		options := jsbeautifier.DefaultOptions()
		fbuf, err := jsbeautifier.Beautify(&str, options)
		//fbuf, err := jsbeautifier.Beautify(&str, optargs.MapType{"indent_size": 2, "space_in_empty_paren": true})
		if err != nil {
			return err
		}

		out.WriteString(fbuf)
		out.Close()
	}

	appServiceJsPath := ""
	for _, v := range nameList {
		if strings.Contains(v, "app-service.js") {
			appServiceJsPath = v
			break
		}
	}
	//fix js
	fp_appServiceJs, err := os.OpenFile(appServiceJsPath, os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	serverdata, err := io.ReadAll(fp_appServiceJs)
	if err != nil {
		return err
	}
	parseData := strings.Split(string(serverdata), "define(\"")
	//wxmlData := parseData[0]

	//fmt.Println(wxmlData)

	sliceList := []SliceFile{}
	for _, slice := range parseData[1:] {
		line := SliceFile{}
		arr := strings.SplitN(slice, "\",", 2)
		line.Name = arr[0]
		line.Data = arr[1][:strings.LastIndexAny(arr[1], "});")+1]
		sliceList = append(sliceList, line)
	}

	for _, sfile := range sliceList {
		if !strings.Contains(sfile.Name, ".png") || !strings.Contains(sfile.Name, ".jpg") {

			outFilePath := outRoot + "\\" + wxPackName + "\\" + sfile.Name

			parentDir := util.GetParentDirectory(outFilePath)
			if !util.PathExists(parentDir) {
				err := os.MkdirAll(parentDir, 0666)
				if err != nil {
					return err
				}
			}

			out, err := os.OpenFile(outFilePath, os.O_CREATE|os.O_RDWR, 0666)
			if err != nil {
				return err
			}
			options := jsbeautifier.DefaultOptions()
			fbuf, err := jsbeautifier.Beautify(&sfile.Data, options)
			if err != nil {
				return err
			}

			out.WriteString(fbuf)
			out.Close()
		}
	}
	fmt.Println("end")
	return nil
}
