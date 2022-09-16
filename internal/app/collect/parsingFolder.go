package collect

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"server/internal/app/systemsproject"
	"strings"
)

// ParsingFolder пакет для поиска *.data файлов
type ParsingFolder struct {
	config  *Config
	logger  *logrus.Logger
	mapFile map[string]string
}

// for correct checking data...
func (p *ParsingFolder) initMap() {
	p.mapFile = map[string]string{systemsproject.DBilling: "", systemsproject.DEmail: "", systemsproject.DSMS: "", systemsproject.DVoice: ""}
}

// checking...
func (p *ParsingFolder) checkData() error {
	for k := range p.mapFile {
		if p.mapFile[k] == "" {
			return fmt.Errorf("файл %s, не найден в директории %s", k, p.config.DataFolder)
		}
	}
	return nil
}

func (p *ParsingFolder) FindFiles() (bool, error) {
	//init...
	p.initMap()
	//parse files...
	if err := p.readDir(p.config.DataFolder); err != nil {
		return false, nil
	}
	//check files...
	if err := p.checkData(); err != nil {
		return false, err
	}
	return true, nil
}

func (p *ParsingFolder) readDir(directory string) error {
	data, err := os.ReadDir(directory)
	if err != nil {
		return err
	}
	for _, v := range data {
		switch {
		case strings.Split(v.Name(), ".")[len(strings.Split(v.Name(), "."))-1] == "data":
			p.mapFile[v.Name()] = fmt.Sprintf("%s/%s", directory, v.Name())
		case v.IsDir():
			p.readDir(fmt.Sprintf("%s/%s", directory, v.Name()))
		}
	}
	return nil
}
