# -*- coding: utf-8 -*-
# Python

# replace chars into links

import os, sys, re, shutil

mydir = "c:/Users/xah/web/xahlee_org/sl"

findreplace = [
(re.compile(r'''title="(\d+)x(\d+)">❐</a>''',re.U|re.M),
r'''title="\1×\2">❐</a>'''),
]

def findString(filePath):
   "print file names who's content match the regex"

   input = open(filePath,'rb')
   s=str(input.read(),'utf-8')
   input.close()

   for couple in findreplace:
      md = re.search(couple[0], s)
      if md:
         print(filePath)
         for mm in md.groups():
            print(mm.encode('utf-8'))

def myfun(dummy, dirr, filess):
   for child in filess:
      if '.html' == os.path.splitext(child)[1] and os.path.isfile(dirr+'/'+child):
         findString(dirr+'/'+child)
os.path.walk(mydir, myfun, 3)

