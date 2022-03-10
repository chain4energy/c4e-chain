DIR="$(pwd)/$(dirname $0)"
echo "Running in $DIR"
mkdir -p $DIR/.log
echo "Storing logs in $DIR/.log/log.txt"

starport chain serve -f --path  $DIR -v $@ 2>$DIR/.log/log.txt &
tail -f $DIR/.log/log.txt | sed -r "s/\x1B\[(([0-9]+)(;[0-9]+)*)?[m,K,H,f,J]//g" > $DIR/.log/log_proper.txt &
tail -f $DIR/.log/log_proper.txt
