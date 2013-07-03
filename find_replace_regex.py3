# -*- coding: utf-8 -*-
# Python 3

# change all files in a dir. 
# using mulitple regex/replace pairs 

# last used at least: 2012-03-14

import os, sys, shutil, re
import datetime

# if this this is not empty, then only these files will be processed
file_list = [
]

input_dir = "/cygdrive/c/Users/h3/web/xahlee_org/Periodic_dosage_dir/bangu"
input_dir = "/home/xah/web/xahlee_info"

min_level = 1 # files and dirs inside input_dir are level 1.
max_level = 6 # inclusive

find_replace_list = [

(re.compile(r"""<header>.+</header>""", re.U|re.M|re.DOTALL), r"""<nav id="t5">
<ul>
<li><a href="http://xahlee.info/SpecialPlaneCurves_dir/specialPlaneCurves.html">Curves</a></li>
<li><a href="http://xahlee.info/surface/gallery.html">Surfaces</a></li>
<li><a href="http://xahlee.info/Wallpaper_dir/c0_WallPaper.html">Wallpaper Groups</a></li>
<li><a href="http://xahlee.info/MathGraphicsGallery_dir/mathGraphicsGallery.html">Gallery</a></li>
<li><a href="http://xahlee.info/math_software/mathPrograms.html">Software</a></li>
<li><a href="http://xahlee.info/3d/index.html">POV-Ray</a></li>
</ul>
<ul>
<li><a href="http://xahlee.info/linux/linux_index.html">Linux</a></li>
<li><a href="http://xahlee.info/perl-python/index.html">Perl Python Ruby</a></li>
<li><a href="http://xahlee.info/js/js.html">JavaScript</a></li>
<li><a href="http://xahlee.info/java-a-day/java.html">Java</a></li>
<li><a href="http://xahlee.info/php/index.html">PHP</a></li>
<li><a href="http://xahlee.info/js/index.html">HTML</a></li>
<li><a href="http://xahlee.info/js/css_index.html">CSS</a></li>
<li><a href="http://ergoemacs.org/emacs/emacs.html">Emacs</a></li>
<li><a href="http://xahlee.info/comp/comp_lang.html">Syntax</a></li>
<li><a href="http://xahlee.info/comp/unicode_index.html">Symbology</a></li>
<li><a href="http://xahlee.info/kbd/keyboarding.html">Keyboard ⌨</a></li>
</ul>
<form action="http://www.google.com" id="cse-search-box"> <div> <input type="hidden" name="cx" value="partner-pub-5125343095650532:1853288892" /> <input type="hidden" name="ie" value="UTF-8" /> <input type="text" name="q" size="20" /> <input type="submit" name="sa" value="Search" /> </div> </form><script src="http://www.google.com/coop/cse/brand?form=cse-search-box&amp;lang=en"></script>
</nav>"""),

# (re.compile(ur"""<img src="([^"]+?)" alt="([^"]+?)" width="([0-9]+)" height="([0-9]+)">
# <figcaption>""", re.U|re.M),
# ur"""<img src="\1" alt="\2" width="\3" height="\4" />
# <figcaption>"""),

# (re.compile(ur"""title="(\d+)x(\d+)">❐</a>""",re.U|re.M),
# ur"""title="\1×\2">❐</a>"""),
]

def replace_string_in_file(file_path):
   "Replaces all strings by regex in find_replace_list at file_path."
   backup_fname = file_path + "~re~"

   # print "reading:", file_path
   input_file = open(file_path, "rb")
   file_content = str(input_file.read(), "utf-8")
   input_file.close()

   num_replaced = 0
   for a_pair in find_replace_list:
      tem_tuple = re.subn(a_pair[0], a_pair[1], file_content)
      output_text = tem_tuple[0]
      num_replaced += tem_tuple[1]
      file_content = output_text

   if (num_replaced > 0):
      print(("◆ %d %s" % (num_replaced, file_path.replace("/cygdrive/c/Users/h3", "~")) ))

      shutil.copy2(file_path, backup_fname)
      output_file = open(file_path, "r+b")
      output_file.read() # we do this way to preserve file creation date
      output_file.seek(0)
      output_file.write(output_text.encode("utf-8"))
      output_file.truncate()
      output_file.close()

#      os.rename(file_path, backup_fname)
#      os.rename(tempName, file_path)

# def process_file(dummy, current_dir, file_list):
#    for child in file_list:
#       if re.search(r".+\.html$", child, re.U) and os.path.isfile(current_dir + "/" + child):
#          replace_string_in_file(current_dir + "/" + child)

# ────────── ────────── ────────── ────────── ──────────

print(datetime.datetime.now())
print("Input Dir:", input_dir)
for x in find_replace_list:
   print("Find regex:「{}」".format(x[0]))
   print("Replace pattern:「{}」".format(x[1]))
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

# os.path.walk(input_dir, process_file, "dummy")

print("Done.")
