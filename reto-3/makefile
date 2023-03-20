.PHONY: install create-certs renew-certs wordpress mysql nginx

install:
	@sudo apt-get update
	@sudo apt-get install -y certbot python3-certbot-nginx

create-certs:
	@sudo certbot --nginx -d example.com -d www.example.com

renew-certs:
	@sudo certbot renew --dry-run

wordpress:
	@sudo docker-compose -f ./docker-compose.wordpress.yml up

mysql:
	@sudo docker-compose -f ./docker-compose.mysql.yml up

nginx:
	@sudo docker-compose -f ./docker-compose.nginx.yml up

