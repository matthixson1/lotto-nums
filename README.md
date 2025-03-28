# Euromillions Number Generator

A simple command-line tool to generate random Euromillions lottery numbers.

## Description

This application generates random Euromillions lottery lines. Each line consists of:
- 5 main numbers (from 1 to 50)
- 2 star numbers (from 1 to 12)

All numbers are randomly generated and sorted in ascending order.

## Building the Application

To build the application, you need to have Go installed on your system. Then run:

`make`
or
`go build -v -o euromillions main.go Euromillions.go`


and run the application with a specified number of lines (default 3), e.g:
`./euromillions --lines 16`


