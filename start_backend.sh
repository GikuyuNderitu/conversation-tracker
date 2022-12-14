#! /bin/sh
# trap "exit" INT TERM ERR
trap "kill 0" EXIT
export DB_PATH="file:///data/surrealdb/test"
export POCKET_BASE_ADDRESS="127.0.0.1:1732"

# surreal start --log debug --user root --pass root -b 0.0.0.0:9021 &

# ./pocketbase serve --http $POCKET_BASE_ADDRESS &

cp ./.env ./service/notes/cmd/conversation_app

go run ./service/notes/cmd/conversation_app --usepostgres &


go run ./service/notes/cmd/conversation_app_gateway &
wait

rm ./service/notes/cmd/conversation_app/.env