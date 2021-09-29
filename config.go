package behavoir_tree

import (
	"encoding/json"
	"github.com/pkg/errors"
	"os"
)

type NodeConfig struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Type     NodeType `json:"type"`
	Title    string   `json:"title"`
	Desc     string   `json:"desc"`
	Children []string `json:"children"`
	Attr     map[string]interface{}
}

type ProjectConfig struct {
	Trees []string     `json:"trees"`
	Nodes []NodeConfig `json:"nodes"`
	Name  string       `json:"name"`
	Desc  string       `json:"desc"`
}

func LoadProject(file string) (ProjectConfig, error) {
	conf := ProjectConfig{}
	buf, err := os.ReadFile(file)
	if err != nil {
		return conf, errors.Wrap(err, "load project "+file)
	}
	err = json.Unmarshal(buf, &conf)
	if err != nil {
		return conf, errors.Wrap(err, "json decode " + file)
	}
	return conf, nil
}

func (p *ProjectConfig) Run(id string) {

}