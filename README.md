# ichimon

開発用備忘録
docker compose up で air（コードを改変すると自動でコンパイルを行いホットリロードしてくれる。）での開発環境が起動。
postman で API 開発する。
DB は pgAdmin を使用して作業している。コンテナを落とすとデータが消えるので docker compose down した場合は再度データ投入を行う？

リクエスト先は localhost:8080/questions など。
ポート番号 80 は front で使用予定のため。

localhost:8080/csrf （GET）で csrf トークン取得、X-CSRF-TOKEN にてえられた値をセット。
↓
localhost:8080/login （POST）でログイン、ボディに json で email と password を持たせる。(セットクッキーで token がクッキーにセットされる)
↓
任意のエンドポイントにリクエストを飛ばす。questions 作成は動作確認。

API ドキュメント
http://localhost:8080/swagger/index.html

swagger init コマンド
swag init --parseDependency --parseInternal
https://github.com/swaggo/swag/issues/810
