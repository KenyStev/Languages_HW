require_relative '../merge/merge'

def SortEmails(filename)
	name = "#{filename.split(".")[0]}"
	FilterFile("#{name}/#{filename}",/\A[\w+\-.]+@[a-z\d\-]+(\.[a-z\d\-]+)*\.[a-z]+\z/i)
	CreateLeaves("#{name}/#{filename}.filtered",5)
	MergeSort("#{name}/leaves/")

	GetSortedFile(filename)
end

