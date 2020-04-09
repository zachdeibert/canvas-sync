const fs = require("fs");
const path = require("path");

const invalidChars = /[^-a-zA-Z0-9 _]/gi;

function DataBase(dir) {
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

DataBase.prototype.run = function(fileRegex, fileList, fileComparator, update, del, updateFunc, deleteFunc) {
    return fs.promises.readdir(this.path).then(files => {
        let resolve;
        let promise = new Promise((_resolve, reject) => {
            resolve = _resolve;
        });
        let expectedFiles = fileList.slice();
        let fileIdxs = fileList.map((_, i) => i);
        const queue = (cond, func, filename, idx) => {
            if (cond) {
                promise = promise.then(() => func(path.join(this.path, filename), idx));
            }
        };
        files.forEach(file => {
            if (fileRegex.test(file)) {
                let idx = expectedFiles.indexOf(file);
                if (idx >= 0) {
                    if (fileComparator(file, fileIdxs[idx])) {
                        expectedFiles.splice(idx, 1);
                        fileIdxs.splice(idx, 1);
                    } else {
                        queue(update, updateFunc, file, fileIdxs[idx]);
                    }
                } else {
                    queue(del, deleteFunc, file, -1);
                }
            }
        });
        expectedFiles.forEach((file, i) => {
            queue(update, updateFunc, file, fileIdxs[i]);
        });
        resolve();
        return promise;
    });
};

module.exports = {
    "DataBase": DataBase
};
