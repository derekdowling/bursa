#
# Cookbook Name:: bursa
# Recipe:: default
#
# Copyright (c) 2014 The Authors, All Rights Reserved.
#
include_recipe "golang"
include_recipe "bursa-gpm"
include_recipe "bursa-duojs"

# FUTURE ELASTICSEARCH SERVER
# Until we start provisioning different types of production servers, this will be
# sufficient. Otherwise, these cookbooks are better suited to specific run lists
# that are determined by the role played by node.
include_recipe "java"
include_recipe "elasticsearch"

# FUTURE POSTGRESQL SERVER
include_recipe "database"
include_recipe "postgresql"
include_recipe "postgresql::ruby"
include_recipe "postgresql::server"


postgresql_database 'bursa' do
  connection(
    :host => '127.0.0.1',
    :port => 5432,
    :username => 'bursa',
    :password => 'securemebaby'
  )
  action :create
end

# Setup System Environment Vars
bash "GOPATH" do
    environment "GOPATH" => node["bursa"]["gopath"]
end

path = ENV['PATH']
bash "PATH" do
    environment "PATH" => "#{path}:#{node["bursa"]["gopath"]}"
end

# Compile Various Website Bits
gem_package "bundler" do
  action :install
end

bash "build" do
    action :run
    cwd "#{node["bursa"]["gopath"]}/src"
    command "./build-script"
end

# Start our webserver
service "goserver" do
  pattern "goserver"
  supports :start => true
  init_command "#{node["bursa"]["gopath"]}/src go run main.go&"
  action [:enable,:start]
end

# Enable GoConvey Auto Testing on localhost:8181 if not in Production
service "gotests" do
  pattern "gotests"
  supports :start => true
  init_command "#{node["bursa"]["gopath"]}/src go test -port=8181 &"
  case node["bursa"]["testing"]
  when "enabled"
    action [:enable, :start]
  end
end

# NODE (NPM, ReactJS)
include_recipe "nodejs"
include_recipe "nodejs::npm"

nodejs_npm "react-tools"
