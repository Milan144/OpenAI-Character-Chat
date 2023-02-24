# Install go dependencies
cd server
go install
# Print if the installation was successful
if [ $? -eq 0 ]; then
    echo "✅ GO dependencies installed successfully"
else
    echo "❌ Installation failed"
fi

# Install node dependencies
cd ../client
npm install
# Print if the installation was successful
if [ $? -eq 0 ]; then
    echo "✅ Node dependencies installed successfully"
else
    echo "❌ Installation failed"
fi
