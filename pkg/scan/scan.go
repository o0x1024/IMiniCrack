package scan

import (
	"IMiniCrack/pkg/util"
	"bufio"
	"context"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/wailsapp/wails"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gopkg.in/yaml.v2"
)

var Sc *Scan

func init() {
	Sc = NewScan()
	Sc.Init()
}

type Scan struct {
	Regx   []regx `yaml:"regx"`
	Result []string
	rt     *wails.Runtime
	Ctx    context.Context
}

type Sensitive struct {
	Desc     string
	MatchStr string
	LineNo   string
	Path     string
}

type regx struct {
	Id     string `yaml:"id"`
	Desc   string `yaml:"desc"`
	Record string `yaml:"record"` //正则
}

func (s *Scan) SaveResult(content string) string {
	user, err := user.Current()
	if err != nil {
		return err.Error()
	}
	path, err := runtime.SaveFileDialog(s.Ctx, runtime.SaveDialogOptions{
		DefaultDirectory: user.HomeDir + "\\Documents",
		Title:            "save result",
		DefaultFilename:  "result.txt",
	})
	if err != nil {
		return err.Error()
	}

	err = os.WriteFile(path, []byte(content), 0666)
	if err != nil {
		return err.Error()
	}

	return "保存成功"
}

// 敏感信息包括
// 手机号、身份证号、
func (s *Scan) FindSensitiveInfo(content string) (string, string, error) {

	//sst := 	[]types.SenSitiveType{
	//	{"cradno18",`[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]`,},
	//	{"cradno15",`[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{2}`},
	//	{"phone",`(13[0-9]|14[5-9]|15[0-3,5-9]|16[2,5,6,7]|17[0-8]|18[0-9]|19[0-3,5-9])\d{8}`},
	//	{"aliyun_oss_url","[\\w-.]\\.oss.aliyuncs.com"},
	//	{"access_key","[Aa](ccess|CCESS)_?[Kk](ey|EY)|[Aa](ccess|CCESS)_?[sS](ecret|ECRET)|[Aa](ccess|CCESS)_?(id|ID|Id)"},
	//	{"secret_key","[Ss](ecret|ECRET)_?[Kk](ey|EY)"},
	//	{"slack_token","(xox[p|b|o|a]-[0-9]{12}-[0-9]{12}-[0-9]{12}-[a-z0-9]{32})"},
	//	{"slack_webhook","(xox[p|b|o|a]-[0-9]{12}-[0-9]{12}-[0-9]{12}-[a-z0-9]{32})"},
	//	{"mailgun_api","key-[0-9a-zA-Z]{32}"},
	//	{"mailchamp_api","[0-9a-f]{32}-us[0-9]{1,2}"},
	//	{"picatic_api","sk_live_[0-9a-z]{32}"},
	//	{"google_oauth_id","[0-9(+-[0-9A-Za-z_]{32}.apps.qooqleusercontent.com"},
	//	{"amazon_aws_access_key_id","AKIA[0-9A-Z]{16}"},
	//	{"amazon_mws_auth_token","amzn\\.mws\\.[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"},
	//	{"amazonaws_url","s3\\.amazonaws.com[/]+|[a-zA-Z0-9_-]*\\.s3\\.amazonaws.com"},
	//	{"facebook_access_token","EAACEdEose0cBA[0-9A-Za-z]+"},
	//	{"twilio_api_key","SK[0-9a-fA-F]{32}"},
	//	{"twilio_account_sid","AC[a-zA-Z0-9_\\-]{32}"},
	//	{"twilio_app_sid","AP[a-zA-Z0-9_\\-]{32}"},
	//	{"paypal_braintree_access_token","access_token\\$production\\$[0-9a-z]{16}\\$[0-9a-f]{32}"},
	//	{"square_oauth_secret","sq0csp-[ 0-9A-Za-z\\-_]{43}"},
	//	{"square_access_token","sqOatp-[0-9A-Za-z\\-_]{22}"},
	//	{"stripe_standard_api","access_token\\$production\\$[0-9a-z]{16}\\$[0-9a-f]{32}"},
	//	{"stripe_restricted_api","rk_live_[0-9a-zA-Z]{24}"},
	//	{"github_access_token",`[a-zA-Z0-9_-]*:[a-zA-Z0-9_\\-]+@github\\.com*`},
	//	{"private_ssh_key","-----BEGIN PRIVATE KEY-----[a-zA-Z0-9\\S]{100,}-----END PRIVATE KEY——"},
	//	{"private_rsa_key","-----BEGIN RSA PRIVATE KEY-----[a-zA-Z0-9\\S]{100,}-----END RSA PRIVATE KEY-----"},
	//	{"jwt1","[= ]ey[A-Za-z0-9_-].[A-Za-z0-9._-]"},
	//	{"jwt2","[= ]ey[A-Za-z0-9_/+-].[A-Za-z0-9._/+-]"},
	//	{"Email","[\\w!#$%&\\'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&\\'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[\\w](?:[\\w-]*[\\w])?"},
	//	{"Assets",`\b(?:(?:25[0-5]|2[0-4][0-9]|[1]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\b`},
	//}

	color := []string{"red", "blue", "orange", "magenta", "chocolate", "pink", "copper", "thistle", "khaki", "seagreen", "lightslategray", "firebrick"}
	var matches []string
	for _, v := range s.Regx {

		tmpreg, err := regexp2.Compile(v.Record, 0)
		if err != nil {
			return "", "", err
		}
		m, err := tmpreg.FindStringMatch(content)
		if err != nil {
			return "", "", err
		}
		for m != nil {
			matches = append(matches, m.String())
			m, _ = tmpreg.FindNextMatch(m)
		}
		if matches != nil {
			allstr := ""
			for i, v := range matches {
				if i >= len(color) {
					allstr += " <span style=\"color: red;\">" + v + "</span>"
				} else {
					allstr += " <span style=\"color: " + color[i] + ";\">" + v + "</span>"
				}
			}
			return allstr, v.Desc, nil
		}

	}
	return "", "", nil
}

func (s *Scan) ScanSensitive(path string) ([]Sensitive, string) {

	if path == "c:\\" {
		return nil, "请检查待扫描目录是否正确"
	}

	sensitives := []Sensitive{}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fp, err := os.OpenFile(path, os.O_RDWR, 0666)
			if err == nil {
				scanner := bufio.NewScanner(fp)
				var lineNo int = 0
				for scanner.Scan() {
					lineNo++
					text := scanner.Text()
					matchStr, desc, err := s.FindSensitiveInfo(text)
					if err != nil {
						return err
					}

					if matchStr != "" {
						sensitive := Sensitive{}
						matchStr = strings.Replace(matchStr, "\n", "", -1)
						sensitive.Desc = desc
						sensitive.MatchStr = matchStr
						no := strconv.Itoa(lineNo)
						sensitive.LineNo = no
						sensitive.Path = path
						result := fmt.Sprintf("%s | %s | line: %d |  %s", desc, matchStr, lineNo, path)
						//fmt.Println(result)

						s.Result = append(s.Result, result)
						runtime.EventsEmit(s.Ctx, "scan", result)
					}
				}
			}
			fp.Close()
		}
		return nil
	})
	if err != nil {
		return nil, err.Error()
	}

	return sensitives, "扫描结束"
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func NewScan() *Scan {
	return &Scan{}
}

func (s *Scan) WailsInit(runtime *wails.Runtime) error {
	s.rt = runtime
	return nil
}

func (s *Scan) SaveRegex() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	yamlPath := user.HomeDir + "\\.iminicrack\\scan.yaml"
	data, err := yaml.Marshal(s.Regx)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(yamlPath, data, 0777)
	if err != nil {
		return "", err
	}
	return "保存成功", nil
}

func (s *Scan) DelRegex(Id string) {
	newRegx := []regx{}
	for _, v := range s.Regx {
		if v.Id != Id {
			newRegx = append(newRegx, v)
		}
	}
	//fmt.Println(newRegx)
	s.Regx = newRegx
}

func (s *Scan) AddRegex(regexString string, desc string) {
	newId := len(s.Regx) + 1
	n := strconv.Itoa(newId)
	regex := regx{
		Id:     n,
		Record: regexString,
		Desc:   desc,
	}
	fmt.Println(regex)
	s.Regx = append(s.Regx, regex)

	//fmt.Println(s.Regx)
}

func (s *Scan) UpdateRegex(id, desc, record string) {

	for i, v := range s.Regx {
		if v.Id == id {
			s.Regx[i].Desc = desc
			s.Regx[i].Record = record
			//fmt.Println(s.Regx)
			return
		}
	}

}

func (s Scan) GetRegx() interface{} {
	return s.Regx
}

func (s *Scan) Init() {
	fmt.Println(" scan WailsInit")
	user, err := user.Current()
	checkError(err)

	iminiPath := user.HomeDir + "\\.iminicrack"
	if !util.PathExists(iminiPath) {
		//fmt.Println(iminiPath)
		err = os.Mkdir(iminiPath, 0666)
		checkError(err)
	}
	yamlPath := iminiPath + "\\scan.yaml"
	if !util.PathExists(yamlPath) {
		regs := []regx{
			{"1", "access_key", "[Aa](ccess|CCESS)_?[Kk](ey|EY)|[Aa](ccess|CCESS)_?[sS](ecret|ECRET)|[Aa](ccess|CCESS)_?(id|ID|Id)"},
			{"2", "OSS", "([A|a]ccess[K|k]ey[I|i][d|D]|[A|a]ccess[K|k]ey[S|s]ecret)"},
			{"3", "phone", `[^\w]((?:(?:\+|00)86)?1(?:(?:3[\d])|(?:4[5-79])|(?:5[0-35-9])|(?:6[5-7])|(?:7[0-8])|(?:8[\d])|(?:9[189]))\d{8})[^\w]`},
			{"4", "邮箱", `(([a-z0-9][_|\.])*[a-z0-9]+@([a-z0-9][-|_|\.])*[a-z0-9]+\.((?!js|css|jpg|jpeg|png|ico)[a-z]{2,}))`},
			{"5", "Secret_key", "[Ss](ecret|ECRET)_?[Kk](ey|EY)"},
			{"6", "github_access_token", `[a-zA-Z0-9_-]*:[a-zA-Z0-9_\\-]+@github\\.com*`},
			{"7", "JWT", `(eyJ[A-Za-z0-9_-]{10,}\.[A-Za-z0-9._-]{10,}|eyJ[A-Za-z0-9_\/+-]{10,}\.[A-Za-z0-9._\/+-]{10,})`},
			{"8", "Swagger UI", `((swagger-ui.html)|(\"swagger\":)|(Swagger UI)|(swaggerUi))`},
			{"9", "身份证", `[^0-9]((\d{8}(0\d|10|11|12)([0-2]\d|30|31)\d{3}$)|(\d{6}(18|19|20)\d{2}(0[1-9]|10|11|12)([0-2]\d|30|31)\d{3}(\d|X|x)))[^0-9]`},
			{"10", "RCE参数", "((cmd=)|(exec=)|(command=)|(execute=)|(ping=)|(query=)|(jump=)|(code=)|(reg=)|(do=)|(func=)|(arg=)|(option=)|(load=)|(process=)|(step=)|(read=)|(function=)|(feature=)|(exe=)|(module=)|(payload=)|(run=)|(daemon=)|(upload=)|(dir=)|(download=)|(log=)|(ip=)|(cli=))"},
			//{"Amazon AWS URL", `(((([a-zA-Z0-9._-]+\.s3|s3)(\.|\-)+[a-zA-Z0-9._-]+|[a-zA-Z0-9._-]+\.s3|s3)\.amazonaws\.com)|(s3:\/\/[a-zA-Z0-9-\.\_]+)|(s3.console.aws.amazon.com\/s3\/buckets\/[a-zA-Z0-9-\.\_]+)|(amzn\.mws\.[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})|(ec2-[0-9-]+.cd-[a-z0-9-]+.compute.amazonaws.com)|(us[_-]?east[_-]?1[_-]?elb[_-]?amazonaws[_-]?com))`},
			{"11", "Amazon AWS AccessKey ID", `[^0-9]((aws(.{0,20})?(?-i)[''\"][0-9a-zA-Z\/+]{40}[''\"])|((A3T[A-Z0-9]|AKIA|AGPA|AIDA|AROA|AIPA|ANPA|ANVA|ASIA)[a-zA-Z0-9]{16}))[^0-9]`},
			{"12", "Amazon AWS Region", `((us(-gov)?|ap|ca|cn|eu|sa)-(central|(north|south)?(east|west)?)-\d)`},
			{"13", "Password Field", `((|'|")([p](ass|wd|asswd|assword))(|'|")(:|=)( |)('|")(.*?)('|")(|,))`},
			{"14", "Authorization Header", `((basic [a-z0-9=:_\+\/-]{5,100})|(bearer [a-z0-9_.=:_\+\/-]{5,100}))`},
			//{"LinkFind ", `(?:"|')(((?:[a-zA-Z]{1,10}://|//)[^"'/]{1,}\.[a-zA-Z]{2,}[^"']{0,})|((?:/|\.\./|\./)[^"'><,;|*()(%%$^/\\\[\]][^"'><,;|()]{1,})|([a-zA-Z0-9_\-/]{1,}/[a-zA-Z0-9_\-/]{1,}\.(?:[a-zA-Z]{1,4}|action)(?:[\?|#][^"|']{0,}|))|([a-zA-Z0-9_\-/]{1,}/[a-zA-Z0-9_\-/]{3,}(?:[\?|#][^"|']{0,}|))|([a-zA-Z0-9_\-]{1,}\.(?:php|asp|aspx|jsp|json|action|html|js|txt|xml)(?:[\?|#][^"|']{0,}|)))(?:"|')`},
			{"15", "URL", `(?:\b[a-z\d.-]+://[^<>\s]+|\b(?:(?:(?:[^\s!@#$%^&*()_=+[\]{}\|;:'",.<>/?]+)\.)+(?:ac|ad|aero|ae|af|ag|ai|al|am|an|ao|aq|arpa|ar|asia|as|at|au|aw|ax|az|ba|bb|bd|be|bf|bg|bh|biz|bi|bj|bm|bn|bo|br|bs|bt|bv|bw|by|bz|cat|ca|cc|cd|cf|cg|ch|ci|ck|cl|cm|cn|coop|com|co|cr|cu|cv|cx|cy|cz|de|dj|dk|dm|do|dz|ec|edu|ee|eg|er|es|et|eu|fi|fj|fk|fm|fo|fr|ga|gb|gd|ge|gf|gg|gh|gi|gl|gm|gn|gov|gp|gq|gr|gs|gt|gu|gw|gy|hk|hm|hn|hr|ht|hu|id|ie|il|im|info|int|in|io|iq|ir|is|it|je|jm|jobs|jo|jp|ke|kg|kh|ki|km|kn|kp|kr|kw|ky|kz|la|lb|lc|li|lk|lr|ls|lt|lu|lv|ly|ma|mc|md|me|mg|mh|mil|mk|ml|mm|mn|mobi|mo|mp|mq|mr|ms|mt|museum|mu|mv|mw|mx|my|mz|name|na|nc|net|ne|nf|ng|ni|nl|no|np|nr|nu|nz|om|org|pa|pe|pf|pg|ph|pk|pl|pm|pn|pro|pr|ps|pt|pw|py|qa|re|ro|rs|ru|rw|sa|sb|sc|sd|se|sg|sh|si|sj|sk|sl|sm|sn|so|sr|st|su|sv|sy|sz|tc|td|tel|tf|tg|th|tj|tk|tl|tm|tn|to|tp|travel|tr|tt|tv|tw|tz|ua|ug|uk|um|us|uy|uz|va|vc|ve|vg|vi|vn|vu|wf|ws|xn--0zwm56d|xn--11b5bs3a9aj6g|xn--80akhbyknj4f|xn--9t4b11yi5a|xn--deba0ad|xn--g6w251d|xn--hgbk6aj7f53bba|xn--hlcj6aya9esc7a|xn--jxalpdlp|xn--kgbechtv|xn--zckzah|ye|yt|yu|za|zm|zw)|(?:(?:[0-9]|[1-9]\d|1\d{2}|2[0-4]\d|25[0-5])\.){3}(?:[0-9]|[1-9]\d|1\d{2}|2[0-4]\d|25[0-5]))(?:[;/][^#?<>\s]*)?(?:\?[^#<>\s]*)?(?:#[^<>\s]*)?(?!\w))`},
		}
		data, err := yaml.Marshal(regs)
		checkError(err)

		err = os.WriteFile(yamlPath, data, 0666)
		checkError(err)

	} else {
		content, err := os.ReadFile(yamlPath)
		checkError(err)
		err = yaml.Unmarshal(content, &s.Regx)
		checkError(err)
	}
}

//	{"cradno18",`[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]`,},
//	{"cradno15",`[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{2}`},
//	{"phone",`(13[0-9]|14[5-9]|15[0-3,5-9]|16[2,5,6,7]|17[0-8]|18[0-9]|19[0-3,5-9])\d{8}`},
//	{"aliyun_oss_url","[\\w-.]\\.oss.aliyuncs.com"},
//	{"access_key","[Aa](ccess|CCESS)_?[Kk](ey|EY)|[Aa](ccess|CCESS)_?[sS](ecret|ECRET)|[Aa](ccess|CCESS)_?(id|ID|Id)"},
//	{"secret_key","[Ss](ecret|ECRET)_?[Kk](ey|EY)"},
//	{"slack_token","(xox[p|b|o|a]-[0-9]{12}-[0-9]{12}-[0-9]{12}-[a-z0-9]{32})"},
//	{"slack_webhook","(xox[p|b|o|a]-[0-9]{12}-[0-9]{12}-[0-9]{12}-[a-z0-9]{32})"},
//	{"mailgun_api","key-[0-9a-zA-Z]{32}"},
//	{"mailchamp_api","[0-9a-f]{32}-us[0-9]{1,2}"},
//	{"picatic_api","sk_live_[0-9a-z]{32}"},
//	{"google_oauth_id","[0-9(+-[0-9A-Za-z_]{32}.apps.qooqleusercontent.com"},
//	{"amazon_aws_access_key_id","AKIA[0-9A-Z]{16}"},
//	{"amazon_mws_auth_token","amzn\\.mws\\.[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"},
//	{"amazonaws_url","s3\\.amazonaws.com[/]+|[a-zA-Z0-9_-]*\\.s3\\.amazonaws.com"},
//	{"facebook_access_token","EAACEdEose0cBA[0-9A-Za-z]+"},
//	{"twilio_api_key","SK[0-9a-fA-F]{32}"},
//	{"twilio_account_sid","AC[a-zA-Z0-9_\\-]{32}"},
//	{"twilio_app_sid","AP[a-zA-Z0-9_\\-]{32}"},
//	{"paypal_braintree_access_token","access_token\\$production\\$[0-9a-z]{16}\\$[0-9a-f]{32}"},
//	{"square_oauth_secret","sq0csp-[ 0-9A-Za-z\\-_]{43}"},
//	{"square_access_token","sqOatp-[0-9A-Za-z\\-_]{22}"},
//	{"stripe_standard_api","access_token\\$production\\$[0-9a-z]{16}\\$[0-9a-f]{32}"},
//	{"stripe_restricted_api","rk_live_[0-9a-zA-Z]{24}"},
//	{"github_access_token",`[a-zA-Z0-9_-]*:[a-zA-Z0-9_\\-]+@github\\.com*`},
//	{"private_ssh_key","-----BEGIN PRIVATE KEY-----[a-zA-Z0-9\\S]{100,}-----END PRIVATE KEY——"},
//	{"private_rsa_key","-----BEGIN RSA PRIVATE KEY-----[a-zA-Z0-9\\S]{100,}-----END RSA PRIVATE KEY-----"},
//	{"jwt1","[= ]ey[A-Za-z0-9_-].[A-Za-z0-9._-]"},
//	{"jwt2","[= ]ey[A-Za-z0-9_/+-].[A-Za-z0-9._/+-]"},
//	{"Email","[\\w!#$%&\\'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&\\'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[\\w](?:[\\w-]*[\\w])?"},
//	{"Assets",`\b(?:(?:25[0-5]|2[0-4][0-9]|[1]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\b`},
