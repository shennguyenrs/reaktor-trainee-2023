FROM node:18-alpine
RUN npm install -g pnpm
RUN mkdir -p /app
WORKDIR app
COPY . .
RUN pnpm install
RUN pnpm run build
ENTRYPOINT ["node", "build"]
