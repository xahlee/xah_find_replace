# -*- coding: utf-8 -*-
# perl

use utf8;
use strict;


=pod


Description:
This script does find and replace on a given foler recursively.

Features:
• multiple Find and Replace string pairs can be given.
• The find/replace strings can be set to regex or literal.
• Files can be filtered according to file name suffix matching or other criterions.
• Backup copies of original files will be made at a user specified folder that preserves all folder structures of original folder.
• A report will be generated that indicates which files has been changed, how many changes, and total number of files changed.
• files will retain their own/group/permissions settings.

usage:
1. edit the parts under the section
2. edit the subroutine fileFilterQ to set which file will be checked or skipped.

to do:
• in the report, print the strings that are changed, possibly with surrounding lines.
• allow just find without replace.
• add the GNU syntax for unix command prompt.
• Report if backup directory exists already, or provide toggle to overwrite, or some other smarties.

Date created: 2000-02
Author: Xah Lee, XahLee.org

=cut

#-- modules --

use File::Find;
use File::Path;
use File::Copy;
use Data::Dumper;

#-- arguments --

# the folder to be search on.
my $folderPath = q[/Users/xah/web/UnixResource_dir/xx];

# this is the backup folder path.
my $backupFolderPath = q[/Users/xah/xxxb];

my %findReplaceH = (
q[xah]=>q[fgcr],
q[that]=>q[hccg],
);

# $useRegexQ has values 1 or 0. If 1, inteprets the pairs in %findReplaceH
# to be regex.
my $useRegexQ = 0;

# in bytes. larger files will be skipped
my $fileSizeLimit = 500 * 1000;


#-- globals --

$folderPath =~ s[/$][]; # e.g. '/home/joe/public_html'
$backupFolderPath =~ s[/$][]; # e.g. '/tmp/joe_back';

$folderPath =~ m[/(\w+)$];
my $previousDir = $`;   # e.g. '/home/joe'
my $lastDir = $1;       # e.g. 'public_html'
my $backupRoot = $backupFolderPath . '/' . $1; # e.g. '/tmp/joe_back/public_html'

my $refLargeFiles = [];
my $totalFileChangedCount = 0;

#-- subroutines --

# fileFilterQ($fullFilePath) return true if file is desired.
sub fileFilterQ ($) {
        my $fileName = $_[0];

        if ((-s $fileName) > $fileSizeLimit) {
                push (@$refLargeFiles, $fileName);
                return 0;
        };
        if ($fileName =~ m{\.html$}) {
print "processing: $fileName\n";
return 1;};

##        if (-d $fileName) {return 0;}; # directory
##        if (not (-T $fileName)) {return 0;}; # not text file

        return 0;
};

# go through each file, accumulate a hash.
sub processFile {
        my $currentFile = $File::Find::name; # full path spect
        my $currentDir = $File::Find::dir;
        my $currentFileName = $_;

        if (not fileFilterQ($currentFile)) {
                return 1;
        }

# open file. Read in the whole file.
        if (not(open FILE, "<$currentFile")) {die("Error opening file: 
$!");};
        my $wholeFileString;
        {local $/ = undef; $wholeFileString = <FILE>;};
        if (not(close(FILE))) {die("Error closing file: $!");};

# do the replacement.
        my $replaceCount = 0;

        foreach my $key1 (keys %findReplaceH) {
                my $pattern = ($useRegexQ ? $key1 : quotemeta($key1));
                $replaceCount = $replaceCount + ($wholeFileString =~ 
s/$pattern/$findReplaceH{$key1}/g);
        };

        if ($replaceCount > 0) { # replacement has happened
                $totalFileChangedCount++;
# do backup
                # make a directory in the backup path, make a backup copy.
                my $pathAdd = $currentDir; $pathAdd =~ s[$folderPath][];
                mkpath("$backupRoot/$pathAdd", 0, 0777);
                copy($currentFile, "$backupRoot/$pathAdd/$currentFileName") or
                    die "error: file copying file failed on $currentFile\n$!";

# write to the original
                # get the file mode.
                my ($mode, $uid, $gid) = (stat($currentFile))[2,4,5];

                # write out a new file.
                if (not(open OUTFILE, ">$currentFile")) {die("Error opening file: $!");};
                print OUTFILE $wholeFileString;
                if (not(close(OUTFILE))) {die("Error closing file: $!");};

                # set the file mode.
                chmod($mode, $currentFile);
                chown($uid, $gid, $currentFile);

                print "-----^$*%$@#-------------------------------\n";
                print "$replaceCount replacements made at\n";
                print "$currentFile\n";
        }

};


#-- main body --

find(\&processFile, $folderPath);

print "--------------------------------------------\n\n\n";
print "Total of $totalFileChangedCount files changed.\n";

if (scalar @$refLargeFiles > 0) {
        print "The following large files are skipped:\n";
        print Dumper($refLargeFiles);
}


__END__
