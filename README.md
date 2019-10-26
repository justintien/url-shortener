# url-shortener

URL Shortener with mysql support.  

Just practicing the implementation of various languages.

- go
- java
- nodejs

## Example

Only the ports are different.

- go: 3001
- java: 3002
- nodejs: 3003

> test server alive

```curl
curl -X GET http://localhost:3001
```

> generate short url

```sh
curl -X POST -H "Content-Type:application/json" -d "{\"url\": \"http://www.google.com\"}" http://localhost:3001/shorten
# Response:
# {"id":1,"url":"http://www.google.com","shortId":"K234GEDO8716","createdAt":"2019-10-23T10:15:28.000+0000","updatedAt":"2019-12-31T10:15:28.000+0000","deletedAt":null}
```

> Redirect short url to origin url

```text
browser: http://localhost:8081/K234GEDO8716
```
