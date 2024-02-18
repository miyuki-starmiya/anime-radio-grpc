FROM python:3.11-slim

ENV TZ=Asia/Tokyo

# Install dependencies
RUN apt-get update && \
    pip install poetry

WORKDIR /app

COPY pyproject.toml poetry.lock ./

# Create virtualenv at .venv in the project instead of ~/.cache/
# Install production dependencies 
RUN poetry config virtualenvs.in-project true && \
    poetry install --no-dev --no-ansi

COPY . .

EXPOSE 8080

CMD ["poetry", "run", "python", "main.py"]
