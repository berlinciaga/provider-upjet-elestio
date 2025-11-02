# ABOUTME: Dockerfile for building the Elestio Crossplane provider image
# ABOUTME: Creates a minimal runtime image with the provider binary
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY provider /usr/local/bin/
USER 65532:65532

ENTRYPOINT ["provider"]
