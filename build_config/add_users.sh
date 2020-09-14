#!/bin/sh

ADMIN_USER=${ADMIN_USER:-admin}
ADMIN_PASSWORD=${ADMIN_PASSWORD:-admin}

# Create Rabbitmq user
( rabbitmqctl wait --timeout 60 $RABBITMQ_PID_FILE ; \
rabbitmqctl add_user $ADMIN_USER $ADMIN_PASSWORD 2>/dev/null ; \
rabbitmqctl set_user_tags $ADMIN_USER administrator ; \
rabbitmqctl set_permissions -p / $ADMIN_USER  ".*" ".*" ".*" ; \
echo "*** User '$ADMIN_USER' with password '$ADMIN_PASSWORD' completed. ***" ; \
echo "*** Log in the WebUI at port 15672 (example: http:/localhost:15672) ***") &

# $@ is used to pass arguments to the rabbitmq-server command.
# For example if you use it like this: docker run -d rabbitmq arg1 arg2,
# it will be as you run in the container rabbitmq-server arg1 arg2
rabbitmq-server $@
