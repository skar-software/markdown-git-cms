export $(grep -v '^#' .env.dev | xargs -0)
cd pubdeskmd1
go run .
