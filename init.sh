echo "######## Building Image ########"
docker build . -t greedy_games_assignment
echo "######## Starting Image ########"
docker run --name greedy_games_assignment -p 9000:9000 greedy_games_assignment
