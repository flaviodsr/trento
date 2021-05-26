#!/bin/bash
TARGET_HOST=$1
CONSUL_HOST=$2
ACTION=$3
TRENTO_BIN="trento"
CONSUL_BIN="consul"
TRENTO_PATH="/srv/trento/"
CONSUL_PATH="/srv/consul/"
EXAMPLES_PATH="$TRENTO_PATH/examples"
TARGET_USER="root"
CONSUL_DATA_DIR="consul-agent-data"

# Abort when any command in the script fails
set -e

# Abort if no input params
if [ $# -lt 2 ] ; then
    echo "Usage: ./deploy.sh <target-server-ip> <consul-ip> [deploy-agent*|deploy-web]"
    exit 1
fi

# If no action is provided, we assume deploy-agent
if test "X$ACTION" == "X" ; then
    ACTION="deploy-agent"
fi

stop_process () {
    echo "Checking if process $2 is running in $1..."
    while ssh "$TARGET_USER@$TARGET_HOST" pgrep -x "$2" > /dev/null
    do
        echo "Attempting to stop $2 process on $1..."
        ssh "$TARGET_USER@$TARGET_HOST" killall "$2"
        sleep 2
    done
}

# Stop old processes
stop_process "$TARGET_HOST" "trento"
stop_process "$TARGET_HOST" "consul"

# Create directory structure if it doesn't exist
ssh "$TARGET_USER@$TARGET_HOST" mkdir -p "$TRENTO_PATH" || true
ssh "$TARGET_USER@$TARGET_HOST" mkdir -p "$CONSUL_PATH/consul.d" || true

# Upload new binaries & examples
rsync -av ./$TRENTO_BIN "$TARGET_USER@$TARGET_HOST:/$TRENTO_PATH"
rsync -av ./$CONSUL_BIN "$TARGET_USER@$TARGET_HOST:/$CONSUL_PATH"
rsync -av ./examples "$TARGET_USER@$TARGET_HOST:/$EXAMPLES_PATH"

# Give them execution permission
ssh "$TARGET_USER@$TARGET_HOST" chmod +x "$TRENTO_PATH/$TRENTO_BIN"
ssh "$TARGET_USER@$TARGET_HOST" chmod +x "$CONSUL_PATH/$CONSUL_BIN"

# Start 'em
if [ "$ACTION" = "deploy-agent" ] ; then
	ssh "$TARGET_USER@$TARGET_HOST" -f "nohup $CONSUL_PATH/$CONSUL_BIN agent -bind=$TARGET_HOST -data-dir=$CONSUL_DATA_DIR -config-dir=$CONSUL_PATH/consul.d -retry-join=$CONSUL_HOST -ui > /dev/null 2>&1"
	ssh "$TARGET_USER@$TARGET_HOST" -f "nohup $TRENTO_PATH/$TRENTO_BIN agent start --config-dir=$CONSUL_PATH/consul.d --consul-template=$EXAMPLES_PATH/trento-config.tpl $EXAMPLES_PATH/azure-rules.yaml > /dev/null 2>&1"
elif [ "$ACTION" = "deploy-web" ] ; then
	ssh "$TARGET_USER@$TARGET_HOST" -f "nohup $CONSUL_PATH/$CONSUL_BIN agent -server -bootstrap-expect=1 -bind=$TARGET_HOST -data-dir=$CONSUL_DATA_DIR -ui > /dev/null 2>&1"
	ssh "$TARGET_USER@$TARGET_HOST" -f "nohup $TRENTO_PATH/$TRENTO_BIN web serve > /dev/null 2>&1"
fi
