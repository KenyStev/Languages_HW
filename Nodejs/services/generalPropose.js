var fs = require('fs');

exports.upload = function(req, savepath, cb){
    var file;
 
    if (!req.files) {
        cb(new Error("No files were uploaded"))
    }
 
    file = req.files.file;
    if(!file){
        cb(new Error("Cannot find file in files param"))
    }else{
        let folder = savepath + file.name.split(".")[0];
        createFolder(folder);
        var uploadPath = folder + "/" + file.name;
        file.mv(uploadPath, function(err) {
            if (err) {
                cb(new Error('Cannot copy file'));
            }
            else {
                cb(null);
            }
        });
    }
}

var createFolder = function(folder) {
    fs.mkdirSync(folder);
}