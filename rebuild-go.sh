docker stop golang-app
docker rm golang-app
docker rmi golang-app-image:local
# docker build -t golang-app-image:local --progress=plain .
# docker run -d -p 8080:8080 --name golang-app golang-app-image:local