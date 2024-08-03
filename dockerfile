# Start with Node.js for the frontend
FROM node:18 AS frontend

WORKDIR /app/fe

COPY fe/package*.json ./

RUN npm ci

COPY fe/svelte.config.js ./
COPY fe/vite.config.ts ./
COPY fe/tsconfig.json ./
COPY fe/src/ ./src/
COPY fe/static/ ./static/

RUN npm run build
RUN npm prune --production

# Build the backend
FROM golang:alpine AS backend

WORKDIR /app/pubdeskmd1

COPY pubdeskmd1/go.mod pubdeskmd1/go.sum ./
RUN go mod download

COPY pubdeskmd1/ .
RUN go build -o ./main .

# Final stage
FROM alpine:latest

# Install Node.js
RUN apk add --no-cache nodejs npm

# Install dependencies for Node.js
RUN apk add --no-cache libc6-compat

# Copy frontend files
WORKDIR /app/fe
COPY --from=frontend /app/fe ./

# Copy backend binary
WORKDIR /app/pubdeskmd1
COPY --from=backend /app/pubdeskmd1/main ./

# Expose ports (adjust as needed)
EXPOSE 4173 4000

# Start both services
CMD ["sh", "-c", "cd /app/fe && npm run start & cd /app/pubdeskmd1 && ./main"]