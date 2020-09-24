FROM rabbitmq:3-management

# To pause the script execution until the rabbitmq starts up fully.
ENV RABBITMQ_PID_FILE /var/lib/rabbitmq/mnesia/rabbitmq

ADD ./scripts/add_users.sh /init.sh
RUN chmod +x /init.sh

EXPOSE 5672
EXPOSE 15672

# Define default command
CMD ["/init.sh"]
