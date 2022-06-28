# ojosama-slack-app

Slack App ですわ〜〜❗❗

## 概要

テキストを壱百満天原サロメお嬢様風に変換する[jiro4989](https://github.com/jiro4989)さんのライブラリ
[ojosama](https://github.com/jiro4989/ojosama) を使って作ったSlack Appです。

Slack Workspaceごとにサーバーを用意してデプロイが必要な作りになっています。

## 準備

```sh
cp .env.example .env
```

Slack Appを作成する場合は、 [app_amnifests.example.yml](./app_amnifests.example.yml) も参考にしてください🗒

Slack Appを作成して、Slack App Tokenを発行し、`SLACK_APP_TOKEN=`の値を置き換えて下さい。
そのときは、Scope `connections:write` を付与することをお忘れないようにお願いします🌱
同じく、Workspace installをして、発行したSlack Bot Tokenで`SLACK_BOT_TOKEN=`の値を置き換えて下さい。

## Heroku デプロイ

heroku CLI を準備しくてださい❗

こちらが最初の準備です。

```sh
heroku apps:create 任意のアプリの名前
heroku stack:set container
```

`.env` ファイルの `HEROKU_APP_NAME=` を任意のアプリの名前に変更してください。

準備が済めば以下のコマンドでデプロイできます。

```sh
make deploy-heroku
```

## 注意事項

### プログラムの使用について

壱百満天原サロメお嬢様、及びその所属の にじさんじ や、
その関係者、ファンコミュニティの方の迷惑にならないように使ってください。

本プログラムは、にじさんじ所属の壱百満天原サロメお嬢様のキャラクターを題材にした二次創作の一つです。
故に、本プログラムは以下二次創作ガイドラインに従います。

* [ANYCOLOR二次創作ガイドライン](https://event.nijisanji.app/guidelines/)

本プログラムを使う場合も上記ガイドラインを守ってお使いください。
