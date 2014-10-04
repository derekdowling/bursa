#
# Cookbook Name:: bursa
# Recipe:: default
#
# Copyright (c) 2014 The Authors, All Rights Reserved.
#
include_recipe "golang"
include_recipe "bursa-gpm"

# FUTURE ELASTICSEARCH SERVER
# Until we start provisioning different types of production servers, this will be
# sufficient. Otherwise, these cookbooks are better suited to specific run lists
# that are determined by the role played by node.
include_recipe "java"
include_recipe "elasticsearch"

# FUTURE POSTGRESQL SERVER
include_recipe "postgresql"
include_recipe "postgresql::server"

# BURSA APP AND GO ENV
path = ENV['PATH']

bash "GOROOT" do
    environment "GOROOT" => node["bursa"]["path"]
end

# Appends the value to the end with the delimiter in between
bash "PATH" do
    environment "PATH" => "#{path}:#{node["bursa"]["gopath"]}"
end

# NODE (NPM, ReactJS)
include_recipe "nodejs"
include_recipe "nodejs::npm"

nodejs_npm "react-tools"
