FROM node:8.11.0

WORKDIR /usr/src/app

COPY package*.json .
RUN npm install --only=prod

COPY . .

EXPOSE 9012

CMD [ "node", "index.js" ]
