FROM nginx:latest


LABEL maintainer="Hasan Karademir"
LABEL version="1.0"
LABEL name="frontend-app"

RUN apt-get update && \
    apt-get install --no-install-recommends curl -y && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /usr/share/nginx/html

COPY . /usr/share/nginx/html
# COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 CMD curl -f http://localhost/ || exit 1

# ENTRYPOINT ["/bin/bash", "-c", "export IP=$(nslookup backend | grep Address: | tail -n 1 | awk '{print $2}'); echo 'Backend IP: '$IP; exec \"$@\""]

CMD nginx -g 'daemon off;'
# ENTRYPOINT [ "supervisord" ]