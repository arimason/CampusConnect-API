package configs

import (
	"errors"
	"io/ioutil"

	"github.com/go-chi/jwtauth"
	"gopkg.in/yaml.v3"
)

// definindo configurações para o banco de dados e webServer
type config struct {
	DBDriver      string           `yaml:"DB_DRIVER"`
	DBHost        string           `yaml:"DB_HOST"`
	DBPort        int              `yaml:"DB_PORT"`
	DBUser        string           `yaml:"DB_USER"`
	DBPassword    string           `yaml:"DB_PASSWORD"`
	DBName        string           `yaml:"DB_NAME"`
	WebServerPort string           `yaml:"WEB_SERVER_PORT"`
	JWTSecret     string           `yaml:"JWT_SECRET"`
	JWTExperesIn  int              `yaml:"JWT_EXPERESIN"`
	TokenAuth     *jwtauth.JWTAuth `yaml:"-"`
}

// // funçao é lida antes do método main
// func init() {

// }

// utilizado apenas para gerar o arquivo YAML de acordo com a Configs
// func GenerateYAML(cfg *Config, fileName string) error {
// 	yamlData, err := yaml.Marshal(cfg)
// 	if err != nil {
// 		return errors.New("erro ao gerar YAML: " + err.Error())
// 	}
// 	err = ioutil.WriteFile(fileName, yamlData, 0664)
// 	if err != nil {
// 		return errors.New("erro ao escrever arquivo YAML: " + err.Error())
// 	}
// 	return nil
// }

func LoadConfigs(path string) (*config, error) {
	var cfg *config
	yamlData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New("erro ao ler o arquivo YAML: " + err.Error())
	}
	err = yaml.Unmarshal(yamlData, &cfg)
	if err != nil {
		return nil, errors.New("erro ao realizar unmarshal para a configuração: " + err.Error())
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, nil
}
