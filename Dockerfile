# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY ./lab-api-teams ./lab-api-teams
USER 65532:65532

ENTRYPOINT ["./lab-api-teams"]
