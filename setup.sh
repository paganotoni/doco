# This script downloads the tailwindcss binary and renames it to 
# bin/tailwindcss.
echo "ℹ️  Downloading TailwindCSS binary. Please wait..."
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64

echo "✅ TailwindCSS binary downloaded."
mkdir -p bin
chmod +x tailwindcss-macos-arm64
mv tailwindcss-macos-arm64 bin/tailwindcss




