FROM alpine:3.14
LABEL maintainer="AF <343663190@qq.com>" version="1.0" license="MIT"
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories  && apk update
ENV TZ=PRC
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
WORKDIR .
RUN touch .env
COPY ./goApp ./
CMD ["./goApp"]