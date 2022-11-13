FROM node:18 as builder

WORKDIR /app

COPY frontend/package.json .

RUN npm install && npm run build


FROM nginx 

COPY --from=builder /app/dist/ /var/www/html/