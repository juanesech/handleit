FROM alpine

# COPY . /code

RUN apk add yarn nodejs --no-cache

WORKDIR /code
# RUN yarn
ENTRYPOINT [ "yarn", "dev" ]