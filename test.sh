curl -H "Accept: application/json" -X GET http://localhost:9090/articles?page=1&limit=30

curl -i -H "Accept: application/json"  -X POST http://localhost:9090/articles -d \
'{"title": "hello welcome to ceria", "tag": "#Ceriaworkspace", "body": "lorem ipsum lorem ipsum", "Author": {"fullname": "Ceria Lover"}, "Comments": [{"fullname": "udin", "body": "top numero uno"},{"fullname": "ucok", "body": "bahh layyy"}]}'

curl -i -H "Accept: application/json"  -X POST http://localhost:9090/articles -d \
'{"title": "hello welcome to ceria", "tag": "#Ceriaworkspace", "body": "lorem ipsum lorem ipsum", "author_id": 1}'

curl -i -H "Accept: application/json"  -X PUT http://localhost:9090/articles/14 -d \
'{"data": {"title": "iam in ceria"}}'

curl -H "Accept: application/json" -X DELETE http://localhost:9090/articles/13

curl -H "Accept: application/json" -X POST http://localhost:9090/articles/bulkdelete -d \
'{"data": [{"id": 1},{"id": 2, "title": "iam in ceria"}]}'

curl -i -H "Accept: application/json"  -X POST http://localhost:9090/articles/bulkcreate -d \
'[{"title": "hello welcome to ceria", "tag": "#Ceriaworkspace", "body": "lorem ipsum lorem ipsum", "author_id": 1}]'

curl -i -H "Accept: application/json"  -X POST http://localhost:9090/articles/find -d \
'{"condition": {"author_id": 3}}'

curl -i -H "Content-Type: application/json"  -X POST http://localhost:9090/login -d \
'{"username": "admin","password": "admin"}'

curl -i -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDcwNDk1OTEsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU0NzA0NTk5MX0.T3OV4vnnGFJwwkmbkxLIlpMBUfpBJzswd8Tu3wxfaFs"  -X GET http://localhost:9090/auth/comments 

curl -i -H "Content-Type: application/json"  -X POST http://localhost:9200/articles/_search -d '{"id": "14"}'

curl -H "Accept: application/json" -X GET http://localhost:9090/articles?page=1&limit=1&where=title:welcome:LIKE|author_id:1:EQUAL