FROM node:20-alpine3.17 as build

RUN apk update && apk upgrade 

WORKDIR /app
COPY . /app

EXPOSE 4173 

ARG API_URL
ENV VITE_API_URL_DEV $API_URL 

RUN npm install && npm run build
CMD ["npm", "run", "preview"]
