# -*- coding: utf-8 -*-
# Python

# change all files in a dir. 
# using mulitple regex/replace pairs 

# last used at least: 2012-03-14

import os, sys, shutil, re

input_dir = "/home/xah/web/"

find_replace_list = [

   (re.compile(ur"""<h1>([^<]+?)</h1>[ \n]+9ea99fef-0623-fc80-597f-37fdc6ffeb0a""", re.U|re.M|re.DOTALL), ur"""9ea99fef-0623-fc80-597f-37fdc6ffeb0a
<h1>\1</h1>"""),

# (re.compile(ur"""•8017015673""", re.U|re.M|re.DOTALL), ur"""<header>
# <span class="xsignet">∑</span> <a href="http://ergoemacs.org/index.html">ErgoEmacs</a> ◆ <span id="e1α"><a href="http://ergoemacs.org/emacs/emacs.html">Emacs</a> ◇ <a href="http://ergoemacs.org/emacs/elisp.html">Lisp</a></span> ◆ <a href="http://ergoemacs.org/emacs/blog.html">Blog</a> ◆ <span id="e2α"><a href="http://ergoemacs.org/emacs_manual/emacs/index.html">Emacs</a> ◇ <a href="http://ergoemacs.org/emacs_manual/elisp/index.html">Lisp</a></span> ◆ <a href="http://ergoemacs.org/emacs/buy_xah_emacs_tutorial.html">Buy Tutorial</a>
# <form action="http://www.google.com" id="cse-search-box"> <div> <input type="hidden" name="cx" value="partner-pub-5125343095650532:8381157956" /> <input type="hidden" name="ie" value="UTF-8" /> <input type="text" name="q" size="55" /> <input type="submit" name="sa" value="Search" /> </div> </form><script src="http://www.google.com/coop/cse/brand?form=cse-search-box&amp;lang=en"></script>
# </header>"""),


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
   file_content = unicode(input_file.read(), "utf-8")
   input_file.close()

   num_replaced = 0
   for a_pair in find_replace_list:
      tem_tuple = re.subn(a_pair[0], a_pair[1], file_content)
      output_text = tem_tuple[0]
      num_replaced += tem_tuple[1]
      file_content = output_text

   if (num_replaced > 0):
      print ("◆ %d %s" % (num_replaced, file_path.replace("/cygdrive/c/Users/h3", "~")) )

      shutil.copy2(file_path, backup_fname)
      output_file = open(file_path, "r+b")
      output_file.read() # we do this way to preserve file creation date
      output_file.seek(0)
      output_file.write(output_text.encode("utf-8"))
      output_file.truncate()
      output_file.close()

#      os.rename(file_path, backup_fname)
#      os.rename(tempName, file_path)


def process_file(dummy, current_dir, file_list):
   for child in file_list:
#      if "pd.html" == child:
      if re.search(r".+\.html$", child, re.U) and os.path.isfile(current_dir + "/" + child):
         replace_string_in_file(current_dir + "/" + child)

os.path.walk(input_dir, process_file, "dummy")

print "Done."
