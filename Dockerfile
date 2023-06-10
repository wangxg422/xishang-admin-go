FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk update --no-cache 
    && apk add --update gcc g++ libc6-compat 
    && apk add --no-cache ca-certificates
    && apk add --no-cache tzdata
    && mkdir /backend

ENV TZ Asia/Shanghai

COPY ./backend /backend/
COPY ./config.yaml /backend/config.yaml

EXPOSE 8000
RUN  chmod +x /backend/backend

CMD ["/backend/backend","-c", "/backend/config.yaml"]