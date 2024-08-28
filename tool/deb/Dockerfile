# Use a base image with the necessary tools for building RPM packages
FROM centos:7

# Install required packages for building RPMs
RUN yum install -y rpm-build rpmdevtools

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . /app

# Build the RPM package
RUN rpmbuild -bb /app/path/to/specfile.spec

# Add any additional steps or configurations here

# Set the entrypoint command
CMD ["/bin/bash"]