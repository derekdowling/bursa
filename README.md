# Bursa

![Build Status](https://magnum.travis-ci.com/derekdowling/bursa.svg?token=iq92sEsRxzbuqGK3drsX&branch=master)

### Installing the Backend

  * Clone the repository into your go/src/github.com directory

	* CD into /bursa and run "godep restore" you're now ready to develop

### Running The Website

Assuming you have already installed chef/kitchen below:

  * To launch the web server

				go run src/bursa.io/server/server.go or ./start-server

  * To run all the tests:

				go test

### Running the Convey Test Server

  * Simply run the script

				./test-server
