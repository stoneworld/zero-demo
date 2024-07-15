## 开发 api 流程
1. 在 api/desc 目录定义 api 描述文件。

```
// 执行
goctl api go --api api/desc/platform.api --dir . --style go_zero
```

`--style go_zero` 表示文件风格，请保持统一，不然文件会重新生成。

2. 在生成的 Logic 文件中实现逻辑，如下图示例：

![alt text](image.png)

3. 服务发布。