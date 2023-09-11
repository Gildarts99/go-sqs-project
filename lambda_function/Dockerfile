FROM golang:1.20 as build

WORKDIR /go-sqs-lambda

# Copy dependencies list
COPY go.mod go.sum ./

# Build with optional lambda.norpc tag, this reduces the size of the image if you are not using RPC
COPY . ./
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o main main.go

# Copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /go-sqs-lambda/main ./main
ENTRYPOINT [ "./main" ]