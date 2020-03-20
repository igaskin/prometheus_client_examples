FROM ruby:2.5

# throw errors if Gemfile has been modified since Gemfile.lock
RUN bundle config --global frozen 1

WORKDIR /usr/src/app

COPY Gemfile Gemfile.lock ./
RUN bundle install

COPY . .

ENV PORT=8000

CMD ["bundle", "exec", "unicorn", "-p", "8000", "-c", "/usr/src/app/unicorn.conf"]
