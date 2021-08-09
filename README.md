MAC power utilies in Golang
===========================
![build](https://github.com/abhi-g80/gomac/actions/workflows/build.yml/badge.svg)
![test](https://github.com/abhi-g80/gomac/actions/workflows/test.yml/badge.svg)

Exposes metrics collected from [powermetrics](https://www.unix.com/man-page/osx/1/powermetrics/) as endpoints.

For testing,

    sudo go run main.go

Get CPU and GPU temperature,

    curl localhost:8080/smc/cpu/temperature
    curl localhost:8080/smc/gpu/temperature 

CPU temperature can be seen at, [http://localhost:8080/graphs/temp.html](http://localhost:8080/graphs/temp.html).
