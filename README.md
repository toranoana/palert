# プロセス監視ちゃん

指定されたプロセス名を探して(完全一致)なければslackにメンションします

cronで数分に一度回すのに使われることを想定しています

### 実行

- 引数1 探すプロセス名
- 引数2 なかった場合のslackに投げるメッセージ
- 引数3 slack webhook URL

```
./palert hoge "@channel hoge process not found" https://hooks.slack.com/services/XXX
```

### 必要ライブリ

```
go get github.com/mitchellh/go-ps
```

