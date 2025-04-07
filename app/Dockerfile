# build the application with node
FROM node:23.11.0 AS build

# app URLs
ARG WEBSITE_APP_URL_CLEAR
ARG WEBSITE_APP_URL_ONION
ARG WEBSITE_APP_URL_I2P

ENV WEBSITE_APP_URL_CLEAR=$WEBSITE_APP_URL_CLEAR
ENV WEBSITE_APP_URL_ONION=$WEBSITE_APP_URL_ONION
ENV WEBSITE_APP_URL_I2P=$WEBSITE_APP_URL_I2P

# API URLs
ARG WEBSITE_API_URL_CLEAR
ARG WEBSITE_API_URL_ONION
ARG WEBSITE_API_URL_I2P

ENV WEBSITE_API_URL_CLEAR=$WEBSITE_API_URL_CLEAR
ENV WEBSITE_API_URL_ONION=$WEBSITE_API_URL_ONION
ENV WEBSITE_API_URL_I2P=$WEBSITE_API_URL_I2P

# other config
ARG WEBSITE_REPORT_URL
ARG WEBSITE_SOURCE_URL
ARG WEBSITE_DOC_URL

ENV WEBSITE_REPORT_URL=$WEBSITE_REPORT_URL
ENV WEBSITE_SOURCE_URL=$WEBSITE_SOURCE_URL
ENV WEBSITE_DOC_URL=$WEBSITE_DOC_URL

WORKDIR /app
COPY .  /app

RUN apt install -y make sed wget
RUN npm install
RUN make

# run it with bun (a lot faster)
FROM oven/bun:latest AS main

WORKDIR /app

COPY --from=build /app/build             ./build
COPY --from=build /app/package.json      ./package.json
COPY --from=build /app/package-lock.json ./package-lock.json

RUN useradd runner -r -u 1001 -d /app
RUN chown -R runner:runner /app

USER runner
RUN bun install

EXPOSE 7001

ENV PORT=7001
CMD ["bun", "build/index.js"]
