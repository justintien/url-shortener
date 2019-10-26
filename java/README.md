# java-url-shortener

A java URL Shortener with mysql support.  
<!-- Using Bijective conversion between natural numbers (IDs) and short strings -->

## Example

```sh
curl -X POST -H "Content-Type:application/json" -d "{\"url\": \"http://www.google.com\"}" http://localhost:3002/shorten
```

Response:

```sh
{"id":1,"url":"http://www.google.com","shortId":"K234GEDO8716","createdAt":"2019-10-23T10:15:28.000+0000","updatedAt":"2019-12-31T10:15:28.000+0000","deletedAt":null}
```

Redirect

```sh
curl -v localhost:8081/K234GEDO8716
```
