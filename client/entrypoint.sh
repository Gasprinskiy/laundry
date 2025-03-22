#!/bin/sh

# Заменяем переменные окружения в конфиге Nginx
envsubst '${API_PROXY_PATH} ${$NGINX_PORT}' < /etc/nginx/conf.d/default.conf.template > /etc/nginx/conf.d/default.conf

# Запускаем Nginx
nginx -g "daemon off;"