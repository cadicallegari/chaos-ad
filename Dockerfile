FROM python:3.7-alpine3.8

COPY ./requirements.txt /app/requirements.txt
RUN pip install -r /app/requirements.txt

COPY ./mypkg /app/mypkg

ENV PYTHONPATH /app
ENV PYTHONDONTWRITEBYTECODE 1

WORKDIR /app
