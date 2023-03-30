FROM nginx:stable-alpine3.17-slim


LABEL maintainer="Hasan Karademir"
LABEL version="1.0"
LABEL name="frontend-app"

RUN apk add --no-cache curl && \
    rm -rf /var/cache/apk/*


WORKDIR /usr/share/nginx/html

COPY . /usr/share/nginx/html
# COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 CMD curl -f http://localhost/ || exit 1

# ENTRYPOINT ["/bin/bash", "-c", "export IP=$(nslookup backend | grep Address: | tail -n 1 | awk '{print $2}'); echo 'Backend IP: '$IP; exec \"$@\""]

CMD nginx -g 'daemon off;'
# ENTRYPOINT [ "supervisord" ]