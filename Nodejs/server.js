/**
 * Module dependencies.
 */

var express = require('express')
  , http = require('http')
  , fileUpload = require('express-fileupload')
  , gp = require('./services/generalPropose.js');

var app = express();
app.use(fileUpload());

// app.configure(function(){
  // app.set('port', process.env.PORT || 8000);
  // app.use(express.favicon());
  // app.use(express.methodOverride());
  // app.use(app.router);
  // app.use(express.static(__dirname + '/public'));
// });

// app.configure('development', function(){
  // app.use(express.errorHandler());
// });

app.get('/',function(req, res){
    res.json(200, "Hello World" );
  }
);

app.get('/upload', function(req,res){
  res.render("./templates/upload.html");
});

app.post('/upload', function(req, res){
    console.log(req.files);
    gp.upload(req,"./uploads/file",function(err){
      if(err)
        res.status(500).json(err)
      else
        res.status(200).json("Uploaded");
    }
  );
});


http.createServer(app).listen(8000, function(){
  console.log("Express server listening on port %s in %s mode.",  8000, app.settings.env);
});