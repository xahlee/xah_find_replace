# -*- coding: utf-8 -*-
# Python 3

# find & replace strings in a dir

import os, sys, shutil, re
import datetime

# if this this is not empty, then only these files will be processed
file_list = [
]

input_dir = "/home/xah/web/wordyenglish_com/"
input_dir = "/home/xah/web/"

input_dir = os.path.normpath(input_dir)

min_level = 1 # files and dirs inside input_dir are level 1.
max_level = 6 # inclusive

print_file_name_has_no_change = False

# <p class="author_0"><address><a href="https://plus.google.com/112757647855302148298?rel=author" rel="author">Xah Lee</a></address>, <time>2013-06-10</time></p>

# <p class="author_0">Xah Lee,\([^<]+?\)<time>\([0-9][0-9][0-9][0-9]-[0-9][0-9]-[0-9][0-9]\)</time></p>
# <p class="author_0">Xah Lee,\([- ,…<>/time0-9]+?\)<time>\([0-9][0-9][0-9][0-9]-[0-9][0-9]-[0-9][0-9]\)</time></p>
# <div class="author_0"><address><a href="https://plus.google.com/112757647855302148298?rel=author" rel="author">Xah Lee</a></address>,\1<time>\2</time></div>

# <p class="author_0">Xah Lee,\([- ,…0-9]+?\)</p>

# <script>(function() { var po=document.createElement('script');po.type='text/javascript';po.async=true;po.src='https://apis.google.com/js/plusone.js';var s=document.getElementsByTagName('script')[0];s.parentNode.insertBefore(po,s);})();</script>
# <script>!function(d,s,id){var js,fjs=d.getElementsByTagName(s)[0],p=/^http:/.test(d.location)?'http':'https';if(!d.getElementById(id)){js=d.createElement(s);js.id=id;js.src=p+'://platform.twitter.com/widgets.js';fjs.parentNode.insertBefore(js,fjs);}}(document, 'script', 'twitter-wjs');</script>
# <div id="fb-root"></div>
# <script>(function(d, s, id) { var js, fjs = d.getElementsByTagName(s)[0]; if (d.getElementById(id)) return; js = d.createElement(s); js.id = id; js.src = "//connect.facebook.net/en_US/all.js#xfbml=1"; fjs.parentNode.insertBefore(js, fjs); }(document, 'script', 'facebook-jssdk'));</script>

# "691751129124876",

# • add text thta says it's ad
# • many pages don't have double ads

find_replace_list = [

(

"""</h1>

<div class="header66704">""",

"""</h1>

<div class="header66704">""",

) ,

]



for x in find_replace_list:
    if len(x) != 2:
        sys.exit("Error: replacement pair has more than 2 elements. Probably missing a comma.")

def replace_string_in_file(file_path):
    "Replaces all findStr by repStr in file file_path"
    backup_fname = file_path + "~bk~"

    # print "reading:", file_path
    input_file = open(file_path, "r", encoding="utf-8")
    file_content = input_file.read()
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
        if print_file_name_has_no_change == True:
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
