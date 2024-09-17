register_user() {
  curl -X POST http://localhost:1169/api/users/register \
  -H "Content-Type: application/json" \
  --data-binary "@user_req.json"
}

get_auctions() {
  curl -s -X GET 'http://localhost:1169/api/auctions?limit=2&offset=0'
}

#register_user
get_auctions | jq '.'
