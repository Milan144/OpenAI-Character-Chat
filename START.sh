# Start go server
cd server
go run app.go &
# Print if server is running or not
if [ $? -eq 0 ]; then
    echo "🚀 GO server is running"
else
    echo "❌ GO server is not running"
fi
