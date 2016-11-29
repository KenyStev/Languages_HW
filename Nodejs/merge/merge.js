var fs = require('fs'),
	readline = require('readline'),
	LineReaderSync = require("line-reader-sync");

const rootpath = "resources/mergesort/"

var GetLeaves = function(dirpath) {
	// var leaves = [];
	// var leaves_ = fs.readdirSync(rootpath + dirpath);
	// for (var leaf in leaves_) {
	// 	leaves.push(leaves_[leaf]);
	// };
	// return leaves;
	return fs.readdirSync(rootpath + dirpath);
}

var MergeSort = function(dirpath) {
	var leaves = GetLeaves(dirpath);
	let sortedpath = dirpath + "sorted/";
	createFolder(sortedpath);
	for (var i in leaves) {
		// console.log("sort MergeSort: ",dirpath + leaves[i]);
		SortFile(dirpath + leaves[i]);
	}
	leaves = GetLeaves(sortedpath);
	// console.log("leaves: ",leaves, "length: ",leaves.length);
	mergesort(sortedpath,leaves,0);
}

var mergesort = function(path,m,cont){
	if (m.length <= 1) {
		return m;
	}

	console.log("m.length: ",m.length);

	let mid = m.length / 2;
	let left = m.slice(0,mid);
	let right = m.slice(mid,m.length);

	console.log("mid: " + mid);
	console.log("left arr: " + left);
	console.log("right arr: " + right);

	left = mergesort(path,left,cont+1);
	right = mergesort(path,right,cont+2);

	return merge(path,left, right,cont);
}

var merge = function(path,left, right, cont) {
	var result = [];
	let name = cont + ".merged.sorted";
	console.log("sorted name: ",name);
	result.push(name);
	let left_file = new LineReaderSync(rootpath + path + left[0]);
	// defer left_file.Close()
	let right_file = new LineReaderSync(rootpath + path + right[0]);
	// defer right_file.Close()
	// new_file,_ := os.Create(rootpath + path + name);
	// defer new_file.Close()
	let new_file_path = rootpath + path + name;

	// scan_left := bufio.NewScanner(left_file)
	// scan_right := bufio.NewScanner(right_file)
	// pos_left := 0; pos_right := 0

	// var left_word,right_word string

	let left_word = left_file.readline();
	let right_word = right_file.readline();
	while(true) {
		// scan_left.Scan(); scan_right.Scan()
		if (!left_word) {
			left_word = left_file.readline();
		}
		if (!right_word) {
			right_word = right_file.readline();
		}

		if (left_word != null && right_word != null) {
			console.log("lesf: "+left_word + " " + left_word.length);
			console.log(" right: "+right_word+ " " + right_word.length);
			if (left_word <= right_word) {
				// new_file.WriteString(left_word+"\n")
				fs.appendFileSync(new_file_path,left_word+"\n");
				left_word = null;
				// pos_left += len(left_word) + 1
				// right_file.Seek(int64(pos_right),0)
				// scan_right = bufio.NewScanner(right_file)
			} else {
				// new_file.WriteString(right_word+"\n")
				fs.appendFileSync(new_file_path,right_word+"\n");
				right_word = null;
				// pos_right += len(right_word)+ 1
				// left_file.Seek(int64(pos_left),0)
				// scan_left = bufio.NewScanner(left_file)
			}
		} else if (left_word != null) {
			console.log("lesf: "+left_word + " " + left_word.length);
			fs.appendFileSync(new_file_path,left_word+"\n");
			left_word = null;
		} else if (right_word != null) {
			console.log(" right: "+right_word+ " " + right_word.length);
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
	
	// scanner := bufio.NewScanner(file)
	let createdFile = rootpath +filepath + ".filtered";
	//let filteredFile = fs.openSync(createdFile,"a+");
	
	console.log("Readfile: "+rootpath + filepath);
	console.log("filtered: "+createdFile);
	let rd = readline.createInterface({
	    input: fs.createReadStream(rootpath + filepath),
	    output: process.stdout,
	    terminal: false
	});

	rd.on('line',(line) => {
		if (match = regx.exec(line)) {
			// console.log(line, filteredFile);
			fs.appendFileSync(createdFile,line.toLowerCase()+"\n");
		};
	});

	// fs.closeSync(filteredFile);
	fs.closeSync(file);

	rd.on('close',function(){
		cb(null);
	});
}

var CreateLeaves = function(filepath, leafSize, cb) {
	let folder = rootpath + filepath.split("/")[0] + "/leaves/";
	createFolder(folder.replace(rootpath,""));
	// let file = fs.openSync(rootpath + filepath,"r");
	let leafcont = 0;
	var filteredFile;
	console.log("path: " + rootpath + filepath);
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
			console.log(linecont);
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
	data.sort((a,b) => {return a - b;});
	let fullpath = filepath.split("/");
	let filename_temp = fullpath.pop();
	fullpath.push("sorted");
	// createFolder(fullpath.join('/'));
	fullpath.push(filename_temp);
	console.log("fullpath: "+fullpath);
	console.log("swap fullpath: "+fullpath);
	sortedpath = rootpath + fullpath.join('/');
	console.log("path: "+ sortedpath);
	let sortedFile = fs.openSync(sortedpath + ".sorted", "a+");
	data = data.slice(0,data.length-1) //delete '\n'
	for (var line in data) {
		fs.writeSync(sortedFile,data[line]+"\n");
	}
	fs.closeSync(sortedFile);
}

// var files = 
// SortFile("emails/emails.txt");
// console.log(files);

var test = function(){
	let filename = "emails.txt";
	let name = filename.split(".")[0];
	FilterFile(name+"/"+filename,"(\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3})",function(err){
		if (err) {
			console.log("error filter");
		}else{
			console.log("termino filter");
			CreateLeaves(name+"/"+filename+".filtered",5, function(err){
				if(err)
					console.log("error create leaves");
				else{
						console.log("termino filter");
						MergeSort(name+"/leaves/");
					}
			});
		}
	});
	
	
	// LineReaderSync = require("line-reader-sync")
	// lrs = new LineReaderSync(rootpath + "emails/emails.txt");
	// while(true){
	//   var line = lrs.readline()
	//   if(line === null){
	//     console.log("EOF");
	//     break;
	//   }else{
	//     console.log("line without \n",line)
	//   }
	  
	// }
}

test();

// FilterFile("emails/emails.txt",/^(([^<>()\[\]\.,;:\s@\"]+(\.[^<>()\[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i);