/**
 * Module dependencies.
 */

var express = require('express')
  , http = require('http')
  , bodyParser = require('body-parser')
  , fileUpload = require('express-fileupload')
  , gp = require('./services/generalPropose.js')
  , services = require('./services/services.js')
  , fs = require('fs');

const SORTPATH = "resources/mergesort/";
const BITCODEPATH = "resources/bitcode/";

var app = express();
app.use(fileUpload());
app.use(bodyParser.json()); // support json encoded bodies
app.use(bodyParser.urlencoded({ extended: true })); // support encoded bodies

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

app.post('/api/bitcode',function(req,res){
  gp.upload(req,BITCODEPATH,function(err){
    if (err) {
      res.status(500).json(err)
    }else{
      services.HideMessage(req.files.file.name,req.body.message,function(downloadPath){
        res.download(downloadPath);
      });
    }
  });
});

app.post('/api/bitcode/seek',function(req,res){
  gp.upload(req,BITCODEPATH,function(err){
    if (err) {
      res.status(500).json(err)
    }else{
      services.SeekMessage(req.files.file.name,function(downloadPath){
        res.download(downloadPath);
      });
    }
  });
});

app.post('/api/kruskal',function(req,res){
  // gp.upload(req,SORTPATH,function(err){
    // if (err) {
      // res.status(500).json(err)
    // }else{
      console.log("-- kruskal --")
      console.log(req.body)
      // let graph = req.body //JSON.parse(req.body)
      services.Kruskal(req.body,function(new_graph){
        res.status(200).json(new_graph);
      });

      // res.status(200).json("Hola bad")
    // }
  // });
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