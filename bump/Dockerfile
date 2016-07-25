FROM iron/ruby:dev

WORKDIR /app
RUN gem install --no-ri --no-rdoc bump

ENTRYPOINT ["bump", "--no-commit", "--no-bundle"]
