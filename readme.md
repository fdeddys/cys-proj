## Information

### database
create database resourceDb;


### table
CREATE TABLE users (
	id bigserial NOT NULL,
	user_name varchar(100) NULL,
	api_token varchar(200) NULL,
	aws_credential varchar(200) NULL,
	PRIMARY KEY (id)
);

CREATE TABLE resources (
	id bigserial NOT NULL,
	user_name varchar(100) NULL,
	resource_type varchar(200) NULL,
	service_name varchar(200) NULL,
	PRIMARY KEY (id)
);


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
