FROM nvidia/cuda:9.0-base-ubuntu16.04
LABEL maintainer="William Minshew <w@emrys.io>"

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    build-essential \
    cuda-command-line-tools-9-0 \
    cuda-cublas-9-0 \
    cuda-cufft-9-0 \
    cuda-curand-9-0 \
    cuda-cusolver-9-0 \
    cuda-cusparse-9-0 \
    # curl \
    libcudnn7=7.0.5.15-1+cuda9.0 \
    # libfreetype6-dev \
    # libpng12-dev \
    # libzmq3-dev \
    # pkg-config \
    # python \
    # python-dev \
    # rsync \
    # software-properties-common \
    # unzip \
    python3-pip \
    && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN pip3 install virtualenv
