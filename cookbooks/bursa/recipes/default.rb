#
# Cookbook Name:: bursa
# Recipe:: default
#
# Copyright (c) 2014 The Authors, All Rights Reserved.
#
include_recipe "golang"
include_recipe "bursa-gpm"

path = ENV['PATH']

bash "GOROOT" do
    environment "GOROOT" => node["bursa"]["path"]
end

# Appends the value to the end with the delimiter in between
bash "PATH" do
    environment "PATH" => "#{path}:#{node["bursa"]["gopath"]}"
end
