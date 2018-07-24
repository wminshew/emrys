FROM nvidia/cuda:9.0-base-ubuntu16.04 as base
LABEL maintainer="William Minshew <w@emrys.io>"

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    # build-essential \
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
    apt-get autoremove && \
    rm -rf /var/lib/apt/lists/*

RUN pip3 install virtualenv
# RUN virtualenv venv
# ARG REQS
# COPY $REQS $REQS
# RUN ./venv/bin/pip install -r $REQS

# FROM nvidia/cuda:9.0-base-ubuntu16.04 as runtime
# LABEL maintainer="William Minshew <w@emrys.io>"
#
# RUN apt-get update && \
#     apt-get install -y --no-install-recommends \
#     python3 \
#     && \
#     apt-get clean && \
#     rm -rf /var/lib/apt/lists/*

ARG HOME
RUN useradd --create-home --home-dir $HOME user \
  && chown -R user:user $HOME
WORKDIR $HOME
# COPY --from=base --chown=user:user /venv ./venv
RUN virtualenv venv
ARG REQS
COPY $REQS $REQS
RUN ./venv/bin/pip --no-cache-dir install -r $REQS
USER user

ARG MAIN
ENV MAIN $MAIN
COPY --chown=user:user $MAIN $MAIN
CMD ./venv/bin/python $MAIN
