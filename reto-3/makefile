.PHONY: install create-certs renew-certs wordpress mysql nginx

install:
	@sudo apt-get update
	@sudo apt-get install -y certbot python3-certbot-nginx

create-certs:
	@sudo certbot --nginx -d reto-3-wordpress.site -d www.reto-3-wordpress.site -d reto-3.reto-3-wordpress.site -d www.reto-3.reto-3-wordpress.site

renew-certs:
	@sudo certbot renew --dry-run

wordpress:
	@sudo docker-compose -f ./docker-compose.wordpress.yml up

mysql:
	@sudo docker-compose -f ./docker-compose.mysql.yml up

nginx:
	@sudo docker-compose -f ./docker-compose.nginx.yml up

mount:
	@mount 10.128.0.13:/var/nfs/general ./wp_data 
