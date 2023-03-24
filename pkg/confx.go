package pkg

import (
	logs "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// VerifiableConfig config需要实现简单的自我校验
type VerifiableConfig interface {
	Verify() error
}

// absolutePath Returns absolute path
func absolutePath(configFile string) string {
	if configFile == "" {
		configFile = "app.yml"
	}

	// 如果是绝对路径，就不再处理; 否则认为是项目conf目录下的文件
	if !strings.HasPrefix(configFile, "/") {
		configFile = path.Join(RootPath(), "conf", configFile)
	}
	configFile, _ = filepath.Abs(configFile)
	return configFile
}

// LoadConfigFor 加载 config
// 例如:
//
//	appConfig := &AppConfig{}
//	err := LoadConfigFor(appConfig, "app.yml")
func LoadConfigFor(config VerifiableConfig, configFile string) error {
	configFile = absolutePath(configFile)

	oldAddressesFile := configFile
	data, err := os.ReadFile(oldAddressesFile)

	if err != nil {
		logs.Errorf("app config error: %v", err)
		return err
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		logs.Errorf("app config Unmarshal error: %v", err)
		return err
	}

	return config.Verify()
}

// RootPath 获取此项目可执行文件的根目录路径
func RootPath() string {
	var rootDir string

	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	rootDir = filepath.Dir(filepath.Dir(exePath))

	tmpDir := os.TempDir()
	if strings.Contains(exePath, tmpDir) {
		_, filename, _, ok := runtime.Caller(0)
		if ok {
			rootDir = filepath.Dir(filepath.Dir(filename))
		}
	}

	return rootDir
}
