# GeniusLurker

## Install and run

```
git clone http://github.com/AlexanderYAPPO/geniuslurker/
cd geniuslurker/
sudo docker build -t geniuslurker .
docker run -t -i -e GENIUS_TELEGRAM_TOKEN=<token> -e GENIUS_API_TOKEN="Bearer <token>" -p 80:80 geniuslurker
```
