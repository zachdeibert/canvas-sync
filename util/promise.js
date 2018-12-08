module.exports = {
    "syncAll": (...promises) => {
        if (promises.length === 1 && Array.isArray(promises[0])) {
            promises = promises[0];
        }
        let promise = Promise.resolve();
        promises.forEach(p => {
            promise = promise.then(p);
        });
        return promise;
    }
};
