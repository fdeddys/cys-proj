## Information

### database
#### create database resourceDb;




### Register

curl -X POST \
  http://localhost:8888/api/user/register \
  -H 'content-type: application/json' \
  -d '{
	"username":"deddy",
	"apiToken":"123",
	"awsCredential":"12345"
}'


### Gather data

curl -X POST \
  http://localhost:8888/api/user/gather \
  -H 'authorization: 123' \
  -H 'content-type: application/json' \
  -d '{
	"username":"deddy"
}'


### List resource

curl -X POST \
  http://localhost:8888/api/user/list-resources \
  -H 'authorization: 123' \
  -H 'content-type: application/json' \
  -d '{
	"username":"deddy"
}'


### run

exec run.sh
