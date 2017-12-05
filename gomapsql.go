package gomapsql

import (
    "fmt"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
    "bytes"
    _ "os"
    "gopkg.in/doug-martin/goqu.v4"
    _ "gopkg.in/doug-martin/goqu.v4/adapters/sqlite3"
)

func Version() {
    fmt.Println("gomapsql 1.0")
}

func SqlMap(querysql string,data []map[string]interface{})(rtn []map[string]interface{},err error) {

    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    sqlBuilder := goqu.New("sqlite3",db)

    // 根据数据建临时表
    sqlStmt := CreateTableSql(data[0],"data")

    _, err = db.Exec(sqlStmt)
    if err != nil {
        log.Printf("%q: %s\n", err, sqlStmt)
        return
    }
    // 导入数据
    insert := sqlBuilder.From("data").Insert(data)
    if _,err := insert.Exec(); err!=nil{
        fmt.Println(err.Error())
    }

    // 执行sql
    rs := []map[string]interface{}{}
    rows, err := db.Query(querysql)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    columns, _ := rows.Columns()
    length := len(columns)

    for rows.Next(){
        row := make([]interface{},length)
        columnPointers := make([]interface{}, length)
        for i:=0;i<length;i++ {
            columnPointers[i] = &row[i]
        }

        rows.Scan(columnPointers...)
        tmpmap := make(map[string]interface{})

        for i:=0 ; i<length ;i++{
            columnName := columns[i]
            columnValue := row[i]
            tmpmap[columnName] = columnValue
        }
        rs = append(rs,tmpmap)
    }
    rtn = rs
    // 清理
    return

}

func LoadMap() {

}

func DbLoadMap() {

}


func GetSqliteType(i interface{}) string{
     switch i.(type){
        case int:
            return "integer"
        case float64:
            return "float"
        case string:
            return "text"
        case bool:
            return "integer"
        case []byte:
            return "blob"
        // case time.Time:
            // return "datetime"
        default:
            return "text"
    }
}
// 根据map数据返回建表语句
func CreateTableSql(data map[string]interface{},tablename string) (string){
    var sql bytes.Buffer

    sql.WriteString("create table ")
    sql.WriteString(tablename)
    sql.WriteString(" ( ")
    first := 0
    for columnName ,value := range data{
        if first>0 {
            sql.WriteString(" , ")
        }
        sql.WriteString(columnName)
        sql.WriteString(" ")
        sql.WriteString(GetSqliteType(value))
        first ++
    }
    sql.WriteString(" ); ")

    return sql.String()

}
