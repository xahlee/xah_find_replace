# -*- coding: utf-8 -*-
# Python 3

# find & replace strings in a dir

import os, sys, shutil, re
import datetime

# if this this is not empty, then only these files will be processed
file_list = [
]

input_dir = "/home/xah/web/xahlee_info"

input_dir = os.path.normpath(input_dir)

min_level = 1 # files and dirs inside input_dir are level 1.
max_level = 7 # inclusive

print_filename_when_no_change = False

find_replace_list = [
(

"""<nav id="t5">
<ul>
<li><a href="../SpecialPlaneCurves_dir/specialPlaneCurves.html">Curves</a></li>
<li><a href="../surface/gallery.html">Surfaces</a></li>
<li><a href="../Wallpaper_dir/c0_WallPaper.html">Wallpaper Groups</a></li>
<li><a href="../MathGraphicsGallery_dir/mathGraphicsGallery.html">Gallery</a></li>
<li><a href="../math_software/mathPrograms.html">Math Software</a></li>
<li><a href="../3d/index.html">POV-Ray</a></li>
</ul>
<ul>
<li><a href="js.html">JavaScript</a></li>
<li><a href="index.html">HTML</a></li>
<li><a href="css_index.html">CSS</a></li>
</ul>
<ul>
<li><a href="../linux/linux_index.html">Linux</a></li>
<li><a href="../perl-python/index.html">Perl Python Ruby</a></li>
<li><a href="../java-a-day/java.html">Java</a></li>
<li><a href="../php/index.html">PHP</a></li>
<li><a href="http://ergoemacs.org/emacs/emacs.html">Emacs</a></li>
<li><a href="../comp/comp_lang.html">Syntax</a></li>
<li><a href="../comp/unicode_index.html">Unicode 😸 ♥</a></li>
<li><a href="../kbd/keyboarding.html">Keyboard ⌨</a></li>
</ul>
<button""",

"""<nav id="t5">
<button"""

),

]



for x in find_replace_list:
    if len(x) != 2:
        sys.exit("Error: replacement pair has more than 2 elements. Probably missing a comma.")

def replace_string_in_file(file_path):
    "Replaces find/replace pairs in find_replace_list in file_path"
    backup_fname = file_path + "~bk~"

    input_file = open(file_path, "r", encoding="utf-8")
    try:
        file_content = input_file.read()
    except UnicodeDecodeError:
        print("UnicodeDecodeError:{:s}".format(input_file))
        return

    input_file.close()

    num_replaced = 0
    for a_pair in find_replace_list:
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
        if print_filename_when_no_change == True:
            print("no change:", file_path)

#────────── ────────── ────────── ────────── ──────────

print(datetime.datetime.now())
print("Input Dir:", input_dir)
for x in find_replace_list:
   print("Find string:「{}」".format(x[0]))
   print("Replace string:「{}」".format(x[1]))
   print("\n")

if (len(file_list) != 0):
   for ff in file_list: replace_string_in_file(os.path.normpath(ff) )
else:
    for dirPath, subdirList, fileList in os.walk(input_dir):
        curDirLevel = dirPath.count( os.sep) - input_dir.count( os.sep)
        curFileLevel = curDirLevel + 1
        if min_level <= curFileLevel <= max_level:
            for fName in fileList:
                if (re.search(r"\.html$", fName, re.U)):
                    replace_string_in_file(dirPath + os.sep + fName)
                    # print ("level %d,  %s" % (curFileLevel, os.path.join(dirPath, fName)))

print("Done.")
