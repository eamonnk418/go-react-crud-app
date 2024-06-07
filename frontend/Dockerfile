FROM node:20.14.0-bookworm-slim AS builder

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

RUN corepack enable pnpm && corepack use pnpm@9.2.0

USER node
WORKDIR /app

COPY --chown=node:node package.json pnpm-lock.yaml ./
RUN --mount=type=cache,id=pnpm-store,target=/pnpm/store pnpm install --frozen-lockfile

COPY --chown=node:node . ./
RUN pnpm run build

FROM node:20.14.0-bookworm-slim

ENV NODE_ENV=production

USER node
WORKDIR /app

COPY --chown=node:node --from=builder /app/.next/standalone ./
COPY --chown=node:node --from=builder /app/.next/static ./.next/static
COPY --chown=node:node --from=builder /app/public ./public

EXPOSE 3000

CMD ["node", "server.js"]
