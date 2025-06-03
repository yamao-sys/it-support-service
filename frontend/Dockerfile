FROM node:22-bookworm

WORKDIR /app

COPY . .

RUN rm /bin/sh && ln -s /bin/bash /bin/sh

RUN apt update && apt install default-jre -y

RUN npm install -g pnpm \
    && SHELL=bash pnpm setup \
    && source /root/.bashrc \
    && pnpm install turbo --global \
    && pnpm install -g @typespec/compiler

WORKDIR /app/it-support
