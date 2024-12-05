# Используем Node.js LTS версию в качестве базового образа
FROM node:lts-alpine as builder

# Устанавливаем зависимости
WORKDIR /app
COPY package*.json ./
RUN npm install
RUN npm install axios

# Копируем исходный код
COPY . .

# Строим продакшн-версию приложения
RUN npm run build

# Создаем финальный образ на основе Nginx
FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]