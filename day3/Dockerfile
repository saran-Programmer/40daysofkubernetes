# Builder stage: has the full Go toolchain, used only to compile the binary.
FROM golang:1.23-alpine AS builder

# All subsequent paths are relative to this directory inside the build container.
WORKDIR /app

# Copy only go.mod first so the dependency layer is cached independently of source changes.
COPY go.mod ./

# Download dependencies now, as their own layer, so this step is skipped on rebuilds unless go.mod changes.
RUN go mod download

# Now bring in the rest of the source, since dependencies are already resolved.
COPY . .

# Build a fully static binary: CGO_ENABLED=0 removes the libc dependency, GOOS=linux targets the scratch image.
RUN CGO_ENABLED=0 GOOS=linux go build -o day3-todo .

# Final stage: scratch has no OS, shell, or libraries, keeping the image minimal and reducing attack surface.
FROM scratch

# Bring in nothing but the compiled binary from the builder.
COPY --from=builder /app/day3-todo /day3-todo

# Run as a non-root numeric UID, since scratch has no /etc/passwd to resolve a named user against.
USER 10001

# Document the port the app listens on.
EXPOSE 8080

# Exec-form CMD runs the binary directly as PID 1, without a shell that scratch doesn't have.
CMD ["/day3-todo"]
