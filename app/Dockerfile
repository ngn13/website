FROM oven/bun:1.1 as build

WORKDIR /app
COPY . /app

EXPOSE 4173 

ARG API_URL
ENV VITE_API_URL_DEV $API_URL 

RUN bun install && bun run build
CMD ["bun", "run", "preview"]
