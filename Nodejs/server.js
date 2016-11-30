/**
 * Module dependencies.
 */

var express = require('express')
  , http = require('http')
  , fileUpload = require('express-fileupload')
  , gp = require('./services/generalPropose.js')
  , services = require('./services/services.js')
  , fs = require('fs');

const SORTPATH = "resources/mergesort/";
const BITCODEPATH = "resources/bitcode/";

var app = express();
app.use(fileUpload());

app.get('/',function(req, res){
    res.json(200, "Hello World" );
  }
);

app.post('/api/sort',function(req,res){
  gp.upload(req,SORTPATH,function(err){
    if (err) {
      res.status(500).json(err)
    }else{
      services.SortEmails(req.files.file.name,function(downloadPath){
        res.download(downloadPath);
      });
    }
  });
});

app.post('/upload', function(req, res){
    console.log(req.files);
    gp.upload(req,"./uploads/",function(err){
      if(err)
        res.status(500).json(err);
      else
        res.status(200).json("Uploaded");
    }
  );
});


http.createServer(app).listen(8000, function(){
  console.log("Express server listening on port %s in %s mode.",  8000, app.settings.env);
});