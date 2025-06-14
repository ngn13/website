# build the application with node
FROM node:23.11.1 AS build

ARG WEBSITE_REPORT_URL
ARG WEBSITE_SOURCE_URL
ARG WEBSITE_DOC_URL
ARG WEBSITE_API_URL
ARG WEBSITE_API_PATH

ENV WEBSITE_REPORT_URL=$WEBSITE_REPORT_URL
ENV WEBSITE_SOURCE_URL=$WEBSITE_SOURCE_URL
ENV WEBSITE_DOC_URL=$WEBSITE_DOC_URL
ENV WEBSITE_API_URL=$WEBSITE_API_URL
ENV WEBSITE_API_PATH=$WEBSITE_API_PATH

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
