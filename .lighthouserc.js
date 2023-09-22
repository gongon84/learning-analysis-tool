module.exports = {
  ci: {
    collect: {
      url: ['http://localhost:8080/'],
      startServerCommand: 'go run ./go/main.go',
    },
    assert: {
      preset: 'lighthouse:recommended',
    },
    upload: {
      target: 'temporary-public-storage',
    },
  },
};
