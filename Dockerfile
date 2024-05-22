# Base image
FROM python:3-slim AS build-env
COPY . /app
WORKDIR /app

# Distroless image
FROM gcr.io/distroless/python3
COPY --from=build-env /app /app
WORKDIR /app
CMD ["calc.py", "/etc"]
