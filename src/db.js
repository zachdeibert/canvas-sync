const fs = require("fs");
const path = require("path");

const invalidChars = /[^-a-zA-Z0-9 _]/gi;

function DataBase(dir) {
    console.log(dir);
    this.path = dir;
    if (!fs.existsSync(dir)) {
        fs.mkdirSync(dir, {
            "recursive": true
        });
    }
}

DataBase.prototype.forCourse = function(courseId, courseName, courseCode, courseTerm) {
    let d = this.path;
    if (courseTerm) {
        d = path.join(d, courseTerm.replace(invalidChars, "_"));
    }
    d = path.join(d, `${courseCode} - ${courseName} - ${courseId}`.replace(invalidChars, "_"))
    return new DataBase(d);
};

DataBase.prototype.forScraper = function(name) {
    return new DataBase(path.join(this.path, name));
};

module.exports = {
    "DataBase": DataBase
};
