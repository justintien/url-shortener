# url-shortener

URL Shortener with mysql support.  

Just practicing the implementation of various languages.

- go
- java
- nodejs

## Usage

- up server

  ```sh
  ./build.java.sh # build java first
  # up server
  docker-compose up --build -d
  ```

## Api

Only the ports are different.

- go: 3001
- java: 3002
- nodejs: 3003

### test server alive

```sh
curl -X GET http://localhost:3001
# response: server alive!
```

### 縮網址

```http
POST /shorten
```

#### 参数

Attribute | Type | Required | Description
--- | --- | --- | ---
`url` | url | yes | 要縮短的原網址

> example:

```sh
curl -X POST \
-H "Content-Type:application/json" \
-d "{\"url\": \"http://www.google.com\"}" \
http://localhost:3001/shorten
```

##### response

```http
200 OK
```

```json
{
    "id": 1,
    "url": "http://www.google.com",
    "shortId": "k3mau1pvbILo",
    "createdAt": "2019-10-23T10:15:28.000+0000",
    "updatedAt": "2019-12-31T10:15:28.000+0000",
    "deletedAt": null
}
```

or

```http
400 Bad Request
```

```json
{
    "error": "Invalid url"
}
```

### redirect 縮網址

```http
GET /{shortId}
```

#### 参数

Attribute | Type | Required | Description
--- | --- | --- | ---
`{shortId}` | String | yes | 縮網址ID

> example:

```sh
curl -X GET \
http://localhost:3001/k3mau1pvbILo
```

##### response

```http
303 See Other
```

```text
redirect to: http://www.google.com
```

or

```http
404 Not Found
```

```json
{
    "error": "Not Found"
}
```
