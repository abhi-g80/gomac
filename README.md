MAC utilies in Golang
=====================
![go build](https://github.com/abhi-g80/gomac/actions/workflows/go.yml/badge.svg)

Exposes metrics collected from [powermetrics](https://www.unix.com/man-page/osx/1/powermetrics/) as endpoints.

For testing,

    sudo go run main.go

Get CPU and GPU temperature,

    curl localhost:8080/cpu/temperature
    curl localhost:8080/gpu/temperature 

