const app = require("./app");
const proto = require("./proto");

module.exports = {
    "startCollection": db => {
        proto.once("homepage", res => {
            console.log(res);
        });
        app.runScript("homepage");
    }
};
