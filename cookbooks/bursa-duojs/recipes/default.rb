#
# Cookbook Name:: bursa-duojs
# Recipe:: default
#
# Copyright (c) 2014 The Authors, All Rights Reserved.
#
include_recipe "nodejs::npm"

# Install Duo
nodejs_npm "gulp"
nodejs_npm "duo"

# Configures a Github API token for private repos and higher througput.
template "#{node['bursa']['user']['homedir']}/.netrc" do
  source "netrc.erb"
end
