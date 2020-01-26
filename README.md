# auth_api
- signin(認証)
- jwt tokenを発行
- jwt tokenはredisに保存

## 用途
他のAPIはrequest headerに付与されるjwt tokenを使ってredisから認可情報を取得する  