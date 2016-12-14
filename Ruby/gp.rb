def upload(params,savepath)
	unless params[:file] &&
			(tmpfile = params[:file][:tempfile]) &&
			(name = params[:file][:filename])
	    @error = "No file selected"
	    return haml(:upload)
	end
	# STDERR.puts "Uploading file, original name #{name.inspect}"
	new_dir = "#{savepath}#{name.split(".")[0]}"
	Dir.mkdir(new_dir)
	f = File.open("#{new_dir}/#{name}", "w")
	while blk = tmpfile.read(65536)
		f.write blk
		# STDERR.puts blk.inspect
	end
	"Upload complete"
	200
end