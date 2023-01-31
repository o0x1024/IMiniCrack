const fs = require("fs-extra");

module.exports = {
    read: function (_path) {
        let filesData = fs.readFileSync(_path, "utf-8", function (e, data) {
            if (e) throw e;
            return data;
        });
        return filesData;
    },
    write: function (writeStr) {
        fs.open(_path, "w", function (e, fd) {
            if (e) throw e;
            fs.write(fd, writeStr, 0, "utf8", function (e) {
                if (e) throw e;
                fs.closeSync(fd);
            });
        });
    }
}; 