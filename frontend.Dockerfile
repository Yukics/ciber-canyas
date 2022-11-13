FROM node:18 as builder

WORKDIR /app

COPY frontend/ .

RUN npm install && npm run build


FROM nginx 

COPY --from=builder /app/dist/ /usr/share/nginx/html