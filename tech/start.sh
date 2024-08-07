docker-compose pull --remove-orphans
docker-compose up --force-recreate --remove-orphans --build -d
sleep 1
docker system prune -f -a
docker volume prune -f
