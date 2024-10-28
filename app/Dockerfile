FROM node:22.9.0 as build

WORKDIR /app
COPY .  /app

ARG API_URL
ENV VITE_API_URL_DEV $API_URL

RUN npm install && npm run build

FROM oven/bun:1.1.33 as main

WORKDIR /app

COPY --from=build /app/build ./build
COPY --from=build /app/package.json ./package.json
COPY --from=build /app/package-lock.json ./package-lock.json

EXPOSE 4173

RUN bun install

CMD ["bun", "build/index.js"]
