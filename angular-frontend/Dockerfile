FROM node:20-slim AS builder

WORKDIR /app

ARG NG_APP_AUTH_SERVICE_ENDPOINT
ARG NG_APP_BLOGS_SERVICE_ENDPOINT

ENV NG_APP_AUTH_SERVICE_ENDPOINTT=$NG_APP_AUTH_SERVICE_ENDPOINT
ENV NG_APP_BLOGS_SERVICE_ENDPOINT=$NG_APP_BLOGS_SERVICE_ENDPOINT

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

# EXPOSE 4200
# CMD ["npm", "start"]