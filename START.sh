# Start go server
cd server
go run app.go &
# Print if server is running or not
if [ $? -eq 0 ]; then
    echo "🚀 GO server is running"
else
    echo "❌ GO server is not running"
fi
# Start Django app 
cd ../client
python3.9 manage.py runserver 127.0.0.1:5002 &
# Print if Django app is running or not
if [ $? -eq 0 ]; then
    echo "🐍 Django app is running"
else
    echo "❌ Django app is not running"
fi
