# Url shortener

A simple Url shortener using Golang and Redis as LRU Cache.

## Getting Started

To install the app, follow these steps:

1. **Clone the repository:**

    `git clone https://github.com/yourusername/url-shortener.git`

2. **Navigate to project directory:**

    `cd url-shortener`

3. **Build the image and run the container:**

    `docker-compose up`

4. To get a short url send a POST request like this:
```curl POST localhost:8000/create-short-url -d '{"long_url": "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html", "user_id" : "e0dba740-fc4b-4977-872c-d360239e6b10"}'```

5. Now open up ypur browser and hit the url.


## Contributing

Any contributions from the community is welcome! If you'd like to contribute to Realtime Chat App, please follow these guidelines:

1. **Fork the repository and create your branch:**

    `git checkout -b feature/new-feature`

2. **Make your changes and commit them:**

    `git commit -m "Add new feature"`

3. **Push to your branch:**
    `git push origin feature/new-feature`

4. **Create a pull request with a detailed description of your changes.**

## License

This project is licensed under the [MIT License](LICENSE).

