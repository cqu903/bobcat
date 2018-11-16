package config

//全局配置对象，用于加载位于根目录下的配置文件
//*.default.yaml是默认配置文件，首先进行加载
//*.yaml文件是用户配置文件，会覆盖默认配置文件中相同的值
import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cqu903/bobcat/utils"
	yaml "gopkg.in/yaml.v2"
)

type Server struct {
	Port        int    `yaml:"port"`
	EnableHTTPS bool   `yaml:"enableHTTPS"`
	StaticDir   string `yaml:"staticDir"`
}
type App struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Author  string `yaml:"author"`
}
type AppConfig struct {
	Server
	App
}

var Conf = new(AppConfig)

//读取配置文件的值，构建全局配置对象
func Config() {
	//load default config
	Conf.Server.Port = 8080
	Conf.Server.EnableHTTPS = false
	Conf.Server.StaticDir = "/static/"

	Conf.App.Name = "bobcat framework"
	Conf.App.Version = "0.1"
	Conf.App.Author = "Roy Yuan"

	//load user customer config
	currentPath := utils.GetBaseDir()
	dir, err := os.Open(currentPath)
	if err != nil {
		log.Fatalf("can't open the current work path: %v", err)
	}
	defer dir.Close()
	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatalf("error:%v", err.Error())
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yaml") && !strings.HasSuffix(file.Name(), ".default.yaml") {
			err := praseYaml(filepath.Join(currentPath, file.Name()), Conf)
			if err != nil {
				log.Fatalf("load config file error,%v", err)
				break
			}
		}
	}
	if Conf == nil {
		log.Fatalln("Load bobcat config has wrong,can't find the right config file!please check!")
	}
}

func praseYaml(path string, v interface{}) error {
	var (
		data []byte
		err  error
	)
	data, err = ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, v)
	return err
}
