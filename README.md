## 开发 api 流程
1. 在 api/desc 目录定义 api 描述文件。
因为我们是单体服务，所以统一需要统一服务名称。

```
// 执行
goctl api go --api api/desc/platform.api --dir . --style go_zero
```

`--style go_zero` 表示文件风格，请保持统一，不然文件会重新生成。

2. 在生成的 Logic 文件中实现逻辑，如下图示例：

![alt text](image.png)

3. 如果 Logic 代码存在复用的，可以抽离到 biz 目录下，如下图示例：

![alt text](image-1.png)

4. 如果新增表，创建表之后一键生成 model 文件。

```
// 指定 internal/model 目录
goctl model pg datasource --url="postgres://username:root@host:5432/postgres?sslmode=disable" --table="tb_user_core" --dir  internal/model
```

生成的 model 如下示例：

![alt text](image-2.png)

生成完成后，需要统一在 internal/svc/service_context.go 文件中进行实例化，并管理相关依赖。

![alt text](image-3.png)

5. 编写代码逻辑。

6. 服务发布。