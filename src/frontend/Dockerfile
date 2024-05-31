FROM node:18.17.1

RUN mkdir -p /app/atcoder-web-app

WORKDIR /app/atcoder-web-app

COPY package.json package-lock.json ./

RUN npm install

RUN npm install -g next

COPY . .

CMD ["npm", "start"]