#!/bin/sh

# Create Rabbitmq admin user
ADMIN_USER=${ADMIN_USER:-admin}
ADMIN_PASSWORD=${ADMIN_PASSWORD:-admin}

( rabbitmqctl wait --timeout 60 $RABBITMQ_PID_FILE ; \
rabbitmqctl add_user $ADMIN_USER $ADMIN_PASSWORD 2>/dev/null ; \
rabbitmqctl set_user_tags $ADMIN_USER administrator ; \
rabbitmqctl set_permissions -p / $ADMIN_USER  ".*" ".*" ".*" ; \
echo "*** User '$ADMIN_USER' with password '$ADMIN_PASSWORD' completed. ***") &

# Create Rabbitmq server user
SERVER_USER=${SERVER_USER:-server}
SERVER_PASSWORD=${SERVER_PASSWORD:-server}

( rabbitmqctl wait --timeout 60 $RABBITMQ_PID_FILE ; \
rabbitmqctl add_user $SERVER_USER $SERVER_PASSWORD 2>/dev/null ; \
rabbitmqctl set_user_tags $SERVER_USER administrator ; \
rabbitmqctl set_permissions -p / $SERVER_USER  ".*" ".*" ".*" ; \
echo "*** User '$SERVER_USER' with password '$SERVER_PASSWORD' completed. ***") &

# $@ is used to pass arguments to the rabbitmq-server command.
# For example if you use it like this: docker run -d rabbitmq arg1 arg2,
# it will be as you run in the container rabbitmq-server arg1 arg2
rabbitmq-server $@
