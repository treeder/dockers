echo "running ./hello.sh"
for i in {0..10}
do
    time ./hello.sh
done

echo "running ./hello.sh in docker"
for i in {0..10}
do
    time docker run treeder/hello
done

echo "running ./hello.sh in docker with --rm"
for i in {0..10}
do
    time docker run --rm treeder/hello
done

echo "running ./hello.sh in docker just reusing container"
echo "initial run"
docker run --name reuse treeder/hello
for i in {0..10}
do
    time docker start reuse
done
