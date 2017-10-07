# -*- coding: utf-8 -*-
# Python 3

# find & replace mutiple pairs of strings in a dir or list of files

import os, sys, shutil, re
import datetime

# if this list is not empty, then only these files will be processed
# each must be full path to a file, not dir
file_list = [

]

# must be full path
INPUT_DIR = "/Users/xah/web/ergoemacs_org/"

MIN_LEVEL = 1 # files and dirs inside INPUT_DIR are level 1.
MAX_LEVEL = 6 # inclusive

FILE_NAME_REGEX = r"\.html$"

PRINT_FILENAME_WHEN_NO_CHANGE = False

FIND_REPLACE_LIST = [

# <a href="act1p2.html" title="previous">←</a>

(
'''<div class="buyxahemacs97449">
Liket it? Put $5 at''',
'''<div class="buyxahemacs97449">
Liket it? Put $1 at''',
),

]

##################################################

INPUT_DIR = os.path.normpath(INPUT_DIR)

for x in FIND_REPLACE_LIST:
    if len(x) != 2:
        sys.exit("Error: replacement pair has more than 2 elements. Probably missing a comma.")

def replace_string_in_file(file_path):
    "Replaces find/replace pairs in FIND_REPLACE_LIST in file_path"
    backup_fname = file_path + '~bk~'

    input_file = open(file_path, "r", encoding="utf-8")
    try:
        file_content = input_file.read()
    except UnicodeDecodeError:
        # print("UnicodeDecodeError:{:s}".format(input_file))
        return

    input_file.close()

    num_replaced = 0
    for a_pair in FIND_REPLACE_LIST:
        num_replaced += file_content.count(a_pair[0])
        output_text = file_content.replace(a_pair[0], a_pair[1])
        file_content = output_text

    if num_replaced > 0:
        print("◆ ", num_replaced, " ", file_path.replace(os.sep, "/"))
        shutil.copy2(file_path, backup_fname)
        output_file = open(file_path, "w")
        output_file.write(output_text)
        output_file.close()
    else:
        if PRINT_FILENAME_WHEN_NO_CHANGE == True:
            print("no change:", file_path)

##################################################

print(datetime.datetime.now())
print("Input Dir:", INPUT_DIR)
for x in FIND_REPLACE_LIST:
   print("Find string:\n「{}」".format(x[0]))
   print("Replace string:\n「{}」".format(x[1]))
   print("\n")

if (len(file_list) != 0):
   for ff in file_list: replace_string_in_file(os.path.normpath(ff) )
else:
    for dirPath, subdirList, fileList in os.walk(INPUT_DIR):
        curDirLevel = dirPath.count( os.sep) - INPUT_DIR.count( os.sep)
        curFileLevel = curDirLevel + 1
        if (MIN_LEVEL <= curFileLevel) and (curFileLevel <= MAX_LEVEL) and (not re.search(r"emacs_manual|\
REC-SVG11-20110816|\
clojure-doc-1.8|\
ocaml_doc|\
css3_spec_bg|\
css_2.1_spec|\
css_3_color_spec|\
css_transitions|\
dom-whatwg|\
html5_whatwg|\
java8_doc|\
javascript_ecma-262_5.1_2011|\
javascript_ecma-262_6_2015|\
javascript_es2016|\
javascript_es6|\
jquery_doc|\
node_api|\
php-doc|\
python_doc_2.7.6|\
python_doc_3.3.3", dirPath, re.U)):
            # print (dirPath)
            for fName in fileList:
                if (re.search( FILE_NAME_REGEX, fName, re.U)) and (not (re.search(r"#", fName, re.U))):
                    replace_string_in_file(dirPath + os.sep + fName)
                    # print ("level %d,  %s" % (curFileLevel, os.path.join(dirPath, fName)))

print("Done.")
