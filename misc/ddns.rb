require 'yaml'
require 'net/http'
require 'json'

# The idea here is to make this a cron job that updates dnsimple periodically.
# It uses a bursa.local.yml file located in your project root to get the config
# variables, an example is provided.

# Load local config data.
bursa = YAML.load_file(File.expand_path('../../bursa.local.yml', __FILE__))
# Get our ip.
ip = JSON.parse(Net::HTTP.get('jsonip.com','/'))["ip"]

# LOGIN="admin@bursa.io"
# TOKEN="your-api-token"
# DOMAIN_ID="bursa.io"
# RECORD_ID="12345" # Replace with the Record ID
# IP="`curl http://jsonip.com | sed 's/{"ip":"\(.*\)"/\1/g' | sed 's/}//'`"

uri = URI.parse("https://api.dnsimple.com/v1/domains/#{bursa["devdns"]["domain_id"]}/records/#{bursa["devdns"]["record_id"]}")
http = Net::HTTP.new(uri.host, uri.port)
http.use_ssl = true
http.verify_mode = OpenSSL::SSL::VERIFY_NONE

req = Net::HTTP::Put.new(uri.request_uri)

req['Accept'] = 'application/json'
req['Content-Type'] = 'application/json'
req['X-DNSimple-Token'] = "#{bursa["devdns"]["login"]}:#{bursa["devdns"]["token"]}"
req.body = JSON.dump({ :content => ip })
resp = http.request req
puts resp
