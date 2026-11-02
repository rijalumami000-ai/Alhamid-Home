# Multi-stage Dockerfile for Svelte Frontend
FROM node:22-alpine AS builder

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci

COPY . .
RUN npm run build

FROM node:22-alpine AS runner

WORKDIR /app
ENV NODE_ENV=production

COPY package.json package-lock.json ./
RUN npm ci --omit=dev

COPY --from=builder /app/.svelte-kit ./.svelte-kit
COPY --from=builder /app/build ./build 2>/dev/null || true
COPY --from=builder /app/static ./static

EXPOSE 3000
ENV PORT=3000

CMD ["node", "build/index.js"]

