# Making Bursa Work

### Running The Website

Assuming you have already installed chef/kitchen below:

  * To launch the web server
  
		go run src/bursa.io/server/server.go or ./start-server
    
  * To run all the tests:
    
    	go test

### Running the Convey Test Server

  * Simply run the script

		./test-server

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
