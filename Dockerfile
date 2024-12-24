# Use the latest Golang image to build SFS
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container
COPY . .

# Fetch and tidy up dependencies for SFS
RUN go mod tidy

# Build the application using a Makefile
RUN make build

# Use the latest Alpine image for a smaller runtime container
FROM alpine:latest AS runtime

# Install necessary dependencies for the runtime environment
RUN apk add --no-cache libc6-compat bash

# Set the working directory for the runtime container
WORKDIR /root/

# Copy the compiled SFS application from the builder container to the runtime container
COPY --from=builder /app/dist/sfs /root/

# Copy the public directory (e.g., static assets) from the builder container to the runtime container
COPY --from=builder /app/public /root/public/

# Copy the custom entrypoint script from the builder container to the runtime container
COPY --from=builder /app/scripts/docker_entrypoint.sh /usr/bin/docker_entrypoint.sh

# Expose port 3000 to the outside world
EXPOSE 3000

# Set the entrypoint to the bash shell that runs the entrypoint script
CMD ["/bin/bash", "/usr/bin/docker_entrypoint.sh"]
