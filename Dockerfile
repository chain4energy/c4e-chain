# Use the golang image as a starting point.
FROM golang:1.19

# Set the working directory inside the container.
WORKDIR /app

# Copy local code to the container image.
COPY . .

RUN curl https://get.ignite.com/cli@v0.26.1 | bash
RUN mv ignite /usr/local/bin/
# Open the required ports for the blockchain.
EXPOSE 26657 1317

CMD ["ignite", "chain", "serve", "--reset-once", "-v"]
