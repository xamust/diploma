package collect

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

type Collect struct {
	Logger           *logrus.Logger
	Config           *Config
	ParsingDataFiles map[string]string
}

func (c *Collect) Destroy() error {
	c.Logger.Info("Start destroying...")
	if err := c.Start(); err != nil {
		return err
	}
	for s, s2 := range c.ParsingDataFiles {
		log.Print(s2)
		if err := os.Remove(s2); err != nil {
			return err
		}
		c.Logger.Infof("%s удален, находиился по пути %s", s, s2)

	}
	c.Logger.Info("Удаление *.data файлов выполнен успешно!")
	return nil
}

func (c *Collect) Start() error {
	c.Logger.Info("Start collecting...")

	//search *.data files...
	result, err := c.searchDataFiles()
	if err != nil {
		return err
	}
	c.Logger.Info("Поиск *.data файлов выполнен успешно!")

	//c.Logger.Print(result)
	c.ParsingDataFiles = result

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
