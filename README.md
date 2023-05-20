# SGE TechBook vol3

2023年N月の技術書典にて執筆した「gRPC × Goのゲーム開発で少しでも実装コストを削減してみる」のサンプルサービスのソースコードです。

ディレクトリの構成については下記ツリーを参考にしてください。

```
tech-book-vol3-sample
├── README.md
├── Makefile    // 各種実行コマンドをまとめます
├── api/
│   └── protos/ // Protocol Buffersの定義をまとめます
├── build/      // Dockerfileなどサービスのビルドに関わるファイルをまとめます
├── cmd/        // 各種エントリーポイントをまとめます
│   ├── api/
│   └── cli/
├── db/         // マイグレーションに関わるファイルを纏めます
├── internal/   // サンプルサービスの実装をまとめます
├── pkg/
│   └── pb/     // サンプルサービスのProtocol Buffersから作成されるGoのコードをまとめます
└── tools/      // 自動生成やマイグレーションを行うためのツールをまとめます
```

## 起動確認

Dockerがインストールされている必要があります。

```shell
# DBの起動。ポート3306で起動します。
$ make up SERVICE=mysql
$ make db-migrate

# APIの起動。ポート50051で起動します。
$ make up SERVICE=api
```

各種自動生成やcliの呼び出し方。

```shell
# Protocol Buffersの定義からGoのコードを生成する
$ make buf-generate

# DBからエンティティを生成する
$ make go-dbgen

# マスターデータ配信用のハンドラを生成する
$ make go-masterhandlergen

# gRPCのメソッドを呼び出すためのCLIのサブコマンドを生成する
$ make cli-subcmdgen

# CLIからローカルに起動しているgRPCサーバのメソッドを呼び出す
$ make cli-build
$ ./bin/cli grpc <subcmd> <MethodName> -h localhost -p 50051 --insecure [-d '{...}']
```

