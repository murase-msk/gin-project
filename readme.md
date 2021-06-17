# Gin Project

## 開発環境

- Windows10
- WSL2(Ubuntu20.04)
- docker
  - docker-compose
- Visual Studio Code
- Git
  - GitHub
  - SourceTree
- Go
  - Gin


## 環境構築

```
cd ~/projectFilePath
//docker起動
docker-compose up -d
```

## コマンドメモ

```
//docker サービス起動(WSL起動時に毎回実行する)
sudo service docker start
//シェルに入るとき
docker-compose exec go /bin/bash
//dockerを停止するとき
docker-compose stop

```

## トラブルシューティング
- VScodeでファイル編集
  - //プロジェクトディレクトリで下記コマンド実行
    code .
    //VScode画面の左下歯車マーク->コマンドパレット->Remote-Containers:Attach to Runnning Container
    //プロジェクトフォルダを開く
- ホストPCからWSL2のlocalhostにアクセスできない
  - touch c:\Users\<ユーザ名>\.wslconfig
    下記の内容を記述する
    localhostForwarding=True

- debugが動かないとき
  - 最初にvscodeでブレークポイントを作ってデバックを実行する必要がある？理由不明
  PCを再起動or高速スタートアップをしないor WSL再起動(powershellでwsl --shutdown)
  - https://qiita.com/snaka/items/a8eee4cfc8f7d733e6ab

## 参考URL
- https://tech-lab.sios.jp/archives/18811
- https://qiita.com/yiheng-lin/items/510b56454c30c7e00635


