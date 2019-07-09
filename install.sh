
if [[ -z $version ]]; then
    version="0.1.15"
fi

filename="fileserver_${version}_${os}_${arch}.tar.gz"
url="https://github.com/broothie/fileserver/releases/download/v$version/$filename"

wget "$url"
tar -C /usr/local/bin/ -xvzf "$filename"
rm "$filename"
