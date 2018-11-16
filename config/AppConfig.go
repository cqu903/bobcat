package config

//全局配置对象，用于加载位于根目录下的配置文件
//*.default.yaml是默认配置文件，首先进行加载
//*.yaml文件是用户配置文件，会覆盖默认配置文件中相同的值
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/GeertJohan/go.rice"
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

var Conf *AppConfig

//读取配置文件的值，构建全局配置对象
func init() {
	//load default config
	conf := new(AppConfig)
	defaultConfig := rice.MustFindBox(".")
	defaultConfigBytes, err := defaultConfig.Bytes("conf.default.yaml")
	fmt.Println(defaultConfigBytes)
	if err != nil {
		log.Fatalf("load default config worng,errors:%v", err)
	}
	yaml.Unmarshal(defaultConfigBytes, conf)

	//load user customer config
	currentPath := utils.GetBaseDir()
	fmt.Println("读取到当前配置目录：" + currentPath)
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
	fmt.Println(Conf)
	if Conf == nil {
		log.Fatalln("全局配置对象加载失败，请检查!")
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
