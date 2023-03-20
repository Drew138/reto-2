.PHONY: install create-certs renew-certs

install:
    @sudo apt-get update
    @sudo apt-get install -y certbot python3-certbot-nginx

create-certs:
    @sudo certbot --nginx -d example.com -d www.example.com

renew-certs:
    @sudo certbot renew --dry-run

make wordpress:
	@sudo docker-compose -f ./docker-compose.wordpress.yml up

make mysql:
	@sudo docker-compose -f ./docker-compose.mysql.yml up

make nginx:
	@sudo docker-compose -f ./docker-compose.nginx.yml up

make clean:
	@sudo docker-compose down --volumes
	@sudo docker container prune -y
	@sudo docker image prune -y
