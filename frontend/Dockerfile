FROM node:22.2.0 AS BuildStage

WORKDIR /frontend

COPY package*.json ./

# Install dependencies
RUN npm install

COPY . .

# Build the application
RUN npm run build

#Deploy Stage
FROM node:22.2.0-alpine AS secretnoteserverfront

RUN npm install -g http-server

WORKDIR /frontend

COPY --from=BuildStage /frontend/dist .

EXPOSE 5173

# Start the server
CMD ["http-server", "-p", "5173"]
