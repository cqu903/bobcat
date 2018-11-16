package config

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "AppConfig.go",
		FileModTime: time.Unix(1542350470, 0),
		Content:     string("package config\n\n//全局配置对象，用于加载位于根目录下的配置文件\n//*.default.yaml是默认配置文件，首先进行加载\n//*.yaml文件是用户配置文件，会覆盖默认配置文件中相同的值\nimport (\n\t\"fmt\"\n\t\"io/ioutil\"\n\t\"log\"\n\t\"os\"\n\t\"path/filepath\"\n\t\"strings\"\n\n\t\"github.com/GeertJohan/go.rice\"\n\t\"github.com/cqu903/bobcat/utils\"\n\tyaml \"gopkg.in/yaml.v2\"\n)\n\ntype Server struct {\n\tPort        int    `yaml:\"port\"`\n\tEnableHTTPS bool   `yaml:\"enableHTTPS\"`\n\tStaticDir   string `yaml:\"staticDir\"`\n}\ntype App struct {\n\tName    string `yaml:\"name\"`\n\tVersion string `yaml:\"version\"`\n\tAuthor  string `yaml:\"author\"`\n}\ntype AppConfig struct {\n\tServer\n\tApp\n}\n\nvar Conf *AppConfig\n\n//读取配置文件的值，构建全局配置对象\nfunc init() {\n\t//load default config\n\tconf := new(AppConfig)\n\tdefaultConfig := rice.MustFindBox(\".\")\n\tdefaultConfigBytes, err := defaultConfig.Bytes(\"conf.default.yaml\")\n\tfmt.Println(defaultConfigBytes)\n\tif err != nil {\n\t\tlog.Fatalf(\"load default config worng,errors:%v\", err)\n\t}\n\tyaml.Unmarshal(defaultConfigBytes, conf)\n\n\t//load user customer config\n\tcurrentPath := utils.GetBaseDir()\n\tfmt.Println(\"读取到当前配置目录：\" + currentPath)\n\tdir, err := os.Open(currentPath)\n\tif err != nil {\n\t\tlog.Fatalf(\"can't open the current work path: %v\", err)\n\t}\n\tdefer dir.Close()\n\tfiles, err := dir.Readdir(-1)\n\tif err != nil {\n\t\tlog.Fatalf(\"error:%v\", err.Error())\n\t}\n\tfor _, file := range files {\n\t\tif strings.HasSuffix(file.Name(), \".yaml\") && !strings.HasSuffix(file.Name(), \".default.yaml\") {\n\t\t\terr := praseYaml(filepath.Join(currentPath, file.Name()), Conf)\n\t\t\tif err != nil {\n\t\t\t\tlog.Fatalf(\"load config file error,%v\", err)\n\t\t\t\tbreak\n\t\t\t}\n\t\t}\n\t}\n\tfmt.Println(Conf)\n\tif Conf == nil {\n\t\tlog.Fatalln(\"全局配置对象加载失败，请检查!\")\n\t}\n}\n\nfunc praseYaml(path string, v interface{}) error {\n\tvar (\n\t\tdata []byte\n\t\terr  error\n\t)\n\tdata, err = ioutil.ReadFile(path)\n\tif err != nil {\n\t\treturn err\n\t}\n\terr = yaml.Unmarshal(data, v)\n\treturn err\n}\n"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "conf.default.yaml",
		FileModTime: time.Unix(1542350523, 0),
		Content:     string("server:\n  port: 80\n  enableHTTPS: false\n  staticDir: /static/\napp:\n  name: BobCat App\n  version: 0.1\n  author: roy\n"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "rice-box.go",
		FileModTime: time.Unix(1542350536, 0),
		Content:     string(""),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1542350536, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "AppConfig.go"
			file3, // "conf.default.yaml"
			file4, // "rice-box.go"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`.`, &embedded.EmbeddedBox{
		Name: `.`,
		Time: time.Unix(1542350536, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"AppConfig.go":      file2,
			"conf.default.yaml": file3,
			"rice-box.go":       file4,
		},
	})
}
