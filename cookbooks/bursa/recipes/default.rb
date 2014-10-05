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

# BURSA APP AND GO ENV
path = ENV['PATH']

# Appends the value to the end with the delimiter in between
bash "PATH" do
    environment "PATH" => "#{path}:#{node["bursa"]["gopath"]}"
end

bash "build" do
    action :run
    cwd "#{node["bursa"]["gopath"]}/src"
    command "./build-script"
end

# NODE (NPM, ReactJS)
include_recipe "nodejs"
include_recipe "nodejs::npm"

nodejs_npm "react-tools"
