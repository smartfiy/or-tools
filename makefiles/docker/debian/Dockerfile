# Create a virtual environment with all tools installed
# ref: https://hub.docker.com/_/debian
FROM debian:latest AS base
LABEL maintainer="corentinl@google.com"
# Install system build dependencies
ENV PATH=/usr/local/bin:$PATH
RUN apt-get update -qq \
    && apt-get install -yq \
    git wget libssl-dev pkg-config build-essential \
    autoconf libtool zlib1g-dev lsb-release \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Install CMake 3.18.5
RUN wget "https://cmake.org/files/v3.18/cmake-3.18.5-Linux-x86_64.sh" \
    && chmod a+x cmake-3.18.5-Linux-x86_64.sh \
    && ./cmake-3.18.5-Linux-x86_64.sh --prefix=/usr/local/ --skip-license \
    && rm cmake-3.18.5-Linux-x86_64.sh
CMD [ "/bin/bash" ]

# Install SWIG 4.0.1 (currently testing)
FROM base AS swig
RUN echo 'Package: * \nPin: release a=stable \nPin-Priority: 900' > /etc/apt/preferences.d/stable.pref \
    && echo 'Package: * \nPin: release a=testing \nPin-Priority: 400' > /etc/apt/preferences.d/testing.pref \
    && mv /etc/apt/sources.list /etc/apt/sources.list.d/stable.list \
    && echo 'deb http://deb.debian.org/debian testing main \ndeb http://deb.debian.org/debian testing-updates main' \
    > /etc/apt/sources.list.d/testing.list
RUN apt-get update -qq \
    && apt-get install -yq swig/testing \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
