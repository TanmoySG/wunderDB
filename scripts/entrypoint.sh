# echo "Setting Up Server Configurations - server-config.json"
# read -p "Enter PORT: " PORT
# read -p "Enter MAIL SERVER: " SERVER
# read -p "Enter SENDER: " SENDER
# read -p "Enter PASSWORD: " PASSWORD

SERVER_CONFIG={\"port\":\"$PORT\",\"mail-server\":\"$SERVER\",\"password\":\"$PASSWORD\",\"sender\":\"$SENDER\"} 

echo $SERVER_CONFIG > /app/server-config.json

python app.py