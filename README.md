# CPU Balance

Monitors core activity to maximize multiprocessing.

![CPU Balance Animation](./img/CPUBalance.gif)

## Purpose

Modern CPUs are usually multi-core processors. However, unless special considerations are made, processes are usually single-threaded. The goal of this app is to allow remote monitoring of CPU activity across multiple cores and see the difference in CPU core activity.

## Technologies Used

1. [Go](https://go.dev/) Programming Language
2. [HTMX](https://htmx.org/) for dynamic HTML
3. [Websocket](http://github.com/coder/websocket) for connection establishment and data updating
4. [gopsutil](https://github.com/shirou/gopsutil/) for obtaining system information
5. [Bootstrap](https://getbootstrap.com/docs/5.3/getting-started/introduction/) for theming and styles
6. [Chart.js](https://www.chartjs.org/docs/latest/) for charting

## Installation

1. Clone the repository
2. Change directory into `cpu-balance` repository folder
3. Run `go mod tidy` to update and cleanup `go.mod` and `go.sum` files
4. Run `go run ./cmd/main.go` to start the program
5. Access the CPU balance dashboard at `[http://localhost:8080/](http://localhost:8080/)`

## Credits

- [sigrdrifa's go-htmx-websockets-example](https://github.com/sigrdrifa/go-htmx-websockets-example) for inspiration