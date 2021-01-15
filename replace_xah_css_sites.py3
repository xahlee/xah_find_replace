# -*- coding: utf-8 -*-
# Python 3

# find & replace mutiple pairs of strings in a dir or list of files

import os, sys, shutil, re
import datetime

# if this list is not empty, then only these files will be processed
# each must be full path to a file, not dir
file_list = [

"/Users/xah/web/ergoemacs_org/xahemacs.css",
"/Users/xah/web/wordyenglish_com/wordy.css",
"/Users/xah/web/xaharts_org/xaharts.css",
"/Users/xah/web/xahlee_info/lbasic.css",
"/Users/xah/web/xahlee_org/xlo22481.css",
"/Users/xah/web/xahmusic_org/xahmusic.css",
"/Users/xah/web/xahporn_org/xporn.css",
"/Users/xah/web/xahsl_org/xahsl.css",

"/Users/xah/web/wordyenglish_com/lit.css",
"/Users/xah/web/xaharts_org/lit.css",
"/Users/xah/web/xahlee_org/lit.css",
"/Users/xah/web/xahmusic_org/lit.css"

]

# must be full path. dir can end with slash or no
INPUT_DIR = "/Users/xah/web/"
MIN_LEVEL = 1 # files and dirs inside INPUT_DIR are level 1.
MAX_LEVEL = 5 # inclusive
FILE_NAME_REGEX = r"\.html$"
PRINT_FILENAME_WHEN_NO_CHANGE = False
BACKUP_FNAME_EXT = '~bk~'
DO_BACKUP = False

FIND_REPLACE_LIST = [

(
'''strong.r {color:red}''',
'',
),

# more pair here

]

# a regex string. any full path that match is skipped
DIRPATH_SKIP_REGEX =
r"emacs_manual|\
REC-SVG11-20110816|\
clojure-doc-1.8|\
css_2.1_spec|\
css_transitions|\
js_es2011|\
js_es2015|\
js_es2015_orig|\
js_es2016|\
js_es2018|\
node_api"

# <nav class="nav-back-85230"><a href="index.html">FLATLAND</a></nav>
# <iframe class="left_panel_26878" src="../web_design_panel_index_32509.html"></iframe>

# /Users/xah/web/xahlee_info/kbd/keyboard_blog_panel_index.html

# git checkout . && git clean -dxfq && time grep -r -F "same line" --include='*html' /Users/xah/xx_manual/ > xxgrep

##################################################

INPUT_DIR = os.path.normpath(INPUT_DIR)

for x in FIND_REPLACE_LIST:
    if len(x) != 2:
        sys.exit("Error: replacement pair has more than 2 elements. Probably missing a comma.")

def replace_string_in_file(file_path):
    "Replaces find/replace pairs in FIND_REPLACE_LIST in file_path"
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
        file_content = file_content.replace(a_pair[0], a_pair[1])

    if num_replaced > 0:
        print("◆ ", num_replaced, " ", file_path.replace(os.sep, "/"))
        if DO_BACKUP:
            backup_fname = file_path + BACKUP_FNAME_EXT
            os.rename(file_path, backup_fname)
        output_file = open(file_path, "w")
        output_file.write(file_content)
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
# emacs_manual|\

        if (MIN_LEVEL <= curFileLevel) and (curFileLevel <= MAX_LEVEL) and (not re.search( DIRPATH_SKIP_REGEX, dirPath, re.U)):
            # print (dirPath)
            for fName in fileList:
                if (re.search( FILE_NAME_REGEX, fName, re.U)) and (not (re.search(r"#", fName, re.U))):
                    replace_string_in_file(dirPath + os.sep + fName)
                    # print ("level %d,  %s" % (curFileLevel, os.path.join(dirPath, fName)))

print("Done.")
