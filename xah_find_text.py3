# -*- coding: utf-8 -*-
# Python 3

# find text of all files in a dir

# WARNING: WORK IN PROGRESS

import os, sys, shutil, re
import datetime

# if this this is not empty, then only these files will be processed
file_list = [

]

input_dir = "/home/xah/web/xahlee_info/"
input_dir = os.path.normpath(input_dir)

min_level = 1 # files and dirs inside input_dir are level 1.
max_level = 6 # inclusive

print_file_name_has_no_change = False

find_str_list = [
r"""\N{ZERO WIDTH NO-BREAK SPACE}""",
r"""<!-- ~/web/xahlee_info/comp/xx_comp_blog.html -->"""
]



def report_string_in_file(file_path):
   "print occurances of string in file_path
   the strings to search is from find_str_list
   "

   input_file = open(file_path, "r")
   file_content = input_file.read()
   input_file.close()

   for searchStr in find_str_list:
       if mResult = search(searchStr, file_content):
           print(mResult)



print(datetime.datetime.now())
print("Input Dir:", input_dir)
for x in find_str_list:
   print("Find string:「{}」".format(x[0]))
   print("Replace string:「{}」".format(x[1]))
   print("\n")

if (len(file_list) != 0):
   for ff in file_list: report_string_in_file(os.path.normpath(ff) )
else:
    for dirPath, subdirList, fileList in os.walk(input_dir):
        curDirLevel = dirPath.count( os.sep) - input_dir.count( os.sep)
        curFileLevel = curDirLevel + 1
        if min_level <= curFileLevel <= max_level:
            for fName in fileList:
                if (re.search(r"\.html$", fName, re.U)):
                    report_string_in_file(dirPath + os.sep + fName)
                    print ("level %d,  %s" % (curFileLevel, os.path.join(dirPath, fName)))

print("Done.")
