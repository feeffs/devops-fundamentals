FROM nginx:alpine
ARG GREETING="Hello, Docker!"
COPY index.html /usr/share/nginx/html/index.html
RUN sed -i "s|__GREETING__|${GREETING}|g" /usr/share/nginx/html/index.html