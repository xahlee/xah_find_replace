# -*- coding: utf-8 -*-
# Python 3

import os, sys, shutil, re
import datetime

# if this this is not empty, then only these files will be processed
file_list = [
    # "/home/xah/web/ergoemacs_org/emacs/emacs.html", # example

]

find_string = "left_panel_26878"

# must be full path
input_dir = "/home/xah/web/"

input_dir = os.path.normpath(input_dir)

min_level = 1 # files and dirs inside input_dir are level 1.
max_level = 9 # inclusive

print_filename_when_no_change = False

def check_file(file_path):

    input_file = open(file_path, "r", encoding="utf-8")
    try:
        file_content = input_file.read()
    except UnicodeDecodeError:
        print("Unicode Decode Error 33167:「{:s}」".format(file_path))
        return

    input_file.close()

    open_count = 0
    close_count = 0
    open_count += file_content.count("<main")
    close_count += file_content.count("</main>")

    if open_count != close_count:
        print("◆ ", open_count, " ", close_count, file_path)

print(datetime.datetime.now())
print("Input Dir:", input_dir)

if (len(file_list) != 0):
   for ff in file_list:
       check_file(os.path.normpath(ff))
else:
    for dirPath, subdirList, fileList in os.walk(input_dir):
        curDirLevel = dirPath.count( os.sep) - input_dir.count( os.sep)
        curFileLevel = curDirLevel + 1
        if min_level <= curFileLevel <= max_level:
            for fName in fileList:
                if (re.search(r"\.html$", fName, re.U)):
                    check_file(dirPath + os.sep + fName)
                    # print ("level %d,  %s" % (curFileLevel, os.path.join(dirPath, fName)))

print("Done.")
