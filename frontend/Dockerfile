# Build stage
FROM node:20-alpine AS build
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build && ls -la /app/dist

# Production stage
FROM node:20-alpine
WORKDIR /app
RUN mkdir -p dist
COPY --from=build /app/dist ./dist
COPY --from=build /app/package*.json ./
RUN npm install --production
EXPOSE 3000
CMD ["npm", "run", "preview"]
