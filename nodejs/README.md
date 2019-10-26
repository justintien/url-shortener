# node-url-shortener

> use koa2 + graphql + mysql built a url shortened service as skeleton sample.

## Usage

- up server

  ```sh
  # should automatically install `npm modules`
  docker-compose up -d
  ```

- init
  
  ```sh
  # db migration
  docker-compose exec app npm run db:migrate
  ```

## Example

> test server alive

```curl
curl -X GET http://localhost:4000
```

> generate short url

```sh
## graphql api
curl -X POST \
  http://localhost:4000/graphql \
  -H 'Content-Type: application/json' \
  -d '{"query": "mutation { shorten (url: \"https://baidu.com\") { id url shortId createdAt updatedAt deletedAt } }"
}'

# restful api
curl -X POST -H "Content-Type:application/json" -d "{\"url\": \"http://www.google.com\"}" http://localhost:4000/shorten
# Response:
# {"id":1,"url":"http://www.google.com","shortId":"K234GEDO8716","createdAt":"2019-10-23T10:15:28.000+0000","updatedAt":"2019-12-31T10:15:28.000+0000","deletedAt":null}
```

> Redirect short url to origin url

```text
browser: http://localhost:4000/K234GEDO8716
```
