test:
	go test ./controller/... ./email/... ./firewall/... ./kernel/... ./latinum/... ./middleware/... ./models/... ./picasso/... ./renaissance/... ./server/...
	# Excluded: latinum (need to bootstrap bitcoin)
test-circle:
	go test ./controller/... ./email/... ./firewall/... ./kernel/... ./middleware/... ./models/... ./picasso/... ./renaissance/... ./server/...
	# Excluded: latinum (need to bootstrap bitcoin)
