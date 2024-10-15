FROM node:20-slim AS builder

WORKDIR /app
COPY package.json /app
COPY package-lock.json /app

RUN npm install
RUN npm audit fix --force

COPY . .
RUN npm run build --prod

FROM nginx:latest

COPY --from=builder /app/dist/angular-frontend /usr/share/nginx/html
COPY nginx.conf  /etc/nginx/conf.d/default.conf

EXPOSE 80