#!/usr/bin/sh

register_user() {
  curl -X POST http://localhost:1169/api/users/register \
  -H "Content-Type: application/json" \
  --data-binary "@cmd/user_req.json"
}

get_auction_cards() {
  curl -X GET 'http://localhost:1169/api/auctions?limit=17&offset=0&min_price=0&max_price=2000&category=electronics&lotcondition=used'
}

get_auction_table() {
  curl -X GET 'http://localhost:1169/api/get-lots-table?limit=17&offset=0' # ADD Auth header and test
}

#register_user
time get_auction_cards | jq '.'
