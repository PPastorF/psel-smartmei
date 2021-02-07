package misc

import (
    "os"
    "fmt"
    "strings"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    types "github.com/ppastorf/psel-smartmei/internal/types"
)

func InProduction() bool {
    if strings.ToUpper(os.Getenv("DEPLOY_ENV")) == "PROD" {
        return true
    } else {
        return false
    }
} 

func ReadConfig(filename string) (*types.Config, error) {
    buf, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    c := &types.Config{}
    err = yaml.Unmarshal(buf, c)
    if err != nil {
        return nil, fmt.Errorf("Erro ao ler o arquivo %q: %v", filename, err)
    }

    return c, nil
}