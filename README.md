
# Install Beego
```
go get github.com/astaxie/beego

```

# 依赖
```
go get github.com/go-sql-driver/mysql
go get github.com/astaxie/beego/orm
go get github.com/astaxie/beego/toolbox
go get github.com/astaxie/beego/logs
```

# Run
```
bee run
```

# Build
```
GOOS=linux GOARCH=amd64 bee pack
```

# Swagger
API: http://127.0.0.0.1/swagger/


# 可变化字段
- 部门: ```Name```, ```Parentid```
- 用户: ```Mobile```, ```Workplace```, ```Avatar```, ```Position```,```Departmentid```

