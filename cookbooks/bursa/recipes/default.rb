#
# Cookbook Name:: bursa
# Recipe:: default
#
# Copyright (c) 2014 The Authors, All Rights Reserved.
#
include_recipe "golang"
include_recipe "bursa-gpm"
include_recipe "bursa-duojs"
include_recipe "bursa::ruby"
include_recipe "bursa-bitcoincore"

# FUTURE ELASTICSEARCH SERVER
# Until we start provisioning different types of production servers, this will be
# sufficient. Otherwise, these cookbooks are better suited to specific run lists
# that are determined by the role played by node.
include_recipe "java"
include_recipe "elasticsearch"

# FUTURE POSTGRESQL SERVER
include_recipe "database"
include_recipe "database::postgresql"
include_recipe "postgresql::server"

# Create our PGDB and create a user
postgresql_database "bursa" do
  connection node["bursa"]["pg_sa"]
  action [:create]
end

postgresql_database_user node["bursa"]["pg_user"]["username"] do
  connection node["bursa"]["pg_sa"]
  database_name "bursa"
  password node["bursa"]["pg_user"]["password"]
  privileges [:all]
  action [:create, :grant]
end

# Setup System Environment Vars
magic_shell_environment "GOPATH" do
    value "#{node["bursa"]["web_root"]}"
end

magic_shell_environment "PATH" do
  value "$PATH:#{node["bursa"]["web_root"]}/bin"
end

gem_package "bundler" do
  action :install
end

#Start our webserver
# service "goserver" do
  # pattern "goserver"
  # supports :start => true
  # init_command "#{node["bursa"]["gopath"]}/src go run main.go&"
  # action [:enable,:start]
# end

# # Enable GoConvey Auto Testing on localhost:8181 if not in Production
# service "gotests" do
  # pattern "gotests"
  # supports :start => true
  # init_command
  # case node["bursa"]["testing"]
  # when "enabled"
    # action [:enable, :start]
  # end
# end

# NODE (NPM, ReactJS)
include_recipe "nodejs"
include_recipe "nodejs::npm"

nodejs_npm "gulp"
nodejs_npm "coffee-react"
nodejs_npm "duo"

# BASIC TOOLS
include_recipe "vim"
