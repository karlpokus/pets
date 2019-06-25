const mock = {
  isConnected: () => true,
  db: function() { return this },
  collection: function() { return this },
  findOne: ({name, pwd}) => {
    return new Promise((resolve, reject) => {
      if (name == "buck" && pwd == "nasty") {
        return resolve(true)
      }
      return reject(new Error("Unauthorized"))
    });
  }
};

module.exports = mock;
