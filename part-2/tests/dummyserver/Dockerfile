FROM ruby:2.5.3-alpine

RUN gem install sinatra

WORKDIR /app

ADD url-aggregator-api.rb /app/url-aggregator-api.rb

CMD ruby /app/url-aggregator-api.rb
