const fs = require("fs");

module.exports = callback => {
    fs.access("config.js", err => {
        if (err) {
            require("./config-gen")(callback);
        } else {
            callback(require("./config"));
        }
    });
};
