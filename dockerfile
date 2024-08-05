# Start with Node.js for the frontend
FROM node:18 AS frontend

WORKDIR /app/frontend

COPY frontend/package*.json ./

RUN npm ci

COPY frontend/svelte.config.js ./
COPY frontend/vite.config.ts ./
COPY frontend/tsconfig.json ./
COPY frontend/src/ ./src/
COPY frontend/static/ ./static/

RUN npm run build
RUN npm prune --production

# Build the backend
FROM golang:alpine AS backend

WORKDIR /app/backend

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ .
RUN go build -o ./main .

# Final stage
FROM alpine:latest

# Install Node.js
RUN apk add --no-cache nodejs npm

# Install dependencies for Node.js
RUN apk add --no-cache libc6-compat

# Copy frontend files
WORKDIR /app/frontend
COPY --from=frontend /app/frontend ./

# Copy backend binary
WORKDIR /app/backend
COPY --from=backend /app/backend/main ./

# Expose ports (adjust as needed)
EXPOSE 4173 4000

# Start both services
CMD ["sh", "-c", "cd /app/frontend && npm run start & cd /app/backend && ./main"]