# b365-twitter-api-search test

1. Some application properties are defined in an .env file, and I have intentionally not added the .env file in my gitignore file to show a sample of how properties are defined/retrieved.

2. Values like SERVER_PORT, TWITTER_SEARCH_API, ACCESS_TOKEN and so on are defined in the .env file. However, to test the API call to twitter search endpoint, we'll require a personalized ACCESS_TOKEN from the developer dashboard (https://developer.twitter.com/en/portal/dashboard). The TWITTER-API-ACCESS-TOKEN value should be replaced (including the <>) with the generated ACCESS_TOKEN.

3. This simple rest api service exposes two endpoints; a ping and the twitter search endpoints. The Postman collection can be automatically imported using this link (https://www.getpostman.com/collections/28730349b29d70f0c278)