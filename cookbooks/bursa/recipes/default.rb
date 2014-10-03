#
# Cookbook Name:: bursa
# Recipe:: default
#
# Copyright (c) 2014 The Authors, All Rights Reserved.
#
include_recipe "golang"
include_recipe "bursa-gpm"

default["bursa"]["path"] = "/bursa"
default["bursa"]["gopath"] = "#{default["bursa"]["path"]}/app"

env "GOROOT" do
    action :create
    value "#{default["bursa"]["path"]}"
end

# Appends the value to the end with the delimiter in between
env "PATH" do
    action :modify
    value "#{default["bursa"]["path"]}/bin"
    delim ":"
end
