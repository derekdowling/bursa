default["bursa"]["path"] = "/bursa"
default["bursa"]["gopath"] = "#{node["bursa"]["path"]}/app"
default["bursa"]["testing"] = "enabled"

# BURSA-DUOJS
default["bursa"]["user"]["homedir"] = "/home/vagrant"
default["bursa"]["github"]["login"] = "bursa-io"
default["bursa"]["github"]["password"] = "67ffb19b463e20d76d6e5f2968ed423a44c54ff5"

# BURSA POSTGRES
default["bursa"]["pg_user"] = {
  :host => "localhost",
  :port => 5432,
  :username => 'bursa',
  :password => "securemebaby"
}
