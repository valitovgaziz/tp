# Используем образ Node.js для разработки и сборки приложения
FROM node:lts-alpine as builder

# Устанавливаем зависимости
WORKDIR /app
COPY package*.json ./
RUN npm install

# Копируем все файлы проекта
COPY . .

# Собираем статические файлы
RUN npm run build

# Создаем второй этап сборки для минимизации размера образа
FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]