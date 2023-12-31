# 2fa
一个简单的两步验证器

## Usage

### 添加
```shell
2fa add -n NAME -s SECRET
```

其他参数请查看 `2fa add -h`

### 显示
```shell
2fa showall [-d]
```
显示所有账号及其两步验证码，`-d` 不显示 **hotp** 类型的账号（因为显示一次计数器就要 +1）

```shell
2fa show [-s ISSUER] NAME
```
显示指定 NAME 的账号，如果有 **-s** 选项，则还会筛选 ISSUER

配置文件位于 `$HOME/.2fa`，格式为 yaml
