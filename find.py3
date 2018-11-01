# -*- coding: utf-8 -*-
# Python 3

# 2016-01-22 temp script. report file such that count of  “<main” and count of </main> are different

import os, sys, shutil, re
import datetime

# if this this is not empty, then only these files will be processed
file_list = [
    # "/home/xah/web/ergoemacs_org/emacs/emacs.html", # example

"/Users/xah/web/ergoemacs_org/emacs/elisp_htmlize_css_code.css",
"/Users/xah/web/ergoemacs_org/emacs/emacs_org_mode.css",
"/Users/xah/web/ergoemacs_org/emacs/kb-layout-2.css",
"/Users/xah/web/ergoemacs_org/emacs/kb-layout.css",
"/Users/xah/web/ergoemacs_org/emacs/lisp_keyboard_layout.css",
"/Users/xah/web/ergoemacs_org/emacs/xx_keyboard_design.css",
"/Users/xah/web/ergoemacs_org/emacs_manual/style_2016-12-21.css",
"/Users/xah/web/ergoemacs_org/xahemacs.css",
"/Users/xah/web/wordyenglish_com/alice/alice.css",
"/Users/xah/web/wordyenglish_com/arabian_nights/xx/style.css",
"/Users/xah/web/wordyenglish_com/arabian_nights/xx_full_2017-05-13/widgets/drsugiyama.css",
"/Users/xah/web/wordyenglish_com/arabian_nights/xx_full_2017-05-13/widgets/style.css",
"/Users/xah/web/wordyenglish_com/lit.css",
"/Users/xah/web/wordyenglish_com/lojban/hrefgram2/lb.css",
"/Users/xah/web/wordyenglish_com/lojban/skari.css",
"/Users/xah/web/wordyenglish_com/musing/pinyin_frequency.css",
"/Users/xah/web/wordyenglish_com/titus/titus.css",
"/Users/xah/web/wordyenglish_com/words/vc.css",
"/Users/xah/web/wordyenglish_com/wordy.css",
"/Users/xah/web/xaharts_org/lit.css",
"/Users/xah/web/xaharts_org/xaharts.css",
"/Users/xah/web/xahlee_info/REC-SVG11-20110816/images/styling/mystyle.css",
"/Users/xah/web/xahlee_info/REC-SVG11-20110816/style/W3C-REC.css",
"/Users/xah/web/xahlee_info/REC-SVG11-20110816/style/svg-style.css",
"/Users/xah/web/xahlee_info/clojure-doc-1.8/javadoc/stylesheet.css",
"/Users/xah/web/xahlee_info/clojure-doc-1.8/static/clojure.css",
"/Users/xah/web/xahlee_info/clojure-doc-1.8/static/internal.css",
"/Users/xah/web/xahlee_info/clojure-doc-1.8/static/wiki.css",
"/Users/xah/web/xahlee_info/clojure/Clojure_1.8_cheatsheet_files/asciidoctor-mod.css",
"/Users/xah/web/xahlee_info/clojure/Clojure_1.8_cheatsheet_files/clojureorg.css",
"/Users/xah/web/xahlee_info/clojure/Clojure_1.8_cheatsheet_files/css.css",
"/Users/xah/web/xahlee_info/clojure/Clojure_1.8_cheatsheet_files/normalize.css",
"/Users/xah/web/xahlee_info/clojure/Clojure_1.8_cheatsheet_files/webflow.css",
"/Users/xah/web/xahlee_info/clojure/clojure_cheatsheet.css",
"/Users/xah/web/xahlee_info/cmaci/ca/ff.css",
"/Users/xah/web/xahlee_info/cmaci/notation/HTMLFiles/Default.css",
"/Users/xah/web/xahlee_info/css_2.1_spec/style/default.css",
"/Users/xah/web/xahlee_info/css_transitions/CSS_Transitions_files/W3C-WD.css",
"/Users/xah/web/xahlee_info/css_transitions/CSS_Transitions_files/default.css",
"/Users/xah/web/xahlee_info/godoc/golang_spec_files/jquery.css",
"/Users/xah/web/xahlee_info/godoc/golang_spec_files/style.css",
"/Users/xah/web/xahlee_info/java8_doc/api/resources/fonts/dejavu.css",
"/Users/xah/web/xahlee_info/java8_doc/api/stylesheet.css",
"/Users/xah/web/xahlee_info/java8_doc/jdk/api/attach/spec/resources/fonts/dejavu.css",
"/Users/xah/web/xahlee_info/java8_doc/jdk/api/attach/spec/stylesheet.css",
"/Users/xah/web/xahlee_info/java8_doc/jdk/api/javac/tree/resources/fonts/dejavu.css",
"/Users/xah/web/xahlee_info/java8_doc/jdk/api/javac/tree/stylesheet.css",
"/Users/xah/web/xahlee_info/java8_doc/jdk/api/javadoc/doclet/resources/fonts/dejavu.css",
"/Users/xah/web/xahlee_info/java8_doc/jdk/api/javadoc/doclet/stylesheet.css",
"/Users/xah/web/xahlee_info/java8_doc/jdk/api/javadoc/taglet/stylesheet.css",
"/Users/xah/web/xahlee_info/java8_doc/jdk/api/jconsole/spec/resources/fonts/dejavu.css",
"/Users/xah/web/xahlee_info/java8_doc/jdk/api/jconsole/spec/stylesheet.css",
"/Users/xah/web/xahlee_info/java8_doc/jdk/api/jpda/jdi/resources/fonts/dejavu.css",
"/Users/xah/web/xahlee_info/java8_doc/jdk/api/jpda/jdi/stylesheet.css",
"/Users/xah/web/xahlee_info/java8_doc/jre/api/javaws/jnlp/resources/fonts/dejavu.css",
"/Users/xah/web/xahlee_info/java8_doc/jre/api/javaws/jnlp/stylesheet.css",
"/Users/xah/web/xahlee_info/java8_doc/jre/api/management/extension/resources/fonts/dejavu.css",
"/Users/xah/web/xahlee_info/java8_doc/jre/api/management/extension/stylesheet.css",
"/Users/xah/web/xahlee_info/java8_doc/jre/api/net/httpserver/spec/resources/fonts/dejavu.css",
"/Users/xah/web/xahlee_info/java8_doc/jre/api/net/httpserver/spec/stylesheet.css",
"/Users/xah/web/xahlee_info/java8_doc/jre/api/net/socketoptions/spec/stylesheet.css",
"/Users/xah/web/xahlee_info/java8_doc/jre/api/nio/sctp/spec/resources/fonts/dejavu.css",
"/Users/xah/web/xahlee_info/java8_doc/jre/api/nio/sctp/spec/stylesheet.css",
"/Users/xah/web/xahlee_info/java8_doc/jre/api/plugin/dom/org/w3c/dom/css",
"/Users/xah/web/xahlee_info/java8_doc/jre/api/plugin/dom/resources/fonts/dejavu.css",
"/Users/xah/web/xahlee_info/java8_doc/jre/api/plugin/dom/stylesheet.css",
"/Users/xah/web/xahlee_info/java8_doc/jre/api/security/jaas/spec/resources/fonts/dejavu.css",
"/Users/xah/web/xahlee_info/java8_doc/jre/api/security/jaas/spec/stylesheet.css",
"/Users/xah/web/xahlee_info/java8_doc/jre/api/security/jgss/spec/resources/fonts/dejavu.css",
"/Users/xah/web/xahlee_info/java8_doc/jre/api/security/jgss/spec/stylesheet.css",
"/Users/xah/web/xahlee_info/java8_doc/jre/api/security/smartcardio/spec/resources/fonts/dejavu.css",
"/Users/xah/web/xahlee_info/java8_doc/jre/api/security/smartcardio/spec/stylesheet.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/css/guide.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/css/index.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/css/jdk_index_style.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/css/manpage.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/css/tools.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/2d/spec/catalog.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/2d/spec/document.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/beans/spec/catalog.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/beans/spec/document.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/imf/api-sample/stylesheet.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/io/doc.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/javadoc/deprecation/doc.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/javaws/developersguide/styles/style1.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/jmx/overview/catalog.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/jmx/overview/document.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/jmx/tutorial/catalog.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/jmx/tutorial/document.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/jni/spec/catalog.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/jni/spec/document.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/lang/doc.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/management/css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/management/css/advanced.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/management/css/default.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/management/css/ipg.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/management/doc.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/management/tooldoc.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/net/css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/net/css/styles1.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/scripting/css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/scripting/css/styles1.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/sound/css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/sound/css/styles1.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/sound/programmer_guide/styles/style1.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/versioning/spec/catalog.css",
"/Users/xah/web/xahlee_info/java8_doc/technotes/guides/versioning/spec/document.css",
"/Users/xah/web/xahlee_info/javascript_ecma-262_5.1_2011/JavaScript_ECMA-262_5.1/es5.css",
"/Users/xah/web/xahlee_info/javascript_ecma-262_6_2015/es6.css",
"/Users/xah/web/xahlee_info/javascript_es2016/ECMAScript_2016_Language_Specification_files/ecmarkup.css",
"/Users/xah/web/xahlee_info/javascript_es2016/ECMAScript_2016_Language_Specification_files/github.min.css",
"/Users/xah/web/xahlee_info/javascript_es2016/ecmarkup.css",
"/Users/xah/web/xahlee_info/javascript_es6/es6.css",
"/Users/xah/web/xahlee_info/javascript_es6/js_es6.css",
"/Users/xah/web/xahlee_info/javascript_es6/mod55209.css",
"/Users/xah/web/xahlee_info/jquery_doc/assets/bundle.css",
"/Users/xah/web/xahlee_info/js/css_data_uri_scheme.css",
"/Users/xah/web/xahlee_info/js/css_tabs_a.css",
"/Users/xah/web/xahlee_info/js/css_transition.css",
"/Users/xah/web/xahlee_info/js/ex/css_onclick.css",
"/Users/xah/web/xahlee_info/js/ex/icon_css_hack.css",
"/Users/xah/web/xahlee_info/js/ex_css_progress_bar.css",
"/Users/xah/web/xahlee_info/js/google-code-prettify/gcp/prettify.css",
"/Users/xah/web/xahlee_info/js/js_d3_calendar_view.css",
"/Users/xah/web/xahlee_info/js/svg_basics_57328.css",
"/Users/xah/web/xahlee_info/js/tabs/tab.css",
"/Users/xah/web/xahlee_info/js/xx-css_pie_menu_2.css",
"/Users/xah/web/xahlee_info/js/xx-css_rotating_cube.css",
"/Users/xah/web/xahlee_info/js/xx-css_semi-circle_menu.css",
"/Users/xah/web/xahlee_info/js/xx-css_triangle.css",
"/Users/xah/web/xahlee_info/js/xx_webgl_light_ball/brisk_hindsight_83430.css",
"/Users/xah/web/xahlee_info/lbasic.css",
"/Users/xah/web/xahlee_info/llang.css",
"/Users/xah/web/xahlee_info/lmath.css",
"/Users/xah/web/xahlee_info/lpanel_f9j0y.css",
"/Users/xah/web/xahlee_info/lsyntax_highlight.css",
"/Users/xah/web/xahlee_info/math_software/math_software.css",
"/Users/xah/web/xahlee_info/node_api/assets/sh.css",
"/Users/xah/web/xahlee_info/node_api/assets/style.css",
"/Users/xah/web/xahlee_info/ocaml_doc/htmlman/libref/style.css",
"/Users/xah/web/xahlee_info/ocaml_doc/htmlman/manual.css",
"/Users/xah/web/xahlee_info/ocaml_doc/htmlman/manual_modded_orig_2016-07-31.css",
"/Users/xah/web/xahlee_info/php/php.css",
"/Users/xah/web/xahlee_info/python_doc_2.7.6/_static/basic.css",
"/Users/xah/web/xahlee_info/python_doc_2.7.6/_static/default.css",
"/Users/xah/web/xahlee_info/python_doc_2.7.6/_static/pygments.css",
"/Users/xah/web/xahlee_info/python_doc_3.3.3/_static/basic.css",
"/Users/xah/web/xahlee_info/python_doc_3.3.3/_static/default.css",
"/Users/xah/web/xahlee_info/python_doc_3.3.3/_static/pydoctheme.css",
"/Users/xah/web/xahlee_info/python_doc_3.3.3/_static/pygments.css",
"/Users/xah/web/xahlee_info/xxjs/test2018-02-07/test_js_27377.css",
"/Users/xah/web/xahlee_org/lit.css",
"/Users/xah/web/xahlee_org/p/mopi/mopi.css",
"/Users/xah/web/xahlee_org/xahleeorg.css",
"/Users/xah/web/xahmusic_org/lit.css",
"/Users/xah/web/xahmusic_org/xahmusic.css",
"/Users/xah/web/xahporn_org/xporn.css",
"/Users/xah/web/xahsl_org/xahsl.css"

]

find_string = "left_panel_26878"

# must be full path
input_dir = "/home/xah/web/ergoemacs_org/"

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
