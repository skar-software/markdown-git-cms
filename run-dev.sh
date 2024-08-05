export $(grep -v '^#' .env.dev | xargs -0)
cd backend
go run .

