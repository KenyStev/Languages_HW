import merge

def SortEmails(filename):
	name = filename.split(".")[0]
	merge.FilterFile(name+"/"+filename,"(\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3})")
	merge.CreateLeaves(name+"/"+filename+".filtered",5)
	merge.MergeSort(name+"/leaves/")

	return merge.GetSortedFile(filename)