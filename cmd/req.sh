#!/usr/bin/sh

register_user() {
  curl -X POST http://localhost:1169/api/register-user \
  -H "Content-Type: application/json" \
  --data-binary "@cmd/user_req.json"
}

login_user() {
  curl -s  -X POST http://localhost:1169/api/login-user \
       -H "Content-Type: application/json" \
       -d '{
             "email": "test@test",
             "password": "test"
           }'
}

get_auction_cards() {
  curl -X GET 'http://localhost:1169/api/auctions?limit=17&offset=0&min_price=1000&max_price=2000&category=electronics&lotcondition=used'
}

get_full_auction() {
  curl -X GET 'http://localhost:1169/api/auction/e7a1d0b2-eef7-49aa-b1a2-38b001b874d6'
}

get_auction_table() {
  token="$(login_user | jq -r '."auth_token"')"
  curl -H "Authorization: Bearer $token" \
       -X GET 'http://localhost:1169/api/get-lots-table?limit=17&offset=0&max_price=3000&min_price=2000'
}

place_bid() {
  token="$(login_user | jq -r '."auth_token"')"
  curl -H "Authorization: Bearer $token" \
       -X POST 'http://localhost:1169/api/post-bid' \
        -d "{
          \"auction_id\": \"e7a1d0b2-eef7-49aa-b1a2-38b001b874d6\",
          \"amount\": $1
        }"
}

#register_user | jq '.'
#login_user | jq '."auth_token"'

#time get_auction_cards | jq '.'
#time get_auction_table | jq '.'

place_bid 5000 | jq '.'
#get_full_auction | jq '.'
