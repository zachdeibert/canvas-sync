const fs = require("fs");
const app = require("../app");
const canvas = require("../canvas");
const proto = require("../proto");
const scraper = require("./scraper");

scraper.register("Announcements", (courseId, db, update, del) => new Promise((resolve, reject) => {
    const txtExtension = /\.txt$/;
    const callback = res => {
        console.log(res);
        db.run(txtExtension, res.map(a => `${a.id} - ${a.title}.txt`), (filename, idx) => {
            return fs.promises.readFile(filename, {
                "encoding": "utf8"
            }).then(str => {
                var parts = str.split("\x03");
                return parts[0].trim() === res[idx].text.trim() && parts.length - 1 === res[idx].replies;
            });
        }, update, del, (filename, idx) => new Promise((resolve, reject) => {
            console.log(res2);
            const callback = res2 => {
                let plaintext = res[idx].text;
                if (res2.replies.length > 0) {
                    plaintext += `\n\x03\nReplies:\n${res2.replies.map(thread => 
                        thread.map(reply => 
                            `From ${reply.author} at ${new Date(reply.time)}:\n${reply.text}`
                        ).join("\n\x03\n")
                    ).join("\n\x03----------------------------------------\n")}\n`;
                }
                let html = "<!DOCTYPE html>" +
                           "<html>" +
                               "<head>" +
                                   "<meta charset=\"utf-8\" />" +
                                   `<title>${res[idx].title}</title>` +
                               "</head>" +
                               "<body>" +
                                   `<h1>${res[idx].title}</h1>` +
                                   `<h4>Published by ${res2.author} at ${new Date(res[idx].postDate)}</h4>` +
                                   "<main>" +
                                       res2.html
                                   "</main>" +
                                   res2.replies.map(thread =>
                                       "<hr />" +
                                       thread.map(reply =>
                                           `<h3>Reply from ${reply.author} at ${new Date(reply.time)}:</h3>` +
                                           "<div>" +
                                               reply.html +
                                           "</div>"
                                       ).join("")
                                   ).join("")
                               "</body>" +
                           "</html>";
                Promise.all(
                    fs.promises.writeFile(filename, plaintext),
                    fs.promises.writeFile(filename.replace(txtExtension, ".html"), html)
                ).then(resolve).catch(reject);
            };
            proto.once("announcement", callback).then(() => {
                return canvas.load(courseId, "discussion_topics", res[idx].id).then(() => app.runScript("announcement")).catch(e => {
                    proto.remove("announcement", callback);
                    reject(e);
                });
            }).catch(reject);
        }), (filename) => {
            return Promise.all(
                fs.promises.unlink(filename),
                fs.promises.unlink(filename.replace(txtExtension, ".html"))
            );
        }).then(resolve).catch(reject);
    };
    proto.once("announcements", callback).then(() => {
        return canvas.load(courseId, "announcements").then(() => app.runScript("announcements")).catch(e => {
            proto.remove("announcements", callback);
            reject(e);
        });
    }).catch(reject);
}));
