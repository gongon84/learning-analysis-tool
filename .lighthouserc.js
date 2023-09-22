module.exports = {
  ci: {
    collect: {
      url: ['http://localhost:8080/'],
      startServerCommand: 'go run ./go/main.go',
    },
    assert: {},
    upload: {
      target: 'temporary-public-storage',
    },
  },
};
