package config

import (
    "os"
    "encoding/xml"
    "io/ioutil"
//    "time"
//    "strconv"
    "fmt"
)
//xml中的mysql配置
type mysqlhost struct{
    XMLName     xml.Name    `xml:"mysql"`
    Name        string      `xml:"name,attr"`
    Host        string      `xml:"host"`
    Port        string      `xml:"port"`
    User        string      `xml:"user"`
    Password    string      `xml:"password"`
    CharSet     string      `xml:"charSet"`
}
//xml中的table配置
type dbTable struct {
    XMLName     xml.Name    `xml:"table"`
    Link        string      `xml:"host"`
    DbName      string      `xml:"db_name"`
    DbNum       int         `xml:"db_num"`
    TableName   string      `xml:"name,attr"`
    TableNum    int         `xml:"table_num"`
}
//xml中的配置解析
var confXml struct {
    XMLName         xml.Name        `xml:"go-die"`
    Version         string          `xml:"version,attr"`
    Port            string          `xml:"port"`
    Mysqllink       []mysqlhost     `xml:"mysql"`
    DbTableConfig   []dbTable       `xml:"table"`
}

type app_mysql map[string]interface{}
type app_table map[string]interface{}

var AppConfig struct{
    AppPort         string
    ServerVersion   string
    AppMysqlHost    app_mysql
    AppTableConfig  app_table
}

func LoadConfig(file string) error {
//    st := time.Now().UnixNano();
    cr ,err := os.Open(file)
    if err != nil {
        return err
    }
    defer cr.Close()
    data, err := ioutil.ReadAll(cr)
    if err != nil {
        return err
    }
    c_xml := confXml
    err = xml.Unmarshal(data, &c_xml)
    if err != nil {
        return err
    }
    AppConfig.AppPort = c_xml.Port
    AppConfig.ServerVersion = c_xml.Version
    AppConfig.AppMysqlHost = make(app_mysql)
    for _,value := range c_xml.Mysqllink  {
        AppConfig.AppMysqlHost[value.Name] = value
    }
    AppConfig.AppTableConfig = make(app_table)
    for _,value := range c_xml.DbTableConfig  {
        AppConfig.AppTableConfig[value.TableName] = value
    }

    fmt.Println(AppConfig)
//    et := time.Now().UnixNano()
//    miuns := et - st
//    fmt.Println("s:"+strconv.FormatInt(st,10)+" et:"+strconv.FormatInt(et,10))
//    fmt.Println("use : "+strconv.FormatInt(miuns,10))
    return nil
}