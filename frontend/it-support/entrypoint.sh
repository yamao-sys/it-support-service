#!/bin/bash

pnpm i
pnpm playwright install --with-deps

pnpm gen:api-spec:watch &
pnpm dev
