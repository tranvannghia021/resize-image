FROM ubuntu:22.04
    
RUN mkdir /app
            
COPY . /app
    
ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install -y \
            build-essential \
            wget \
            git \
            pkg-config
    
RUN apt-get install -y python3 python3-pip python3-setuptools \
                           python3-wheel
RUN pip3 install meson ninja
    
RUN apt-get install -y \
            libexpat1-dev \
            librsvg2-dev \
            libpng-dev \
            libjpeg-dev \
            libwebp-dev \
            libexif-dev \
            liblcms2-dev \
            libglib2.0-dev \
            liborc-dev \
            libgirepository1.0-dev \
            gettext 
    
ARG VIPS_VER=8.14.2
ARG VIPS_DLURL=https://github.com/libvips/libvips/releases/download
RUN cd /usr/local/src \
            && wget ${VIPS_DLURL}/v${VIPS_VER}/vips-${VIPS_VER}.tar.xz \
            && tar xf vips-${VIPS_VER}.tar.xz \
            && cd vips-${VIPS_VER} \
            && meson setup build --buildtype=release \
            && cd build \
            && meson compile \
            && meson test \
            && meson install
RUN ldconfig

RUN apt install golang-go && go version \
    && export GOROOT=/usr/local/go \
    && export GOPATH=$HOME/go \
    && export PATH=$GOPATH/bin:$GOROOT/bin:$PATH \
    && go version
            
RUN export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/vips/lib