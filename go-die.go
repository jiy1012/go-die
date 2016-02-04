package main

import (
    "github.com/go-martini/martini"
    "./config"
    "runtime"
    "fmt"
)

func main() {
    init_config()
    m := martini.Classic()
    port := config.AppConfig.AppPort
    m.RunOnAddr(":"+port)
    m.Get("/id", func() string {
        return "Hello world!"
    })
    m.Run()
}

func init_config() {
    goos := runtime.GOOS //linux windows
    fmt.Println(goos)
    config_path := ""
    if "windows" == goos {
        config_path = "dev/"
        martini.Env = martini.Dev
    }else if "linux" == goos {
        config_path = "prod/"
        martini.Env = martini.Dev
    }else if "mac" == goos {
        config_path = "dev/"
        martini.Env = martini.Prod
    }
    file_path := "./config/"+config_path+"global.xml"
    fmt.Println(file_path)
    err := config.LoadConfig(file_path)
    if err != nil {
        panic(err)
    }
    fmt.Println("init config success");
}