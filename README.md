# Tickets-API

## _What is this?_  

A RESTful API written in Golang to store and retrieve ticket information.

This is a RESTful API built on top of the gin-gonic/gin package that uses a MySQL database in order to store and retrieve ticket information, though it can also be used for other types of data. 

![](mysql.png)

## Installation

To test out the API, clone this repository using `git clone https://github.com/komi-alasse/tickets-api.git`

Note: In order to use this API you must have access to a MySQL server, you can install MySQL using Homebrew `brew install mysql` or through downloading [MySQL](https://www.mysql.com/) you will also need [Golang](https://go.dev/) which you can install using `brew install go`.

## Usage

From there you can simply run `go run .` from the project repository to initiate a server which will run from your local device on port 3306. 

At any point you can enter `Ctrl-C` to stop the process.



