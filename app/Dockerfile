FROM node:20-alpine3.17 as build

RUN apk update && apk upgrade && adduser -D svelte
USER svelte

WORKDIR /app
COPY --chown=svelte:svelte . /app

RUN npm install && npm run build
CMD ["npm", "run", "preview"]
