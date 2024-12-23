# Stage 1: Builder
# Use the official Go image as the base image for building the application
FROM golang:latest AS builder

# Set the working directory inside the container for the build process
WORKDIR /app

# Copy the entire source code and Makefile from the host into the container
COPY . .

# Fetch the Go dependencies and tidy up the module cache to ensure all dependencies are resolved
RUN go mod tidy

# Run `make build` to compile the Go application (Makefile should define the build process)
RUN make build

# Stage 2: Runtime (Alpine)
# Use the lightweight Alpine Linux image as the base for the runtime environment
FROM alpine:latest AS runtime

# Install the necessary libraries for the Go binary (e.g., to ensure compatibility with glibc)
RUN apk add --no-cache libc6-compat

# Set the working directory inside the container where the final app will run
WORKDIR /root/

# Copy the compiled Go binary from the builder stage into the runtime image
COPY --from=builder /app/dist/sfs .

# Copy the public directory (which contains static files/templates) from the builder stage into the container
COPY --from=builder /app/public /root/public

# Expose the port that the Go application will listen on (adjust if needed)
EXPOSE 3000

# Set the command to run the Go binary (`sfs`) when the container starts
CMD ["./sfs"]
