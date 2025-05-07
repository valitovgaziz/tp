# Используем официальный образ Nginx
FROM nginx:alpine

# Копируем index.html в папку Nginx
COPY index.html /usr/share/nginx/html/

# (Опционально) Можно заменить конфиг Nginx
# COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf

# Порт, который будет слушать Nginx
EXPOSE 80