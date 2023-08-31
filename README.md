# dynamic_service
## Задача:

Требуется реализовать сервис, хранящий пользователя и сегменты, в которых он состоит (создание, изменение, удаление сегментов, а также добавление и удаление пользователей в сегмент)

## Реализация:
#### На основе файла .env_example создать файл .env
### Запуск приложения
```bash
	go mod download
	swag init
	go run main.go
```

### Запуск контейнера
```bash
	docker-compose up --build -d
```

#### Добавление пользователя:
Endpoint:
<summary><code>POST</code> <code>http://127.0.0.1:3000/create_user</code></summary>

**Parameters**

|          Name | Required |  Type   | Description                                                                                                                                                           |
| -------------:|:--------:|:-------:| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|     `name` | required | string  | Имя   

#### Пример:

Input:
```json
{
	"name": "Adelina"
}
```

Output:
```json
{
	"id": 1,
	"name": "Adelina",
	"Segments": null
}
```

#### Добавление пользователя в сегмент:
Endpoint:
<summary><code>POST</code> <code>http://127.0.0.1:3000/add_user_to_segment</code></summary>

**Parameters**

|          Name | Required |  Type   | Description                                                                                                                                                           |
| -------------:|:--------:|:-------:| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|     `segments` | required | List of strings  | Названия сегментов   
|     `user_id` | required | Int  | Id user'а

#### Пример:

Input:
```json
{
  "segments": [
    "avito_segment",
    "old_segment"
  ],
  "user_id": 2
}
```

Output:
```json
{
	"Segments added"
}
```


#### Удаление пользователя из сегмента:
Endpoint:
<summary><code>POST</code> <code>http://127.0.0.1:3000/delete_user_from_segment</code></summary>

**Parameters**

|          Name | Required |  Type   | Description                                                                                                                                                           |
| -------------:|:--------:|:-------:| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|     `segments` | required | List of strings  | Названия сегментов   
|     `user_id` | required | Int  | Id user'а

#### Пример:

Input:
```json
{
  "segments": [
    "avito_segment",
    "old_segment"
  ],
  "user_id": 2
}
```

Output:
```json
{
	"User deleted"
}
```

#### Добавление сегмента:
Endpoint:
<summary><code>POST</code> <code>http://127.0.0.1:3000/create_segment</code></summary>

**Parameters**

|          Name | Required |  Type   | Description                                                                                                                                                           |
| -------------:|:--------:|:-------:| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|     `name` | required | string  | Название сегмента   

#### Пример:

Input:
```json
{
	"name": "golang_segment"
}
```

Output:
```json
{
	"id": 4,
	"name": "golang_segment"
}
```

#### Удаление сегмента:
Endpoint:
<summary><code>POST</code> <code>http://127.0.0.1:3000/delete_segment</code></summary>

**Parameters**

|          Name | Required |  Type   | Description                                                                                                                                                           |
| -------------:|:--------:|:-------:| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|     `name` | required | string  | Название сегмента   

#### Пример:

Input:
```json
{
	"name": "golang_segment"
}
```

Output:
```json
{
	"Segment deleted"
}
```

#### Просмотр активных сегментов у пользователя:
Реализован просмотр именно активных сегментов (тех, у которых deleted_at = NULL)
Endpoint:
<summary><code>GET</code> <code>http://127.0.0.1:3000/show_segments/:user<int></code></summary>

