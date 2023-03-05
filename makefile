
generate_range_certs:
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -subj /CN=localhost \
		-keyout ./range/playbooks/files/nginx.key -out ./range/playbooks/files/nginx.crt

setup_range:
	generate_range_certs