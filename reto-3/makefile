.PHONY: install create-certs renew-certs

install:
    sudo apt-get update
    sudo apt-get install -y certbot python3-certbot-nginx

create-certs:
    sudo certbot --nginx -d example.com -d www.example.com

renew-certs:
    sudo certbot renew --dry-run


