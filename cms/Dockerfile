FROM node:20-alpine as build

WORKDIR /app

COPY package.json /app
COPY yarn.lock /app
COPY nuxt.config.ts /app

RUN yarn set version 1.22.18
RUN yarn install
COPY . .
RUN yarn build

FROM node:20-alpine as product
WORKDIR /app
COPY --from=build /app/.output /app/.output
CMD ["node", ".output/server/index.mjs"] 
