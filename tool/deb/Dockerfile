# Use a base image with the necessary tools for building deb packages
FROM ubuntu:latest

# Install build dependencies
RUN apt-get update && apt-get install -y build-essential

# Copy the source code into the container
COPY . /app

# Set the working directory
WORKDIR /app

# Build the deb package
RUN dpkg-buildpackage -us -uc

# Cleanup unnecessary files
RUN rm -rf /app

# Set the entrypoint to the deb package
CMD ["ls", "-l", "*.deb"]