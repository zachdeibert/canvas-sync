const electron = require("electron");
const fs = require("fs");

let loadPromise = electron.app.whenReady();

module.exports = {
    "window": null,
    "onLoad": callback => {
        if (typeof(callback) === "function") {
            loadPromise = loadPromise.then(() => callback());
        }
    },
    "runScript": name => new Promise((resolve, reject) => {
        fs.readFile(`${__dirname}/scrapers/injections/${name}.js`, "utf8", (err, data) => {
            if (err) {
                reject(err);
            } else {
                module.exports.window.webContents.executeJavaScript(data).then(resolve).catch(reject)
            }
        });
    })
};

module.exports.onLoad(() => {
    module.exports.window = new electron.BrowserWindow({
        "show": true,
        "webPreferences": {
            "enableRemoteModule": false,
            "webSecurity": false
        }
    });    
    module.exports.window.webContents.userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:74.0) Gecko/20100101 Firefox/74.0";
});
