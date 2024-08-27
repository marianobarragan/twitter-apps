# Create user
curl --request POST \
  --url http://localhost:8083/users \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/9.3.3' \
  --header 'x-auth-token: 1234' \
  --data '{
	"name": "user name",
	"followings": [
		{
			"user_id": 222
		},
		{
			"user_id": 333
		}
	],
	"followers": [
		{
			"user_id": 444
		},
		{
			"user_id": 555
		}
	]
}'

# GET user
curl --request GET \
  --url http://localhost:8083/users/1724718790043 \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/9.3.3' \
  --header 'x-auth-token: 1234'

# POST tweet
curl --request POST \
  --url http://localhost:8082/tweets \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/9.3.3' \
  --header 'x-auth-token: 1234' \
  --data '{
	"text": "some text"
}'

# GET user timeline
curl --request GET \
  --url 'http://localhost:8081/users/444/timeline?from=1&to=1824631501' \
  --header 'User-Agent: insomnia/9.3.3'
