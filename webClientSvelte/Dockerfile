FROM node:latest

WORKDIR /usr/src/app

COPY package.json ./
COPY yarn.lock* ./

RUN yarn install

COPY . .

RUN yarn build

EXPOSE 8081

CMD ["yarn", "dev", "--host"]
