# Atcoder web app

## 動かし方
- `mv .env.sample .env`
- `docker-compose -f docker-compose.yml up --build`
- htpp://localhost:3000

## DB設計
### usersテーブル
- id : UUID primary key
- email 
- password
- atcoder_id

### raival_usersテーブル
- id : UUID primary key
- id_from
- id_to

## APIエンドポイント
### POST /signup
- ユーザー作成
- email, password, atcoder_id入力
### POST /login
- ログイン
### POST /logout
- ログアウト、ログイン画面にもどる
### POST /user/rival/{target_id}
- ライバルユーザー追加
- target_id(登録したいatcoder_idを入力)
### DELETE /user/rival/{target_id}
- 既存のライバルユーザー削除
### GET /user/rival
- ライバルユーザーリスト取得
### GET /user/table
- ライバルユーザーのrating, streakのリスト取得
### GET /user/submission/
### GET /user/profile
- 自分のatcoder_id取得
### POST /user/profile/{ID}
- 自分のatcoder_id変更


## フロントcomponent
### login, signup画面
- email, password入力でログイン
- email, password, atcoder_id入力でsignup->login
### user画面
- ログイン後最初に表示
- タイトル下にlogoutリンク
- 上半分にライバルユーザーtableコンポーネント
  - その下に編集リンク
- 下半分にライバルユーザー提出リストコンポーネント
### edit画面
- 自分のatcoder id変更
- ライバルユーザー追加
- ライバルユーザー削除


## todo
- デプロイ
- GET Table, GET Submissionでユーザーが存在しないときなどerrを返さずに0,0を返したい
- atcoder idのvalidationのバグ修正

## メモ
- postmanでテスト<->chromeでテストの切り替えの際に以下が必要
  - user_controllerのcookie.secure変更
  - middlewareのsamemode
- docker-compose.ymlで動かす<->npm start, go run main.goで別々に動かす際にdb.goのhost変える必要がある