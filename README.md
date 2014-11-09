#HelpTheSheriff webapp
Test of web application written in Go and deployed (temporarily) on Heroku.
https://helpthesheriff.herokuapp.com/.

I am sorry for this crappy name (fileserver), but it was a different app at the beginning, but I was too lazy to create a new one.

## Prerequisites
* [Go latest version](https://golang.org/doc/install): the backend is written in Go. I think it's compatible also with older versions, but there's no reason to not use the latest one.
* [Heroku toolbelt](https://toolbelt.heroku.com/): to push to heroku server using Git.
* [Account on Heroku](http://heroku.com/): to create a new app.

## Install
* [install and configure Go](http://golang.org/doc/install) and the [Heroku toolbelt](https://toolbelt.heroku.com/) on your machine
* create an account on Heroku
* fork and clone this repository on your local machine

## Run Locally
In order to run it locally you just need to have Go installed and configured ($GOPATH!!).
* *cd /fileserver*
* *go build*
* *./fileserver* : at this point you should be able to run this webapp locally at [http://localhost:8080](http://localhost:8080).

## Customization
In order to use the names of your teammates, you have to modify the file *hello.go* : 
```
func AssignPTR(rw http.ResponseWriter, r *http.Request) {
    teammate := []Teammate{
        {"Alessandro"},
        {"Davide"},
        {"Marcello"},
        {"Sarah"},
        {"Fred"},
        {"Khaled"},
        {"Herve"},
        {"Bo"},
        {"Remi"},
    }
```
## Deploy on Heroku
* follow the official procedure [here](https://toolbelt.heroku.com/). It's easy as hell.