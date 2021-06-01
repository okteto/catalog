# syntax = docker/dockerfile:experimental
FROM node:12 as builder
RUN mkdir /usr/src/app
WORKDIR /usr/src/app

ENV PATH /usr/src/app/node_modules/.bin:$PATH
ENV YARN_CACHE_FOLDER=/root/.cache/yarn/v6
COPY package.json yarn.lock ./
RUN --mount=type=cache,target=/root/.cache/yarn/v6 yarn install --frozen-lockfile

COPY . .
RUN --mount=type=cache,target=.cache yarn run build

FROM nginx:1.19-alpine
COPY --from=builder /usr/src/app/dist /usr/share/nginx/html
COPY nginx.default.conf /etc/nginx/conf.d/default.conf
EXPOSE 8080
CMD ["/bin/sh", "-c", "exec nginx -g 'daemon off;'"]