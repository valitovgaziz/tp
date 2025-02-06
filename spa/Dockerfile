# Используем официальный образ Node.js на основе Alpine Linux
FROM node:16-alpine as build-stage

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем package.json и package-lock.json (если есть)
COPY package*.json ./

# Устанавливаем зависимости
RUN npm install

# Копируем все файлы проекта в рабочую директорию
COPY . .

# Собираем проект
RUN npm run build

# Используем легковесный образ Nginx для serving статики
FROM nginx:stable-alpine as production-stage

# Копируем собранные файлы из предыдущего этапа в директорию Nginx
COPY --from=build-stage /app/dist /usr/share/nginx/html

# Копируем конфигурацию Nginx (если нужно)
# COPY nginx.conf /etc/nginx/conf.d/default.conf

# Открываем порт 80
EXPOSE 80

# Запускаем Nginx
CMD ["nginx", "-g", "daemon off;"]