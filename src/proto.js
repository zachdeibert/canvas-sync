const electron = require("electron");
const path = require("path");
const app = require("./app");

electron.protocol.registerSchemesAsPrivileged([{
    "scheme": "static",
    "privileges": {
        "standard": true,
        "secure": true
    }
}]);

app.onLoad(() => {
    electron.protocol.registerFileProtocol("static", (req, callback) => {
        let url = req.url.split("/").slice(3).join("/");
        if (url.startsWith("node_modules")) {
            callback({
                path: path.normalize(`${__dirname}/../${url}`)
            });
        } else {
            callback({
                path: path.normalize(`${__dirname}/../html/${url}`)
            });
        }
    });
});
