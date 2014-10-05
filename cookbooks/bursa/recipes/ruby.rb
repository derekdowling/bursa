apt_repository "ruby-ng" do
    uri "http://ppa.launchpad.net/brightbox/ruby-ng/ubuntu"
    distribution 'precise'
    components ['main']
    keyserver "keyserver.ubuntu.com"
    key "C3173AA6"
    action :add
end

package "ruby1.9.3" do
      action :upgrade
end

package "rubygems" do
      action :upgrade
end
