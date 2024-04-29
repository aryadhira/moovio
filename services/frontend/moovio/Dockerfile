FROM node:18-alpine

WORKDIR /app

COPY package.json .

RUN npm install

COPY . .

RUN npm run build

CMD [ "npm", "run", "start" ]

# RUN chmod +x docker-entrypoint.sh

# ENTRYPOINT [ "/bin/sh", "docker-entrypoint.sh" ]