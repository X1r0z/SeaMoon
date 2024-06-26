# CHANGELOG
## SeaMoon 2.0.0

2.0.0 版本终于算是完成了一个小的里程碑，打出来第一个版本，在稳定性和使用上可能还是有一定的问题，预计在几个版本内实现fix。

###❤️ What's New
*  🎨 refactor: 重构了整体的架构，优化客户端使用
*  🎨 refaotr: 优化服务端日志结构代码
* ✨ feat: 支持了 vmess / ss / vless 等加密协议
* ✨ feat: 支持了 clash / shadowrocket 客户端一键导入
* ✨ feat: 支持了客户端登录认证
* ✨ feat: 支持了客户端服务流量统计、测速等常用功能。
* 🔧 fix(ci)：重新拆分了 ci，支持 2.0 特性外增加了可读性。
* 📝 docs：手册 2.0 版本更新
* .... 更多功能更新欢迎下载试用体验

### What's Changed

* fcde201 Merge pull request #61 from DVKunion/2.0-dev
* ff358b7 feat: add issue templates (#80)
* 32144e7 fix: some bugs and docs update (#83)
* 2bcec2e fix: update ci action (#84)

**Full Changelog**: https://github.com/DVKunion/SeaMoon/compare/1.2.0-beta.2...2.0.0


## SeaMoon 1.2.0-beta.2

### ❤️ What's New
* 📝 docs: 订正手册 (#49)
* 📝 docs: 手册增加了 [gost](https://github.com/go-gost/gost) 作为本地客户端的使用方式 (#57)
* 🔧 fix(client): 修复 ISSUE 提到的连接缓慢问题 (#48)(#23)(#44)

### What's Changed
* fix: use release binary instead of self compiling by @DVKunion in https://github.com/DVKunion/SeaMoon/pull/24
* fix: client error && websocket compress error by @DVKunion in https://github.com/DVKunion/SeaMoon/pull/48
* docs: update manual && update yarn && fix a little bug by @DVKunion in https://github.com/DVKunion/SeaMoon/pull/49
* fix: roll back vdoing by @DVKunion in https://github.com/DVKunion/SeaMoon/pull/50

**Full Changelog**: https://github.com/DVKunion/SeaMoon/compare/1.2.0-beta.1...1.2.0-beta.2

## SeaMoon 1.2.0-beta.1

>  1.2.0 是一个临时版本，由于想要加入一些新的功能与适配，整体代码改动量较大。因此先拆分了部分功能代码。
>  整体的稳定性上会有所欠缺，更多的是下一个版本的新功能尝试

### ❤️ What's New
* 📝 docs: 修正了首页的一些图床丢失 (#18)
* ✨ feat(server): 新增隧道协议：grpc (#19)
* ⚡️ zap(server):  将协议解码转移在服务端处理，轻量化客户端 (#17)
* 🎨 refartor(server): 重构服务端模式，参考 gost 代理项目优化网络传输 (#15)
* 🎨 refartor(client): 减轻客户端传递功能，做更多兼容性适配，配合server重构 (#20)
* 🔧 fix(ci): 修正一些重构导致的ci错误(#16)(#21)(#22)

**Full Changelog**: https://github.com/DVKunion/SeaMoon/compare/1.1.3...1.2.0-beta.1

* 75db770 feat: support grpc tunnel (#19)
* 122435f fix: Dockerfile (#22)
* ae49ee0 fix: page ci dir changes (#16)
* be04fa1 fix: tag ci error (#21)
* f40e8f0 refactor: client (#20)
* 22bc49c refactor: server code && upgrade go mod (#15)
* ef2bf71 zap: change socks handle from client to server. (#17)


## SeaMoon 1.1.3

### ❤️ What's New

* 📝 docs: 增加手册页面sitemap站点地图(#7)(#8)
* ✨ feat(server): 修改了阿里云默认的部署资源类型(vcpu 0.1/mem 128)，来降低普通用户使用的价格消费 (#10)
* ✨ feat(server): 增加了sealos部署方案，用更加便宜的价格使用seamoon (#11)
* ✨ feat(server): 增加了docker server， 现在可以通过docker来启动服务端  (#12)
* 🔧 fix(config): 用更友好的方式来使用config，不再单一的通过域名特征来判断服务端地址类型。(#13)

**Full Changelog**: https://github.com/DVKunion/SeaMoon/compare/1.1.2...1.1.3

* 41c5ce8 feat(docker): add docker server (#12)
* 1414293 feat: low cpu && mem cost (#10)
* 99c98fd fix(client): use more friendly config (#13)

## SeaMoon 1.1.2

### ❤️ What's New

* 🔧 fix(websocket): 修正了protocol error detect 时仍挂起gorouting导致卡死的问题 (#6)
* ✨ feat(dockerfile): 增加了docker client， 现在可以通过docker来启动客户端  (#6)

**Full Changelog**: https://github.com/DVKunion/SeaMoon/compare/1.1.1...1.1.2

## SeaMoon 1.1.1

### ❤️ What's New

* 🔧 fix(websocket): 修正了 websocket 在超出 32768 slice导致的 panic。 (#4)
* 🔧 fix(websocket): 修整了 websocket 在 close 时写入 message 导致的 panic。 (#4)
* 🔧 fix(websocket): 忽略了大量 websocket 链接导致的 1006 abnormal close 报错。 (#4)
* 🔧 fix(s.yaml): 修整了 serverless-devs 工具编排文件，目前可以通过 serverless-devs 工具`s deploy`一件部署至阿里云。 (#4)
* 🔧 fix(ci): 修整了 go-releaser ci 配置 (#3)
* 🔧 fix(docs): 更新了 README.md 较为过时的使用手册。

### 🌈 Small Talk

> Hi, 各位，SeaMoon成功挤入2023Kcon兵器谱，使得整个项目获得了一批关注；在此感谢大家对SeaMoon项目的浓厚兴趣与支持，谢谢各位🙏    
> 由于工作原因，以及个人的一些想法枯竭，项目于去年创建，直到现在目前也仅支持了阿里云一个demo QAQ，因此整体给人并不是一个较为完善的使用效果。1.1.1 版本后，我会尽量保证一些活跃性质的更新，以及一些比较有意思的想法demo迭代。  
> 也欢迎对serverless感兴趣的安全小伙伴留言来交个朋友～

**Full Changelog**: https://github.com/DVKunion/SeaMoon/compare/1.1.0...1.1.1

* bc209a9 doc: update README.md
* a2e7360 fix: go-releaser ci config (#3)
* 8f51e63 fix: readme.md
* fe658ff fix: some websocket error optimization (#4)
* c316527 hotfix: some docs and code format

## 1.1.0 (2022-09-27)

### Bug Fixes

* optimize connection ([70dfc5a](https://github.com/DVKunion/SeaMoon/commit/70dfc5ad4d25fd5b529097183c873d87ec37f126))
* optimize connection ([2b416c0](https://github.com/DVKunion/SeaMoon/commit/2b416c0b106ad0a6a21aa3da838cf311061e9ef8))

## 1.0.0 (2022-09-09)

### Features

* **ci:** add build
  client ([215400c](https://github.com/DVKunion/SeaMoon/commit/215400cb7a3ae6c3f5f12df6828c8735156b810b))
* **pkg/socks5:** socks5 proxy beta
  version ([20d586c](https://github.com/DVKunion/SeaMoon/commit/20d586ce1ac36f143c1e340aa3bf9132e35af230))
* **pkg/http:** http proxy beta
  version ([3b41846](https://github.com/DVKunion/SeaMoon/commit/3b41846f75fe6d9510a9d040d76f97b35ce8c494))



