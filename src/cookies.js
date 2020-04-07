const electron = require("electron");
const app = require("./app");

app.onLoad(() => {
    let cookies = electron.session.defaultSession.cookies;
    cookies.on("changed", (ev, cookie, cause, removed) => {
        if (cookie.session && !removed) {
            let url = `${cookie.secure ? "https" : "http"}://${cookie.domain}${cookie.path}`;
            cookies.set({
                "url": url,
                "name": cookie.name,
                "value": cookie.value,
                "domain": cookie.domain,
                "path": cookie.path,
                "secure": cookie.secure,
                "httpOnly": cookie.httpOnly,
                "expirationDate": Math.floor(new Date().getTime()/1000) + 12096000
            }).catch(() => {});
        }
    });
});
