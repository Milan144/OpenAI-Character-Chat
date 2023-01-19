# Install go dependencies
cd server
go install
# Print if the installation was successful
if [ $? -eq 0 ]; then
    echo "✅ GO dependencies installed successfully"
else
    echo "❌ Installation failed"
fi
# Install pip dependencies
cd ../client
pip install -r requirements.txt
# Print if the installation was successful
if [ $? -eq 0 ]; then
    echo "✅ Pip dependencies installed successfully"
else
    echo "❌ Installation failed"
fi
