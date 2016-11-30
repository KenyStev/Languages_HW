var fs = require('fs'),
	readline = require('readline'),
	LineReaderSync = require("line-reader-sync");

const ENV = "dev";
const rootpath = "resources/mergesort/";

var print = (function () {
    return {
        log: function() {
            var args = Array.prototype.slice.call(arguments);
            if (ENV === 'dev') {
            	console.log.apply(console, args);
            }
        }
    }
}());

var GetLeaves = function(dirpath) {
	return fs.readdirSync(rootpath + dirpath);
}

var MergeSort = function(dirpath,cb) {
	var leaves = GetLeaves(dirpath);
	let sortedpath = dirpath + "sorted/";
	createFolder(sortedpath);
	for (var i in leaves) {
		// print.log("sort MergeSort: ",dirpath + leaves[i]);
		SortFile(dirpath + leaves[i]);
	}
	leaves = GetLeaves(sortedpath);
	// print.log("leaves: ",leaves, "length: ",leaves.length);
	mergesort(sortedpath,leaves,0);
	cb(null);
}

var mergesort = function(path,m,cont){
	if (m.length <= 1) {
		return m;
	}

	print.log("m.length: ",m.length);

	let mid = m.length / 2;
	let left = m.slice(0,mid);
	let right = m.slice(mid,m.length);

	print.log("mid: " + mid);
	print.log("left arr: " + left);
	print.log("right arr: " + right);

	left = mergesort(path,left,cont+1);
	right = mergesort(path,right,cont+2);

	return merge(path,left, right,cont);
}

var merge = function(path,left, right, cont) {
	var result = [];
	let name = cont + "_" + (left[0]).split(".")[0] + "_" + (right[0]).split(".")[0] + ".merged.sorted";
	if (cont == 0) {name = cont + ".merged.sorted";};
	print.log("sorted name: ",name," left: ",left," right: ",right);
	result.push(name);
	let left_file = new LineReaderSync(rootpath + path + left[0]);
	let right_file = new LineReaderSync(rootpath + path + right[0]);
	let new_file_path = rootpath + path + name;

	let left_word = left_file.readline();
	let right_word = right_file.readline();
	while(true) {
		if (left_word == null) {
			left_word = left_file.readline();
		}
		if (right_word == null) {
			right_word = right_file.readline();
		}

		if (left_word != null && right_word != null) {
			print.log("left: "+left_word + " " + left_word.length);
			print.log(" right: "+right_word+ " " + right_word.length);
			if (left_word <= right_word) {
				fs.appendFileSync(new_file_path,left_word+"\n");
				left_word = null;
			} else {
				fs.appendFileSync(new_file_path,right_word+"\n");
				right_word = null;
			}
		} else if (left_word != null) {
			print.log("lesf: "+left_word + " " + left_word.length);
			fs.appendFileSync(new_file_path,left_word+"\n");
			left_word = null;
		} else if (right_word != null) {
			print.log(" right: "+right_word+ " " + right_word.length);
			fs.appendFileSync(new_file_path,right_word+"\n");
			right_word = null;
		}else{
			break;
		}
	}

	return result
}

var createFolder = function(filename) {
	let folder = filename.split(".")[0];
	fs.mkdirSync(rootpath+folder);
}

var FilterFile = function(filepath, pattern,cb) {
	let regx = new RegExp(pattern);
	let file = fs.openSync(rootpath + filepath,'r');
	
	let createdFile = rootpath +filepath + ".filtered";
	
	// print.log("Readfile: "+rootpath + filepath);
	// print.log("filtered: "+createdFile);
	let rd = readline.createInterface({
	    input: fs.createReadStream(rootpath + filepath),
	    output: process.stdout,
	    terminal: false
	});

	rd.on('line',(line) => {
		if (match = regx.exec(line)) {
			// print.log(line, filteredFile);
			fs.appendFileSync(createdFile,line.toLowerCase()+"\n");
		};
	});

	fs.closeSync(file);

	rd.on('close',function(){
		cb(null);
	});
}

var CreateLeaves = function(filepath, leafSize, cb) {
	let folder = rootpath + filepath.split("/")[0] + "/leaves/";
	createFolder(folder.replace(rootpath,""));
	let leafcont = 0;
	var filteredFile;
	// print.log("path: " + rootpath + filepath);
	var rd = readline.createInterface({
	    input: fs.createReadStream(rootpath + filepath),
	    output: process.stdout,
	    terminal: false
	});

	var linecont = 0;
	rd.on('line',(line) => {
		if (linecont%leafSize == 0) {
			if(linecont>0)
				fs.closeSync(filteredFile);
			// print.log(linecont);
			filteredFile = fs.openSync(folder +"leaf" + leafcont,"a+");
			leafcont++;
		}
		fs.writeSync(filteredFile,line+"\n");
		linecont++;
	});

	rd.on('close', function(){
		cb(null);
	});
}

var SortFile = function(filepath) {
	let fileData = fs.readFileSync(rootpath + filepath,'utf8');
	let data = fileData.split("\n");
	data.sort((a,b) => {
		if (a < b)
			return -1;
		else if (a === b) 
			return 0;
		else
			return 1;
	});
	let fullpath = filepath.split("/");
	let filename_temp = fullpath.pop();
	fullpath.push("sorted");
	// createFolder(fullpath.join('/'));
	fullpath.push(filename_temp);
	// print.log("fullpath: "+fullpath);
	// print.log("swap fullpath: "+fullpath);
	sortedpath = rootpath + fullpath.join('/');
	// print.log("path: "+ sortedpath);
	let sortedFile = fs.openSync(sortedpath + ".sorted", "a+");
	data = data.slice(1,data.length) //delete '\n'
	for (var line in data) {
		fs.writeSync(sortedFile,data[line]+"\n");
	}
	fs.closeSync(sortedFile);
}

var GetSortedFile = function(name){
	return rootpath + name.split(".")[0] +"/leaves/sorted/0.merged.sorted";
}

// var files = 
// SortFile("emails/emails.txt");
// print.log(files);

var test = function(){
	let filename = "emails.txt";
	let name = filename.split(".")[0];
	FilterFile(name+"/"+filename,"(\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3})",function(err){
		if (err) {
			print.log("error filter");
		}else{
			print.log("termino filter");
			CreateLeaves(name+"/"+filename+".filtered",5, function(err){
				if(err)
					print.log("error create leaves");
				else{
						print.log("termino filter");
						MergeSort(name+"/leaves/",function(err){
							if (err) {
								print.log("error MergeSort");
							}
						});
					}
			});
		}
	});
}

exports.FilterFile = FilterFile;
exports.CreateLeaves = CreateLeaves;
exports.MergeSort = MergeSort;
exports.GetSortedFile = GetSortedFile;

// test();

// FilterFile("emails/emails.txt",/^(([^<>()\[\]\.,;:\s@\"]+(\.[^<>()\[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i);