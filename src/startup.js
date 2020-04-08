const app = require("./app");
const db = require("./db");
const proto = require("./proto");
const options = require("./options");

const authenticatedRegex = /^https:\/\/[-a-zA-Z0-9]+.instructure.com\/$/;

app.onLoad(() => {
    proto.once("startup", query => {
        app.window.webContents.loadURL(query["url"]);
        let listener = () => {
            const url = app.window.webContents.getURL();
            if (authenticatedRegex.test(url)) {
                app.window.webContents.removeListener("did-redirect-navigation", listener);
                app.window.webContents.removeListener("did-navigate", listener);
                options.startCollection(new db.DataBase(query["out"]));
            }
        };
        app.window.webContents.on("did-redirect-navigation", listener);
        app.window.webContents.on("did-navigate", listener);
    });
    app.window.loadURL("static://startup/startup.html");
});
