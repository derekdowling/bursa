# GO RELATED
default["bursa"]["path"] = "/bursa"
default["bursa"]["web_root"] = "#{node["bursa"]["path"]}"
default["bursa"]["testing"] = "enabled"

# BURSA-DUOJS
default["bursa"]["user"]["name"] = "vagrant"
default["bursa"]["user"]["homedir"] = "/home/vagrant"
default["bursa"]["github"]["login"] = "bursa-io"
default["bursa"]["github"]["password"] = "67ffb19b463e20d76d6e5f2968ed423a44c54ff5"

# BURSA POSTGRES OWNER
# password comes from roles => all.rb
default["bursa"]["pg_sa"] = {
  :host => "127.0.0.1",
  :port => 5432,
  :username => "postgres",
  :password => "sa_securemebaby"
}

# BURSA POSTGRES USER
default["bursa"]["pg_user"] = {
  :username => 'bursa',
  :password => "securemebaby"
}
