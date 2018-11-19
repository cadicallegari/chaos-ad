FROM python:3.7-alpine3.8

# RUN apk add --update && \
#     apk add ca-certificates gcc && \
#     apk del make g++ && \
#     rm /var/cache/apk/*

COPY ./requirements.txt /app/requirements.txt
RUN pip install -r /app/requirements.txt

COPY ./chaosad_py /app/chaosad_py

ENV PYTHONPATH /app
ENV PYTHONDONTWRITEBYTECODE 1

WORKDIR /app
