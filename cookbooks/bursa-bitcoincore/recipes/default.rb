#
# Cookbook Name:: bursa-bitcoincore
# Recipe:: default
#
# Copyright (c) 2014 The Authors, All Rights Reserved.
#
include_recipe "apt"

apt_repository "bitcoin" do
   uri "http://ppa.launchpad.net/bitcoin/bitcoin/ubuntu "
   distribution node['lsb']['codename']
   components ["main"]
   keyserver "keyserver.ubuntu.com"
   key "8842CE5E"
end

apt_package "bitcoind"

template "/etc/init/bitcoin-server.conf"

directory "#{node['bursa']['user']['homedir']}/.bitcoin" do
  owner node['bursa']['user']['name']
  group node['bursa']['user']['name']
end

template "#{node['bursa']['user']['homedir']}/.bitcoin/bitcoin.conf" do
  owner node['bursa']['user']['name']
  group node['bursa']['user']['name']
end

service "bitcoin-server" do
  provider Chef::Provider::Service::Upstart
  action :start
end

magic_shell_alias "rt" do
  command 'bitcoin-cli -regtest'
end
