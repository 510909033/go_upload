FROM registry.plt.babytree-inc.com/base/centos:7.6.1810

COPY . /var/www

EXPOSE 9601 

ENV GORACE="log_path=/tmp/golang_race/report"

RUN mkdir -p /opt/go/logs \
	&& mkdir -p /tmp/golang_race

#RUN /var/www/go_upload/run
