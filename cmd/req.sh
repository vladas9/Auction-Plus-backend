#!/usr/bin/sh

register_user() {
  curl -X POST http://localhost:1169/api/users/register \
  -H "Content-Type: application/json" \
  --data-binary "@cmd/user_req.json"
}

get_auctions() {
  curl -X GET 'http://localhost:1169/api/auctions?limit=17&offset=0&min_price=700&max_price=2000&category=electronics&lotcondition=new'
}

#register_user
time get_auctions | jq '.'
