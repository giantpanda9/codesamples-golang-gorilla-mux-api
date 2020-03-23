# codesamples-golang-gorilla-mux-api
Golang API application dependencies: gorilla/mux (http://github.com/gorilla/mux), git
# Description
Simple API on Go lang code example - made using gorilla/mux middleware framework - http://github.com/gorilla/mux
Almost the same as https://github.com/giantpanda9/codesamples-golang-gorilla-mux-api , but contains no Marhalling,
which makes the output JSON look more correct. Let's keep this project to demonstrate the ability to do Marshalling.
# Purposes
To demonstrate ability to code on Go lang and to study this great language a little more
# Requirements
1) gorilla/mux - http://github.com/gorilla/mux
2) git
3) codesamples-python-scrapy-scrap-golang-site - https://github.com/giantpanda9/codesamples-python-scrapy-scrap-golang-site
# Installation instructions (approximate, not the last ones to follow):
On Ubuntu Ubuntu 18.04.4 LTS
1) sudo apt-get install git
2) go get 'github.com/gorilla/mux'
3) clone this project
4) cd to /your/git/folder/this/project/folder
5) go run goApi.go
6) [Optional] Clone and run codesamples-python-scrapy-scrap-golang-site - https://github.com/giantpanda9/codesamples-python-scrapy-scrap-golang-site - if you wish to update data in output.json file
# How to run?
1) http://localhost:8000/goapi/getdata
2) http://localhost:8000/goapi/getdata/rsa - instead of rsa you can input probably any module name from here: https://golang.org/pkg/, if Python scrapper - https://github.com/giantpanda9/codesamples-python-scrapy-scrap-golang-site - re-scrapped the page propely you will get the new modules otherwise only the modules and version from data/output.json file (available as part of this project in data folder)
3) http://localhost:8000/goapi/getdata/bydate/February%202019 - should display Go 1.12 version available since February 2019
# Notes
1) Example input data conatined within data/output.json file
2) ouput.json named like so because it actually an output file of the project - codesamples-python-scrapy-scrap-golang-site - https://github.com/giantpanda9/codesamples-python-scrapy-scrap-golang-site
3) project kept with simple structure to allow quick overview - I know MVC pattern, but shall demonstrate this abilty in dedicated project, if you do not mind - and this is solely to demonstrate the API creation abilities
