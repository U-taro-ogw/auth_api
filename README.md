# auth_api
- signin(認証)
- jwt tokenを発行
- jwt tokenはredisに保存(key: jwt_token value: 認可情報)

# redis
- 課題1  
  redisの扱い  
  メソッドを生やすものか  
  connectionを引数とした関数にするか
  
- 課題2  
  UserHandler内でredisのconnectionを作成するのはおかしい。  
  この辺をスマートにやりたい
  
- 課題3  
  テストを書く