#! /bin/sh
trap "exit" INT TERM ERR
trap "kill 0" EXIT
export DB_PATH=file:///data/surrealdb/test

surreal start --log debug --user root --pass root -b 0.0.0.0:9021 &

go run ./service/notes &

go run ./service/notes_fe &
wait