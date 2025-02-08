# SPDX-License-Identifier: Unlicense

use strict;
use warnings;

use Getopt::Long;

my $VERSION = '0.1.0';

GetOptions( 'help' => \my $help
          , 'version' => \my $version
          );

if ($help) {
    print 'Print the user name associated with the current effective user ID.
Same as id -un.

      --help        display this help and exit
      --version     output version information and exit
';
    exit 0;
}

if ($version) {
    print 'whoami (perl polyutils) ' . $VERSION . "\n";
    exit 0;
}

print getlogin() . "\n";

