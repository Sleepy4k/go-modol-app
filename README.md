# Golang MODOL (Modular Online Learning) APP

## About The Project

MODOL is an e-learning application designed to provide a modular learning experience, where each module can stand alone or be integrated with other modules.
The application is built using the Go programming language for efficiency and high performance.

## Disclaimer

**IMPORTANT:**

- This project is for educational purposes
- Any names or assets in this project only used as data display
- **I don't take any responsibility for what you do with this program in the future**

## Overview

### Interface

- This application is designed to be Console Line Interface (CLI)
- User can access and use the application easily
- Our application provided clear command to help user solving their module step by step

### Storage

- Your data will be stored at local disk, maybe cloud storage if we need it
- Any progress had you done will be saved too, so you can continue your progress
- Altough you change your directory we will restore default data from the application

## Getting Started

Before you install this project on your local machine or somewhere else,
you need to install golang interpreter from their website and make sure
to test it, whether it work or have some problem

## Pre-requisites

Here are things for you to run this program:

- Golang interpreter
- Bash Script

## Installation

### Using Git

- Clone this project from github using git

~~~bash
git clone https://github.com/sleepy4k/go-modol-app.git
~~~

- Go to the project directory

~~~bash
cd go-modol-app
~~~

### Using Zip File

- Download from github website using this url

~~~bash
https://github.com/Sleepy4k/go-modol-app/archive/refs/heads/main.zip
~~~

- Unzip the zip file

## Build Project

### Using Bash Script

- Run shell script (please open it from bash terminal to prevent error)

~~~bash
./build.sh
~~~

### Using Command Line

- Run this command on console line (terminal)

~~~bash
go build
~~~

- Run the built app

~~~bash
./modol-app
~~~

## Conclusion

From this project we learn how to manage content like moodle application, such as
submission, subject, and quiz. And we learn how validation so important for the
app, because it will prevent as from errors and security issues. And for the last
but not least, we learn how we store data on json file and manipulate it using
file stream so we get relationship data and much more.

## Feedback

If you have any feedback for me or this project, feel free to contact me or make
an issue with detail description and also tag me 😁
