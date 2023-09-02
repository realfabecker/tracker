FROM node:18-slim as base

FROM base as dev
WORKDIR /app
COPY package.json .
RUN npm install