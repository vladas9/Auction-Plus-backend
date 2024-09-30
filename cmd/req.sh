#!/usr/bin/sh

register_user() {
  curl -X POST http://localhost:1169/api/user/register \
  -H "Content-Type: application/json" \
  --data-binary "@cmd/user_req.json"
}

login_user() {
  curl -s  -X POST http://localhost:1169/api/user/login \
       -H "Content-Type: application/json" \
       -d '{
             "email": "test@test",
             "password": "test"
           }'
}

get_auction_cards() {
  curl -X GET 'http://localhost:1169/api/auction/cards?limit=17&offset=0&min_price=1000&max_price=2000&category=electronics&lotcondition=used'
}

get_full_auction() {
  curl -X GET "http://localhost:1169/api/auction/$1"
}

get_auction_table() {
  token="$(login_user | jq -r '."auth_token"')"
  curl -H "Authorization: Bearer $token" \
       -X GET 'http://localhost:1169/api/auctions/table?limit=17&offset=0&max_price=3000&min_price=2000'
}

place_bid() {
  token="$(login_user | jq -r '."auth_token"')"
  curl -H "Authorization: Bearer $token" \
       -X POST 'http://localhost:1169/api/bid/post' \
        -d "{
          \"auction_id\": \"e7a1d0b2-eef7-49aa-b1a2-38b001b874d6\",
          \"amount\": $1
        }"
}

post_auct() {
  token="$(login_user | jq -r '."auth_token"')"
  curl -H "Authorization: Bearer $token" \
       -H "Content-Type: application/json" \
       -X POST 'http://localhost:1169/api/auction/post' \
       -d '{
                "img_src": null,
                "title": "Test item",
                "description": "This is a test item.",
                "category_name": "electronics",
                "condition": "new",
                "start_date": "2024-09-17T12:42:51.788144+03:00",
                "end_date": "2024-09-23T12:42:51.788144+03:00",
                "start_price": "165"
           }'
}

#register_user | jq '.'
#login_user | jq '."auth_token"'

#time get_auction_cards | jq '.'
#time get_auction_table | jq '.'

#place_bid 6650 | jq '.'

post_auct
get_full_auction $id | jq '.'

