# Making Bursa Work

### Running The Website

Assuming you have already installed chef/kitchen below:
	
  * Go to Bursa Root, then set Go Paths:
  
		source gvp

  * Install Go Deps:
  
    	gpm install
  
  * Install node depencies()frontend asset management)
    
		npm install

  * Build our assets
  
		gulp css:vendor
    	gulp css:main

   	Also - Live compile assets with:
    
	    gulp watch

  * To launch the web server
  
		go run src/bursa.io/server/server.go
    
  * To run all the tests:
    
    	go test bursa-io

### Running the Convey Test Server

  * Build convey (from Bursa root):
	
		go build smartystreets/goconvey -o .godeps/bin/goconvey \ github.com/smartystreets/goconvey

  * Run Convey:
	
		./.godeps/bin/goconvey

### Chef/Kitchen Requirements

* Install the chef development kit, which has tools to replace the vagrant berkshelf plugin:

		wget https://opscode-omnibus-packages.s3.amazonaws.com/mac_os_x/10.8/x86_64/chefdk-0.3.0-1.dmg
* Install neccessary gems in bursa root:
 
  		bundle install
* Within the cookbooks/bursa directory:

	**NOTE** It'd be nice to have this right in the root. [There's a known issue](https://github.com/opscode/chef-dk/issues/50)

	* Install cookbook dependencies for our **application cookbook**:
			
			berks install

	* Create the VM:
	
			kitchen create all
			
	* Provision the VM:
			
			kitchen converge
	  
	* Login to the kitchen-generated vm:
	  
	  		kitchen login
	* See if gpm is available:
			
			gpm --version
   
### References

  
* [Getting Started with Chef DK](http://tcotav.github.io/chefdk_getting_started.html)

### Cookbook References

* https://github.com/phlipper/chef-postgresql
