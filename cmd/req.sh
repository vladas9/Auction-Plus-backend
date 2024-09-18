#!/usr/bin/sh

register_user() {
  curl -X POST http://localhost:1169/api/users/register \
  -H "Content-Type: application/json" \
  --data-binary "@cmd/user_req.json"
}

get_auctions() {
  curl -X GET 'http://localhost:1169/api/auctions?limit=15&offset=3&min_price=100&max_price=500'
}

#register_user
time get_auctions | jq '.'
