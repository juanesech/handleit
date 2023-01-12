FROM alpine:3

# COPY . /code

RUN apk add yarn nodejs ruby-dev build-base --no-cache

WORKDIR /code
RUN gem install sass
# RUN yarn
ENTRYPOINT [ "yarn", "dev" ]