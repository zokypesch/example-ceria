# Example for ceria API WRAPPER

# Just register your struct / model, and BOOOOMMM ! your api is already automaticlly

`CD/CI with jenkins and kubernetes soon will be updated`

# please install docker & docker-compose first 
if you dont have postgresql, rabbitMQ, elastic and redis in your local machine
<a href="https://docs.docker.com/install/"> https://docs.docker.com/install/ </a>

`
Note: elastic version ^5 please set your memory minim 4GB in your virtual machine
if you want to skip elastic just change the status in config.development.yaml to false
`

# run docker-compose 
`cd docker && docker-compose up --build`
Run the examples


# Install Dep tools
`
see the instruction
<a href="https://github.com/golang/dep">https://github.com/golang/dep</a>
`

# Install ceria package
`
dep ensure
`

# Run example apps
```
cd examples/wrapper && go run main.go
cd examples/rabbitMQ && go run main.go
```

# Open your Postman or using Curl (i expect your using examples/wrapper)
```
Get Data
curl -H "Accept: application/json" -X GET http://localhost:9090/articles?page=1&limit=30

Response

{"data":[{"Author":{"Fullname":"udin"},"AuthorID":"2","Body":"lorem ipsum lorem ipsum lorem ipsum","Comments":[],"Model":{"CreatedAt":"2019-01-09 05:57:16.017519 +0000 UTC","DeletedAt":"","ID":"1","UpdatedAt":"2019-01-09 05:57:16.017519 +0000 UTC"},"Tag":"#thanks","Title":"Hi there iam fine thanks"}],"error":"","message":"","page":"1","status":true,"total_data":"1"}%

Insert with inheritance
curl -i -H "Accept: application/json"  -X POST http://localhost:9090/articles -d \
'{"title": "hello welcome to ceria", "tag": "#Ceriaworkspace", "body": "lorem ipsum lorem ipsum", "Author": {"fullname": "Ceria Lover"}, "Comments": [{"fullname": "udin", "body": "top numero uno"},{"fullname": "ucok", "body": "bahh layyy"}]}'

Response
{"data":{"ID":14,"CreatedAt":"2019-01-09T21:49:50.911122+07:00","UpdatedAt":"2019-01-09T21:49:50.911122+07:00","DeletedAt":null,"title":"hellowelcome to ceria","tag":"#Ceriaworkspace","body":"lorem ipsum lorem ipsum","author":{"ID":5,"CreatedAt":"2019-01-09T21:49:50.90907+07:00","UpdatedAt":"2019-01-09T21:49:50.90907+07:00","DeletedAt":null,"fullname":"Ceria Lover"},"author_id":5,"Comments":[{"ID":2,"CreatedAt":"2019-01-09T21:49:50.915091+07:00","UpdatedAt":"2019-01-09T21:49:50.915091+07:00","DeletedAt":null,"article_id":14,"fullname":"udin","body":"top numero uno"},{"ID":3,"CreatedAt":"2019-01-09T21:49:50.920292+07:00","UpdatedAt":"2019-01-09T21:49:50.920292+07:00","DeletedAt":null,"article_id":14,"fullname":"ucok","body":"bahh layyy"}]},"error":"","message":"","status":true}%

Normal insert
curl -i -H "Accept: application/json"  -X POST http://localhost:9090/articles -d \
'{"title": "hello welcome to ceria", "tag": "#Ceriaworkspace", "body": "lorem ipsum lorem ipsum", "author_id": 1}'

Response
{"data":{"ID":15,"CreatedAt":"2019-01-09T21:51:10.295422+07:00","UpdatedAt":"2019-01-09T21:51:10.295422+07:00","DeletedAt":null,"title":"hellowelcome to ceria","tag":"#Ceriaworkspace","body":"lorem ipsum lorem ipsum","author":{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"fullname":""},"author_id":1,"Comments":null},"error":"","message":"","status":true}

Update
curl -i -H "Accept: application/json"  -X PUT http://localhost:9090/articles/14 -d \
'{"data": {"title": "iam in ceria"}}'

Delete
curl -H "Accept: application/json" -X DELETE http://localhost:9090/articles/13

Response
{"data":null,"error":"","message":"","status":true}%

Bulk Delete
curl -H "Accept: application/json" -X POST http://localhost:9090/articles/bulkdelete -d \
'{"data": [{"id": 1},{"id": 2, "title": "iam in ceria"}]}'

Response
{"data":{"data":[{"id":1},{"id":2,"title":"iam in ceria"}]},"error":"","message":"","status":true}%

Bulk Create
curl -i -H "Accept: application/json"  -X POST http://localhost:9090/articles/bulkcreate -d \
'[{"title": "hello welcome to ceria", "tag": "#Ceriaworkspace", "body": "lorem ipsum lorem ipsum", "author_id": 1}]'

response
{"data":[{"ID":16,"CreatedAt":"2019-01-09T21:55:18.507644+07:00","UpdatedAt":"2019-01-09T21:55:18.507644+07:00","DeletedAt":null,"title":"hello welcome to ceria","tag":"#Ceriaworkspace","body":"lorem ipsum lorem ipsum","author":{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"fullname":""},"author_id":1,"Comments":null}],"error":"","message":"","status":true}

Find
curl -i -H "Accept: application/json"  -X POST http://localhost:9090/articles/find -d \
'{"condition": {"author_id": 3}}'

Response
{"data":[{"Author":{"Fullname":""},"AuthorID":"3","Body":"lorem ipsum lorem ipsum lorem ipsum","Comments":[],"Model":{"CreatedAt":"2019-01-09 07:02:19.743132 +0000 UTC","DeletedAt":"","ID":"9","UpdatedAt":"2019-01-09 07:02:19.743132 +0000 UTC"},"Tag":"#thanks","Title":"Hi there iam not fine thanks"}],"error":"","message":"","status":true}%

Login to get JWT
curl -i -H "Content-Type: application/json"  -X POST http://localhost:9090/login -d \
'{"username": "admin","password": "admin"}'

Response
{"code":200,"expire":"2019-01-09T22:59:51+07:00","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDcwNDk1OTEsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU0NzA0NTk5MX0.T3OV4vnnGFJwwkmbkxLIlpMBUfpBJzswd8Tu3wxfaFs"}

get token and accept to barrier
curl -i -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDcwNDk1OTEsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU0NzA0NTk5MX0.T3OV4vnnGFJwwkmbkxLIlpMBUfpBJzswd8Tu3wxfaFs"  -X GET http://localhost:9090/auth/comments 

Response
{"data":[{"ArticleID":"12","Body":"hai udin great !","Fullname":"ajinomoto","Model":{"CreatedAt":"2019-01-09 07:08:53.706371 +0000 UTC","DeletedAt":"","ID":"1","UpdatedAt":"2019-01-09 07:08:53.706371 +0000 UTC"}},{"ArticleID":"14","Body":"top numero uno","Fullname":"udin","Model":{"CreatedAt":"2019-01-09 14:49:50.915091 +0000 UTC","DeletedAt":"","ID":"2","UpdatedAt":"2019-01-09 14:49:50.915091 +0000 UTC"}},{"ArticleID":"14","Body":"bahh layyy","Fullname":"ucok","Model":{"CreatedAt":"2019-01-09 14:49:50.920292 +0000 UTC","DeletedAt":"","ID":"3","UpdatedAt":"2019-01-09 14:49:50.920292 +0000 UTC"}}],"error":"","message":"","status":true}

check your elastic search
curl -i -H "Content-Type: application/json"  -X POST http://localhost:9200/articles/_search -d '{"id": "14"}'

for query using expression like
curl -H "Accept: application/json" -X GET http://localhost:9090/articles?page=1&limit=1&where=title:welcome:LIKE|author_id:1:EQUAL

Response
{"data":[{"Author":{"Fullname":"admin"},"AuthorID":"1","Body":"lorem ipsum lorem ipsum lorem ipsum","Comments":[],"Model":{"CreatedAt":"2019-01
-09 06:24:29.73612 +0000 UTC","DeletedAt":"","ID":"8","UpdatedAt":"2019-01-09 06:24:29.73612 +0000 UTC"},"Tag":"#thanks","Title":"Hi there iam
not fine thanks"}],"error":"","message":"","page":"1","status":true,"total_data":"1"}

```
