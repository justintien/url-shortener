# go-url-shortener

A golang URL Shortener with mysql support.  
<!-- Using Bijective conversion between natural numbers (IDs) and short strings -->

## Usage

```sh
docker-compose up --build
```

## Example

```sh
curl -X POST -H "Content-Type:application/json" -d "{\"url\": \"http://www.google.com\"}" http://localhost:3001/shorten
```

Response:

```sh
{"id":1,"url":"http://www.google.com","shortId":"K234GEDO8716","createdAt":"2019-10-23T10:15:28.000+0000","updatedAt":"2019-12-31T10:15:28.000+0000","deletedAt":null}
```

Redirect

```sh
curl -v localhost:3001/K234GEDO8716
```
