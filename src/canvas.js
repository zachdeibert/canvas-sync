const app = require("./app");

module.exports = {
    "baseURL": null,
    "load": (courseId, ...path) => {
        return app.window.webContents.loadURL(`${module.exports.baseURL}courses/${courseId}/${path.join("/")}`);
    }
};
