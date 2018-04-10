a="one\\rtwo\\rthree\\rfour"
echo $a
b=${a//\\r/,}
OLD_IFS="$IFS"
IFS=","
arr=($b)
IFS="$OLD_IFS"
for s in ${arr[@]}
do
    echo "$s"
done

length=${#arr[@]}
echo $length
echo ${arr[$length-1]}