Команди:

1. Отримання курсу BTC до UAH(GET): http://localhost:8080/api/rate
2. Підписати email на розсилку(POST): http://localhost:8080/api/subscribe
3. Здійснити розсилку з поточним курсом(GET): http://localhost:8089/api/sendEmails

To run Docker:
1. docker build -t btc-app .  
2. docker run --name=btc-app -p 8080:8080 btc-app    

Архітектура:
https://app.diagrams.net/#Lbtc-application-layers