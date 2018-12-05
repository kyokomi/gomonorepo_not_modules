gomonorepo
=============================

非vgo版

```
└── project1
    ├── client // UnityとかAndroidとかiOSとかなんかいろいろ
    ├── server
    │   ├── pkg
    │   └── src
    │       ├── appname
    │       │   ├── api1
    │       │   │   ├── hoge
    │       │   │   │   └── hoge.go
    │       │   │   └── main.go
    │       │   ├── api2
    │       │   │   ├── fuga
    │       │   │   │   └── fuga.go
    │       │   │   └── main.go
    │       │   ├── common
    │       │   │   └── piyo
    │       │   │       └── piyo.go
    │       │   ├── go.mod
    │       │   ├── go.sum
    │       │   └── vendor
    │       └── tools
    │           └── hogecli
    │               ├── go.mod
    │               ├── go.sum
    │               └── main.go
    └── web // frontendのJavaScriptとかなんかいろいろ
```

- server/src/appname内のapi1とapi2でvendorを共有する
- server/src/appname/commonはapi1とapi2で共通のライブラリとして利用している
- api1 -> api2 や api2 -> api1が呼べてしまうので、もしNGにする場合はそういうlintツール導入しないとだめです
- tools下はツール毎にvendorを設定している（もしapiと揃えたいならserver/src/appname下へ）
- server下がgolang前提になってるのでもしgolang以外のapiとかがあるなら `server/go` とかにするといいかも？
