var merge = require("../merge/merge.js");

exports.SortEmails = function(filename,cb) {
	let name = filename.split(".")[0];
	merge.FilterFile(name+"/"+filename,"(\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3})",function(err){
		if (err) {
			console.log("error filter");
		}else{
			console.log("termino filter");
			merge.CreateLeaves(name+"/"+filename+".filtered",5, function(err){
				if(err)
					console.log("error create leaves");
				else{
						console.log("termino filter");
						merge.MergeSort(name+"/leaves/",function(err){
							if (err) {
								console.log("error merge");
							}else{
								console.log("done: ",filename);
								cb(merge.GetSortedFile(filename));
							}
						});
					}
			});
		}
	});
};

