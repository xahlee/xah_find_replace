# Python 3

# 2018-08-24

# change all files in a dir by 1 or more regex/replace pairs

# web site: http://xahlee.info/python/findreplace_regex.html

import os, sys, shutil, re
import datetime

# if this list is not empty, then only these files will be processed
file_list = [
    # "/home/xah/web/ergoemacs_org/emacs/emacs.html", # example
]

# must be full path
input_dir = "/Users/xah/xx_manual/"

file_extension_regex = r"\.html$"

min_level = 1 # files and dirs inside input_dir are level 1.
max_level = 9 # inclusive

do_backup = True
backup_suffix = "~~"

# find and replace pairs here. Each is a 2-tuple. first element is regex object, second is replace string
find_replace_list = [
   (re.compile(r'''<meta name="[A-Za-z]+" content="[-.,0-9]+">''', re.U|re.M|re.DOTALL),
    r''),

   # more find and replace pairs here
]

##################################################

def replace_string_in_file(fpath):
   "Replaces all strings by regex in find_replace_list at fpath."

   input_file = open(fpath, "r", encoding="utf-8")

   try:
      file_content = input_file.read()
   except UnicodeDecodeError:
      print("UnicodeDecodeError:{:s}".format(input_file))
      return

   input_file.close()

   num_replaced = 0
   for a_pair in find_replace_list:
      tem_tuple = re.subn(a_pair[0], a_pair[1], file_content)
      output_text = tem_tuple[0]
      num_replaced += tem_tuple[1]
      file_content = output_text

   if (num_replaced > 0):
      print(("* changed %d %s" % (num_replaced, fpath) ))

      if do_backup:
         shutil.copy2(fpath, fpath + backup_suffix)

      output_file = open(fpath, "r+b")
      output_file.read() # we do this way to preserve file creation date
      output_file.seek(0)
      output_file.write(output_text.encode("utf-8"))
      output_file.truncate()
      output_file.close()

##################################################

print(datetime.datetime.now())
print("Input Dir:", input_dir)
for x in find_replace_list:
   print("Find regex:{}".format(x[0]))
   print("Replace pattern:{}".format(x[1]))
   print("\n")

if (len(file_list) != 0):
   for ff in file_list: replace_string_in_file(os.path.normpath(ff) )
else:
    for dirPath, subdirList, fileList in os.walk(input_dir):
        curDirLevel = dirPath.count( os.sep) - input_dir.count( os.sep)
        curFileLevel = curDirLevel + 1
        if min_level <= curFileLevel <= max_level:
            for fName in fileList:
                if (re.search(file_extension_regex, fName, re.U)):
                    replace_string_in_file(dirPath + os.sep + fName)

print("Done.")
