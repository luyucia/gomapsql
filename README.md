# gomapsql
A go package for sql operate on map,you can use it to write sql to operate a map list


### 依赖:
需要安装sqlite3
如果可以翻墙直接执行下面命令
```
go get github.com/mattn/go-sqlite3
```

如果不可以,需要手动安装它的依赖golang.org/x/net
```
mkdir -p $GOPATH/src/golang.org/x/
cd $GOPATH/src/golang.org/x/
git clone https://github.com/golang/net.git net
go install net
然后再执行
go get github.com/mattn/go-sqlite3
```


### 安装gomapsql
如果前面的依赖装好了,直接执行
```
go get github.com/luyucia/gomapsql
```

### 使用
## 装载
```
mapsql.LoadMap(datamap,"tablename")
```

## 在map上执行sql操作
```
mapsql.Sql(datamap,sql)
```

## 样例
```
    import "github.com/luyucia/gomapsql"
    import "fmt"

    func main() {

    data := []map[string]interface{}{
        map[string]interface{}{"id":1,"uv":101.1,"pv":10,"name":"luyu","city":"长春"},
        map[string]interface{}{"id":1,"uv":111.0,"pv":10,"name":"luyu","city":"长春"},
        map[string]interface{}{"id":2,"uv":100.0,"pv":10,"name":"luyu","city":"长春"},
        map[string]interface{}{"id":2,"uv":100.5,"pv":10,"name":"luyu","city":"长春"},
        map[string]interface{}{"id":3,"uv":100.3,"pv":10,"name":"luyu","city":"长春"},
    }

    // 初始化
    mapsql := &gomapsql.Mapsql{}
    mapsql.New()

    // 装载map到t1
    mapsql.LoadMap(data,"t1")

    // 装载map到t2
    mapsql.LoadMap(data,"t2")

    // 在map上操作sql
    sql := `select * from t1 a join t2 b on a.id=b.id order by uv asc`
    fmt.Println(mapsql.Sql(sql))

    // 清理
    mapsql.Clean()


    }

```
