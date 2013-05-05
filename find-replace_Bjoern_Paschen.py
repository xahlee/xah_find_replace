# -*- coding: utf-8 -*-
# Python

# find & replace strings in a dir
# by Bjoern Paschen, 2012-04-08

import os, re

# if this this is not empty, then only these files will be processed
my_files = []

input_dir = "c:/Users/h3/web/xahlee_org/"

min_level = 4 # files and dirs inside input_dir are level 1
max_level = 4 # inclusive

find_replace_list = [
    (u'<iframe style="width:100%;border:none" src="http://xahlee.org/footer.html"></iframe>', 
     u'<iframe style="width:100%;border:none" src="../footer.html"></iframe>',
     ),
    (u'testing',
     u'done',
     ),
    (u'utf8-test',
     u'ƿƺ'
     ),
    ]

def replace_string_in_file(f):
    content = u''
    temp = u''
    writeout = False
    
    with open(os.path.normpath(f), 'rb') as data:
        content = unicode(data.read(), 'utf8')

    for rep in find_replace_list:
        temp = re.sub(rep[0], rep[1], content, re.M)
        if temp != content:
            content = temp
            writeout = True

    if writeout:
        with open(os.path.normpath(f), 'w+b') as data:
            data.write(content.encode('utf8'))

def calc_level(input_dir, actual_dir):
    def depth(d):
        # strings are sequences, so we can use filter()
        # calculates the length of the sequence containing only teh seperators
        return len(filter(lambda x: x == os.sep, os.path.abspath(d)))
    
    # relative depth, starts at 1, not at 0
    return depth(actual_dir) - depth(input_dir) + 1

def process_dir(root, dirs, files):
    # calculate depth level
    level = calc_level(input_dir, root)

    if min_level <= level <= max_level:
        [replace_string_in_file(os.path.join(root, f)) for f in files if (os.path.splitext(f)[1] == '.html' and os.path.isfile(os.path.join(root, f)))]

# main routine
if __name__ == "__main__":
    if my_files:
        [replace_string_in_file(x) for x in my_files]
    else:
        [process_dir(root, dirs, files) for root, dirs, files in os.walk(input_dir)]