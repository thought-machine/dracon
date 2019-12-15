#!/bin/bash -e

if [[ "${OSTYPE}" == "linux-gnu" ]]; then
    arch="linux"
else
    echo "-> unsupported OS: ${OSTYPE}"
    exit 1
fi
echo "-> determined ${arch} OS"

install_path="${HOME}/.local/bin"
if [ $(whoami) == "root" ]; then
    install_path="/usr/local/bin"
fi
echo "-> installing under ${install_path}"
mkdir -p ${install_path}

bin_path="${install_path}/dracon"
repo="thought-machine/dracon"
download_url=$(curl -s https://api.github.com/repos/${repo}/releases/latest \
  | grep browser_download_url \
  | grep ${arch} \
  | cut -d '"' -f 4)
echo "-> downloading ${download_url} to ${bin_path}"
curl -L $download_url -o ${bin_path}
chmod +x ${bin_path}

echo ""
echo "$PATH"|grep -q ${install_path} || echo "-> You need to add ${install_path} to your PATH variable. e.g. export PATH=${install_path}:\$PATH;"
echo ""

echo "$PATH"|grep -q ${install_path} || export PATH=${install_path}:$PATH
dracon -h
