
#!/usr/bin/env bash

BASE_URL="http://localhost:8080"
EMAIL="test@test.com"
PASSWORD="123456"

echo "===> Register"
curl -s -X POST "$BASE_URL/auth/register" \
  -H "Content-Type: application/json" \
  -d "{
    \"email\": \"$EMAIL\",
    \"username\": \"testuser\",
    \"password\": \"$PASSWORD\"
  }"
echo -e "\n"

echo "===> Login"
TOKEN=$(curl -s -X POST "$BASE_URL/auth/login" \
  -H "Content-Type: application/json" \
  -d "{
    \"email\": \"$EMAIL\",
    \"password\": \"$PASSWORD\"
  }" | jq -r '.access_token')

echo "Token: $TOKEN"
echo ""

if [ "$TOKEN" == "null" ] || [ -z "$TOKEN" ]; then
  echo "❌ Login failed"
  exit 1
fi

AUTH_HEADER="Authorization: Bearer $TOKEN"

echo "===> Create Transaction"
curl -s -X POST "$BASE_URL/transactions" \
  -H "Content-Type: application/json" \
  -H "$AUTH_HEADER" \
  -d '{
    "amount": 150.75,
    "currency": "TRY",
    "category": "food",
    "note": "Lunch",
    "occured_at": "2024-01-10T12:00:00Z"
  }'
echo -e "\n"

echo "===> List Transactions"
curl -s -X GET "$BASE_URL/transactions?limit=10&offset=0" \
  -H "$AUTH_HEADER"
echo -e "\n"

START="2024-01-01T00:00:00Z"
END="2024-12-31T23:59:59Z"

echo "===> Sum By Day"
curl -s -X GET "$BASE_URL/transactions/daily?start=$START&end=$END" \
  -H "$AUTH_HEADER"
echo -e "\n"

echo "===> Sum By Month"
curl -s -X GET "$BASE_URL/transactions/monthly?start=$START&end=$END" \
  -H "$AUTH_HEADER"
echo -e "\n"

echo "===> Sum By Year"
curl -s -X GET "$BASE_URL/transactions/yearly?start=$START&end=$END" \
  -H "$AUTH_HEADER"
echo -e "\n"

echo "✅ All tests executed"
