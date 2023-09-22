module.exports = {
  ci: {
    collect: {
      url: ['http://localhost:8080/'],
      startServerCommand: 'go run ./go/main.go',
    },
    settings: {
      extraHeaders: {
        'ngrok-skip-browser-warning': 'true',
      }
    },
    upload: {
      target: 'temporary-public-storage',
    },
    githubToken: process.env.GITHUB_TOKEN,
  },
};
