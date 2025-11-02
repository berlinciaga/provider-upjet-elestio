# ABOUTME: Dockerfile for building the Elestio Crossplane provider image
# ABOUTME: Creates a minimal runtime image with the provider binary
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY _output/bin/linux_arm64/provider /usr/local/bin/provider
USER 65532:65532

ENTRYPOINT ["provider"]
