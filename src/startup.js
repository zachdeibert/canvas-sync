const app = require("./app");
const options = require("./options");

const authenticatedRegex = /^https:\/\/[-a-zA-Z0-9]+.instructure.com\/$/;

app.onLoad(() => {
    app.window.webContents.once("will-navigate", (ev, url) => {
        ev.preventDefault();
        let query = {};
        url.split("?", 2)[1].split("&").forEach(s => {
            let p = s.split("=");
            query[p[0]] = decodeURIComponent(p[1]);
        });
        app.window.webContents.loadURL(query["url"]);
        let listener = () => {
            const url = app.window.webContents.getURL();
            if (authenticatedRegex.test(url)) {
                app.window.webContents.removeListener("did-redirect-navigation", listener);
                app.window.webContents.removeListener("did-navigate", listener);
                options.startCollection(query["out"]);
            }
        };
        app.window.webContents.on("did-redirect-navigation", listener);
        app.window.webContents.on("did-navigate", listener);
    });
    app.window.loadURL("static://startup/startup.html");
});
