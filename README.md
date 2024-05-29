# Atcoder web app


## 要件
- サインアップ、ログイン、ログアウト機能
- ライバルユーザー登録、削除機能
- ライバルユーザーたちと自分のatcoder_id, streak, rating表示機能
- ライバルユーザーたちと自分の直近の提出表示機能

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
### /signup
- ユーザー作成
- email, password, atcoder_id入力
- responseはstatus codeのみでok
### /login
- ログイン
- レスポンス成功時にatcoder_id返したい
### /logout
- ログアウト、ログイン画面にもどる
### POST /user/rival/{target_id}
- ライバルユーザー追加
- target_id(登録したいatcoder_idを入力)
### delete /user/rival/{target_id}
- 既存のライバルユーザー削除
### GET /user/rival
- ライバルユーザーのrating, streakのリスト取得
### GET /user/submission/{atcoder_id}
- ライバルユーザーたちの提出リスト取得(1日)
### POST /user/profile/{ID}
- 自分のatcoder_id変更
### GET /user/profile
- 自分の情報取得
- streak
- 提出リスト(1週間)


## フロントcomponent
### login画面
- email, password入力
- signupと切り替えられるボタン
### signup画面
- email, password, atcoder_id入力
- loginと切り替えられるボタン
### ホーム画面
- ログイン後最初に表示
- 上半分にライバルユーザーtableコンポーネント
- 下半分にライバルユーザー提出リストコンポーネント
- 右上とかにプロフィールボタンとサインアウトボタン
### tableコンポーネント
- atcoder_id, streak, ratingをテーブル表示
- 各ユーザーの横に削除ボタン
- テーブルの下にユーザー追加ボタンとテキストボックス

### 提出リストコンポーネント
- ライバルユーザーたちと自分の今日のac情報表示(ACのみ)
- ユーザーid, 問題、時刻、difficulty

### プロフィールコンポーネント
- streak, rating, 1週間のac情報表示(ACのみ)
- atcoder_id変更ボタン

## API依存メモ
### ログイン、サインイン、サインアウトまわり
- user_controller -> user_usecase -> user_repository -> db
### プロフィールまわり
- 自分の1週間の提出、レート表示(get profile)
- atcoder id 変更(post proile)
- user_controller -> user_usecase -> user_repository,infra -> atcoder problems
- user_controllerは自分のIDとる
- user_usecaseは自分のIDからatcoder id取得してinfraのgetstreak(id)とgetsubmission(id, 7days)たたいていい感じにレスポンス変形
### rival
- ライバル追加、削除
- rival_controller -> rival_usecase -> rival_repository -> db
### streak
- ライバルと自分のstreak表示
- rival_controller -> rival_usecase -> infra
- rival_controllerは自分のIDからrivalのatcoder idすべてに対してgetstreak(id)とgetrating(id)たたいていい感じにレスポンス変形
### submission
- ライバルと自分の1週間の提出表示
- rival_controller -> rival_usecase -> infra
- infraはatcoder problemsたたくだけ
- usecaseはatcoder problems jsonから自分のsubmission response型に変換
- controllerはクライアントで使用するjsonに変換

## todo
- infra部分実装
  - getstreak : string -> int
  - getrating : string -> int
  - getsubmission : (string, time) -> []submission
- 上3つがusecaseからたたくもの
- getstreak, getsubmissionはfetchSubmissionたたいて全件取得する
- getratingは最後のsubmissionだけ取得すればよい
- 実装後に使っていない関数削除

## 動作確認
- GET /csrf : ok
- POST /signup : ok
  - リクエストbodyに、email, password, atcoder_id
  - atcoder_idのvalidationはできていないぽい
- POST /login : ok
  - リクエストbodyに、email, password
- POST /logout : ok
- POST /user/rival : ok
- DELETE /user/rival/:id : ok
- POST /profile/:id : updateで更新されたかどうかcheckできていない(同じatcoder idわたしてもレスポンス成功になっている)
- GET /table : ok
- GET /submission : ok