# auth_api
- signin(認証)
- jwt tokenを発行
- jwt tokenはredisに保存

## 用途
他のAPIはrequest headerに付与されるjwt tokenを使ってredisから認可情報を取得する  

### 今後
順序は逆だけど、テストを書く。  
auth_apiのテストを書いたら、他のAPIをTDDで実装していく。
