var   bitFuncs = require('./bitFuncs.js')
	, fs = require('fs-extra');

const rootpath = "resources/bitcode/";

var HideMessage = function(imagename,message) {
	let folder = imagename.split(".")[0] + "/";
	saveMessage(rootpath+folder+"message.txt",message);

	try{
		fs.copySync(rootpath+folder+imagename,rootpath+folder+"hidden_"+imagename);
	}catch(err){
		console.error(err)
	}
	let imageOut = fs.openSync(rootpath+folder+"hidden_"+imagename,'r+');

	let start = getImageOffsetStart(imagename);

	console.log("-----> len <-----");
	console.log("msg len: %d",message.length);
	let messageLen = new Buffer(4);
	messageLen.writeUInt32LE(message.length,0);
	console.log("msg len size: %d",messageLen.length);
    // binary.LittleEndian.PutUint32(messageLen, uint32(len(message)))
    console.log(messageLen);
	console.log("start before: ",start);
	writeBits(imageOut,messageLen,start);
	console.log("start after: ",start);
	console.log("-----> msg <-----");
	let messagebits = new Buffer(message,'utf8');
	writeBits(imageOut,messagebits,start + messageLen.length*8);

	fs.closeSync(imageOut);
}

var SeekMessage = function(imagename){
	let folder = imagename.split(".")[0] + "/";
	let image = fs.openSync(rootpath+folder+imagename,'r');
	// defer image.Close()

	let start = getImageOffsetStart(imagename);
	// image.Seek(int64(start),0);
	console.log("-----> len <-----");
	let messageLenBytes = readBits(image,4,start);
	console.log(messageLenBytes);
	let messageLen = messageLenBytes.readUInt32LE(0);
	console.log("msg len: %d",messageLen)
	console.log("-----> msg <-----");
	let message = readBits(image,messageLen,start + 4*8).toString('utf8',0,messageLen);
	console.log("msg: ",message);
	
	saveMessage(rootpath+folder+"message.txt",message);
}

var writeBits = function(fileout, bytes, offset) {
	bytes.forEach(function(byte_i) {
		for (var i = 0; i < 8; i++) {
			console.log("%t",bitFuncs.hasBit(byte_i,i));
			let imgByte = new Buffer(1);
			fs.readSync(fileout,imgByte,0,1,offset);

			console.log("imgByte before: ",imgByte);
			if (bitFuncs.hasBit(byte_i,i)) {
				imgByte[0] = bitFuncs.setBit(imgByte[0],0);
			}else{
				imgByte[0] = bitFuncs.clearBit(imgByte[0],0);
			}
			console.log("imgByte after: ",imgByte);
			fs.writeSync(fileout,imgByte,0,1,offset);
			offset++;
		}
	});
}

var readBits = function(filein, byteslen, offset) {
	let bytes = new Buffer(byteslen);
	for (bit = 0; bit<byteslen; bit++) {
		let new_byte = 0;
		for (i = 0; i < 8; i++) {
			let fileByte = new Buffer(1);
			fs.readSync(filein,fileByte,0,1,offset++);
			if (bitFuncs.hasBit(fileByte[0],0)) {
				new_byte = bitFuncs.setBit(new_byte,i);
			}else{
				new_byte = bitFuncs.clearBit(new_byte,i);
			}
		}
		bytes[bit] = new_byte;
	}
	return bytes;
}

var getImageOffsetStart = function(imagename) {
	let folder = imagename.split(".")[0] + "/";
	let image = fs.openSync(rootpath + folder+imagename,'r');

	let bytesToRead = 2;
	var buffer = new Buffer(bytesToRead);
    fs.readSync(image, buffer, 0, bytesToRead, 28);
	
	bitsPerColor = buffer.readUInt8(0, 2);
	console.log("bitsPerColor: ",buffer.readUInt8(0, 2));

	let numColors = 0;
	if (bitsPerColor <= 8) {
		numColors = Math.pow(2,bitsPerColor);
		console.log("numColors: ",numColors);
	}

	let headerSize = 54;
	let colorTableSize = 4*numColors;
	let imageOffset = headerSize + colorTableSize;

	console.log("imageOffset dec: ",imageOffset,", hex: ",imageOffset.toString(16));
	fs.closeSync(image);
	return imageOffset;
}

var saveMessage = function(filename,message) {
	fs.appendFileSync(filename,message);
}

var GetHidden = function(name){
	return rootpath + name.split(".")[0] + "/hidden_" + name;
}

var GetMessage = function(name){
	return rootpath + name.split(".")[0] + "/message.txt";
}

exports.HideMessage = HideMessage;
exports.SeekMessage = SeekMessage;
exports.GetHidden = GetHidden;
exports.GetMessage = GetMessage;

/*var test = function(){
	// let file = fs.openSync("resources/bitcode/oceano/hide_oceano.bmp",'r+');
	// let buf = new Buffer(4);
	// buf.writeUInt32LE(56,0);
	// writeBits(file,buf,54);
	// let num = readBits(file,4,54).readUInt32LE(0);
	// console.log(readBits(file,4,54));
	// console.log(num);
	// fs.createReadStream('resources/bitcode/oceano/oceano.bmp').pipe(fs.createWriteStream('resources/bitcode/oceano/hide_oceano.bmp'));
	// HideMessage("oceano.bmp","Hola Mundo, que tal estas?");
	// SeekMessage("hidden_oceano.bmp",54);
}

test();*/