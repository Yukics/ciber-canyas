FROM node:18

WORKDIR /app

# COPY frontend/package.json .

# RUN npm install

CMD ["npm", "run", "dev"]
