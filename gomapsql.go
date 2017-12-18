package gomapsql

import (
    "fmt"
    "database/sql"
    _"github.com/mattn/go-sqlite3"
    "log"
    // "bytes"
    _ "os"
    "github.com/luyucia/gocatdb"

)

type Mapsql struct{
    db *sql.DB
    cdb gocatdb.Catdb
    err error
}

func (this *Mapsql) New(){
    // fmt.Println("init")
    this.db, this.err = sql.Open("sqlite3", ":memory:")
    if this.err != nil {
        log.Fatal(this.err)
    }
    // defer this.db.Close()

    this.cdb = gocatdb.Catdb{}
    this.cdb.BindDb(this.db,"sqlite3")

}


func Version() {
    fmt.Println("gomapsql 1.0")
}

// 异常情况
// sql语句异常
// map数据类型异常
// 数据插入异常

func (this *Mapsql)  Sql(querysql string)(rtn []map[string]interface{}) {
    rtn = this.cdb.Query(querysql)
    this.err = this.cdb.GetError()
    return
}

func (this *Mapsql) LoadMap(data []map[string]interface{},tablename string) {
    this.cdb.Table(tablename).Create(data[0])
    this.err = this.cdb.GetError()
    if this.err == nil {
        this.cdb.Table(tablename).Insert(data)
        this.err = this.cdb.GetError()
    }

}

func (this *Mapsql) GetError() (error) {
    return this.err
}

func (this *Mapsql) Clean() {
    this.db.Close()
}

func DbLoadMap() {

}
