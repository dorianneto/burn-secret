cert:
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout certs/localhost.key -out certs/localhost.crt -subj "/C=CA/ST=Newfoundland/L=St. John\'s/O=Dorian Neto/OU=Development/CN=localhost/emailAddress=doriansampaioneto@gmail.com"
