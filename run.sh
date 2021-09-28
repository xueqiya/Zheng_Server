docker build -t zheng_server . && \
docker run -d -p 8888:8888 --rm --name zheng_server zheng_server