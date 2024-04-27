package docker

type Dockerfile struct {
	Node            string
	NodeWithDataEnv string
	FE              string
}

const nodeBase = `FROM node:16-alpine

WORKDIR /app

COPY . .

RUN mkdir config

EXPOSE 3000

`

var DockerfileStr = Dockerfile{
	Node: nodeBase + "CMD cd config && node ../sub-store.bundle.js",

	NodeWithDataEnv: nodeBase + "CMD SUB_STORE_DATA_BASE_PATH=/app/config node sub-store.bundle.js",

	FE: `FROM debian:bullseye-slim AS downloader

WORKDIR /app
RUN apt-get update && \
    apt-get install -y curl unzip && \
    rm -rf /var/lib/apt/lists/* && \
    curl -LJO https://sub-store-org.github.io/resource/ssm/nginx.conf

FROM nginx:alpine AS runner

WORKDIR /app
COPY dist.zip /app/
RUN unzip dist.zip -d /app && mv /app/dist /app/www && rm dist.zip
COPY --from=downloader /app/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]`,
}
