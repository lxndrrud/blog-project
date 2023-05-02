FROM node:14.13.1
WORKDIR /var/www/api
COPY ./ /var/www/api/
RUN npm install
RUN npm run build
ENTRYPOINT ["npm", "start"]
