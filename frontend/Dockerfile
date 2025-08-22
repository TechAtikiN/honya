# Use Node.js base image
FROM node:20-alpine

# Install pnpm globally
RUN npm install -g pnpm

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and pnpm-lock.yaml if it exists
COPY package*.json ./
COPY pnpm-lock.yaml* ./

# Install dependencies using pnpm
RUN pnpm install --frozen-lockfile

# Copy the rest of the application code
COPY . .

# Build the Next.js app (this creates the .next folder)
RUN pnpm run build

# Expose the port the app will run on
EXPOSE 3000

# Start the Next.js app in production mode
CMD ["pnpm", "start"]