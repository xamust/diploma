package collect

import (
	"github.com/sirupsen/logrus"
)

type Collect struct {
	Logger           *logrus.Logger
	Config           *Config
	ParsingDataFiles map[string]string
}

func (c *Collect) Start() error {
	c.Logger.Print("Start collecting...")

	//search *.data files...
	result, err := c.searchDataFiles()
	if err != nil {
		return err
	}
	c.Logger.Print("Поиск *.data файлов выполнен успешно!")

	//c.Logger.Print(result)
	c.ParsingDataFiles = result
	//init sms...

	return nil
}

func (c *Collect) searchDataFiles() (map[string]string, error) {
	parse := ParsingFolder{c.Config, c.Logger, make(map[string]string)}
	_, err := parse.FindFiles()
	if err != nil {
		return nil, err
	}
	return parse.mapFile, nil
}
